package middlewares

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Logger : is a simple middleware that logs incoming requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		clientIPAdress := r.RemoteAddr
		fullpath := r.URL.Path
		if r.URL.RawQuery != "" {
			fullpath += "?" + r.URL.RawQuery
		}
		logRecord := []string{currentTime, clientIPAdress, r.Method, fullpath}
		print := fmt.Sprintf(
			"\033[36m[AXO]\033[0m \033[90m%s\033[0m | \033[33m%s\033[0m | \033[32m%s\033[0m | \033[34m%s\033[0m\n",
			currentTime, clientIPAdress, r.Method, fullpath,
		)

		if os.Getenv("NOLOG") == "" {
			fmt.Print(print)

			// Logs klasörünü oluştur (eğer yoksa)
			logDir := "logs"
			if _, err := os.Stat(logDir); os.IsNotExist(err) {
				_ = os.Mkdir(logDir, os.ModePerm)
			}

			logFilePath := filepath.Join(logDir, "log.clf")
			logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				fmt.Printf("Error opening file: %v", err)
				return
			}
			defer logFile.Close()

			csvWriter := csv.NewWriter(logFile)
			defer csvWriter.Flush()

			if err := csvWriter.Write(logRecord); err != nil {
				fmt.Printf("Error writing to CSV: %v", err)
			}

			// Log dosya boyutu kontrolü (8MB)
			fileInfo, _ := logFile.Stat()
			if fileInfo.Size() > 8*1024*1024 {
				archiveOldLog(logFilePath, logDir)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// archiveOldLog log.clf dosyasını tarihli bir zip dosyasına arşivler
func archiveOldLog(logFilePath, logDir string) {
	timestamp := time.Now().Format("20060102_1504")
	zipFileName := filepath.Join(logDir, fmt.Sprintf("log_%s.zip", timestamp))

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		fmt.Printf("Error creating zip file: %v", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	logFile, err := os.Open(logFilePath)
	if err != nil {
		fmt.Printf("Error opening log file: %v", err)
		return
	}
	defer logFile.Close()

	logWriter, err := zipWriter.Create("log.clf")
	if err != nil {
		fmt.Printf("Error writing to zip: %v", err)
		return
	}

	_, err = io.Copy(logWriter, logFile)
	if err != nil {
		fmt.Printf("Error copying to zip: %v", err)
		return
	}

	// Eski log dosyasını sıfırla
	_ = os.Truncate(logFilePath, 0)
}
