package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var log string

// Logger : is a simple middleware that logs incoming requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var currentTime = time.Now().Format("2006-01-02 15:04:05")
		var clientIPAdress = r.RemoteAddr
		var fullpath = r.URL.Path
		if r.URL.RawQuery != "" {
			fullpath = fullpath + "?" + r.URL.RawQuery
		}
		var logText = fmt.Sprintf("[AXO]:[%s] 👤|%s| 🚦[%s] 🔗 %s \n", currentTime, clientIPAdress, r.Method, fullpath)
		var print = fmt.Sprintf(
			"\033[36m[AXO]\033[0m \033[90m%s\033[0m | \033[33m%s\033[0m | 🚦 \033[32m%s\033[0m | \033[34m%s\033[0m\n",
			currentTime, clientIPAdress, r.Method, fullpath,
		)

		if os.Getenv("NOLOG") == "" {
			fmt.Print(print)
			log = fmt.Sprintf("%v\n%v", log, logText)
			// Write log to file
			var logFile, err = os.OpenFile("log.clf", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				fmt.Printf("Error opening file: %v", err)
			}
			defer logFile.Close()
			_, err = logFile.WriteString(logText)
			if err != nil {
				fmt.Printf("Error writing to file: %v", err)
			}

			//if log file size is bigger than 8MB, rename it log.clf.old
			fileInfo, _ := logFile.Stat()
			if fileInfo.Size() > 8*1024*1024 {
				os.Rename("log.clf", "log.clf.old")
			}
		}

		next.ServeHTTP(w, r)
	})
}
