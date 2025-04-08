package img

import (
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

func disabled(router *http.ServeMux, path string) {
	router.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		var src = r.URL.Query().Get("src")
		//add header to prevent caching
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		http.Redirect(w, r, src, http.StatusMovedPermanently)
	})
}

func enabled(router *http.ServeMux, route string, staticPath string) {
	targetPath = staticPath
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

	router.HandleFunc(route, Optimize)
}

func Init(router *http.ServeMux, route string, staticPath string, enabledOrDisabled string) {
	switch enabledOrDisabled {
	case "true":
		enabled(router, route, staticPath)
	case "false":
		disabled(router, route)
	default:
		disabled(router, route)
	}
}
