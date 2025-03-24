package img

import (
	"axo/axo"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/h2non/bimg"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

var (
	staticPath     string
	UseBreakpoints bool
	BreakpointList struct {
		Width  []int
		Height []int
	}
	cacheDir       string
	maxImageWidth  int
	maxImageHeight int
	cacheExp       int
	maxCacheSize   int
	memoryCache    *expirable.LRU[string, []byte]
	cacheMutex     sync.RWMutex
)

func Init() {
	staticPath = os.Getenv("STATIC_PATH")
	UseBreakpoints = os.Getenv("USE_BREAKPOINTS") == "true"
	BreakpointList = struct {
		Width  []int
		Height []int
	}{
		Width:  parseEnvToIntSlice("BREAKPOINT_WIDTHS"),
		Height: parseEnvToIntSlice("BREAKPOINT_HEIGHTS"),
	}
	cacheDir = os.Getenv("CACHE_DIR")
	maxImageWidth = parseEnvToInt("MAX_IMAGE_WIDTH", 1440)
	maxImageHeight = parseEnvToInt("MAX_IMAGE_HEIGHT", 990)
	cacheExp = parseEnvToInt("CACHE_EXPIRATION", 14400)  // 4 hours
	maxCacheSize = parseEnvToInt("MAX_CACHE_SIZE", 1024) // Maximum amount of items in the cache
	memoryCache = expirable.NewLRU[string, []byte](maxCacheSize, nil, time.Duration(cacheExp)*time.Second)
}

func parseEnvToInt(envVar string, defaultValue int) int {
	valueStr := os.Getenv(envVar)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func parseEnvToIntSlice(envVar string) []int {
	valueStr := os.Getenv(envVar)
	if valueStr == "" {
		return nil
	}
	values := strings.Split(valueStr, ",")
	intValues := make([]int, len(values))
	for i, v := range values {
		intValues[i], _ = strconv.Atoi(v)
	}
	return intValues
}

func generateCacheKey(params ...string) string {
	hash := sha256.New()
	hash.Write([]byte(strings.Join(params, "-")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func getImageType(format string) bimg.ImageType {
	switch format {
	case "jpg", "jpeg":
		return bimg.JPEG
	case "png":
		return bimg.PNG
	case "webp":
		return bimg.WEBP
	case "tiff":
		return bimg.TIFF
	case "gif":
		return bimg.GIF
	case "heif":
		return bimg.HEIF
	default:
		return bimg.UNKNOWN
	}
}

func roundToNearestBreakpoint(value int, breakpoints []int) int {
	if len(breakpoints) == 0 {
		return value
	}
	closest := breakpoints[0]
	for _, b := range breakpoints {
		if b >= value && (closest < value || b < closest) {
			closest = b
		}
	}
	return closest
}

func checkAndCreateCacheDir() error {
	if cacheDir == "memory" {
		return nil
	}
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create cache directory: %w", err)
		}
	}
	return nil
}

func getCachedImage(cacheKey string) (io.Reader, error) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	if cacheDir == "memory" {
		if cachedImg, found := memoryCache.Get(cacheKey); found {
			return bytes.NewReader(cachedImg), nil
		}
		return nil, fmt.Errorf("cache miss or expired")
	}

	cacheFilePath := filepath.Join(cacheDir, cacheKey)
	if cachedImg, err := os.ReadFile(cacheFilePath); err == nil {
		fileInfo, err := os.Stat(cacheFilePath)
		if err == nil && time.Since(fileInfo.ModTime()) < time.Duration(cacheExp)*time.Second {
			return bytes.NewReader(cachedImg), nil
		}
		os.Remove(cacheFilePath)
	}
	return nil, fmt.Errorf("cache miss or expired")
}

func writeCacheFile(cacheKey string, img []byte) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if cacheDir == "memory" {
		memoryCache.Add(cacheKey, img)
		return nil
	}

	cacheFilePath := filepath.Join(cacheDir, cacheKey)
	cacheFile, err := os.Create(cacheFilePath)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %w", err)
	}
	defer cacheFile.Close()

	_, err = cacheFile.Write(img)
	if err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

func Optimize(w http.ResponseWriter, r *http.Request) {
	defer func() {
		runtime.GC()
		// debug.FreeOSMemory()
	}()
	if err := checkAndCreateCacheDir(); err != nil {
		axo.Error(w, "Error creating cache directory", http.StatusInternalServerError)
		return
	}

	img := r.URL.Query().Get("src")
	if img == "" || strings.Contains(img, "..") {
		axo.Error(w, "src is required", http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(img, "http") {
		switch {
		case strings.HasPrefix(img, fmt.Sprintf("/api/%s/", staticPath)):
		case strings.HasPrefix(img, fmt.Sprintf("/%s/", staticPath)):
			img = path.Join("/api", img)
		default:
			img = path.Join("/api", staticPath, img)
		}
	}

	qualityStr := r.URL.Query().Get("quality")
	quality, err := strconv.Atoi(qualityStr)
	if err != nil || qualityStr == "" || quality < 0 || quality > 100 {
		quality = 90
	}

	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")

	width, err := strconv.Atoi(widthStr)
	if err != nil || widthStr == "" {
		width = 0
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil || heightStr == "" {
		height = 0
	}

	if width <= 0 && height <= 0 {
		width = maxImageWidth
	}

	if UseBreakpoints {
		if width > 0 {
			width = roundToNearestBreakpoint(width, BreakpointList.Width)
		}
		if height > 0 {
			height = roundToNearestBreakpoint(height, BreakpointList.Height)
		}
	}

	source, err := url.Parse(img)
	if err != nil {
		axo.Error(w, "src is invalid", http.StatusBadRequest)
		return
	}

	// if !slices.Contains(config.Get().AllowList, source.Host) && strings.HasPrefix(source.Path, "http") {
	// 	axo.Error(w, fmt.Sprintf("NOT_ALLOWED %v", source.Host), http.StatusForbidden)
	// 	return
	// }

	filePath := filepath.Join(strings.TrimPrefix(source.Path, "/api/"))
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		axo.Error(w, "IMAGE_ERROR", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading image: %v\n", err)
		axo.Error(w, "IMAGE_READ_ERROR", http.StatusInternalServerError)
		return
	}

	if width <= 0 && height <= 0 {
		width = maxImageWidth
		height = maxImageHeight
	} else {
		if width <= 0 {
			if height > 0 {
				imgSize, err := bimg.NewImage(buf).Size()
				if err != nil {
					axo.Error(w, "IMAGE_SIZE_ERROR", http.StatusInternalServerError)
					return
				}
				aspectRatio := float64(imgSize.Width) / float64(imgSize.Height)
				width = int(float64(height) * aspectRatio)
				if width > maxImageWidth {
					width = maxImageWidth
					height = int(float64(width) / aspectRatio)
				}
			}
		}
		if height <= 0 {
			if width > 0 {
				imgSize, err := bimg.NewImage(buf).Size()
				if err != nil {
					axo.Error(w, "IMAGE_SIZE_ERROR", http.StatusInternalServerError)
					return
				}
				aspectRatio := float64(imgSize.Height) / float64(imgSize.Width)
				height = int(float64(width) * aspectRatio)
				if height > maxImageHeight {
					height = maxImageHeight
					width = int(float64(height) / aspectRatio)
				}
			}
		}
	}

	if width > maxImageWidth {
		width = maxImageWidth
	}
	if height > maxImageHeight {
		height = maxImageHeight
	}

	format := r.URL.Query().Get("format")
	if format == "" {
		format = "jpg"
	}

	cacheKey := generateCacheKey(source.Path, strconv.Itoa(width), strconv.Itoa(height), format, strconv.Itoa(quality))

	if cachedImg, err := getCachedImage(cacheKey); err == nil {
		w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))
		w.Header().Set("X-Cache-Status", "HIT")
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%v", cacheExp))
		w.WriteHeader(http.StatusOK)
		io.Copy(w, cachedImg)
		return
	}

	options := bimg.Options{
		Width:   width,
		Height:  height,
		Quality: quality,
		Type:    getImageType(format),
	}

	resizedImg, err := bimg.NewImage(buf).Process(options)
	if err != nil {
		fmt.Printf("Error processing image: %v\n", err)
		axo.Error(w, "IMAGE_PROCESS_ERROR", http.StatusInternalServerError)
		return
	}

	if err := writeCacheFile(cacheKey, resizedImg); err != nil {
		fmt.Printf("Error writing cache file: %v\n", err)
		axo.Error(w, "CACHE_WRITE_ERROR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Cache-Status", "MISS")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%v, must-revalidate", cacheExp))
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))
	w.WriteHeader(http.StatusOK)
	if cacheDir == "memory" {
		if cachedImg, found := memoryCache.Get(cacheKey); found {
			io.Copy(w, bytes.NewReader(cachedImg))
		} else {
			axo.Error(w, "CACHE_READ_ERROR", http.StatusInternalServerError)
		}
	} else {
		cacheFilePath := filepath.Join(cacheDir, cacheKey)
		cacheFile, _ := os.Open(cacheFilePath)
		defer cacheFile.Close()
		cacheFile.Seek(0, 0)
		io.Copy(w, cacheFile)
	}
}
