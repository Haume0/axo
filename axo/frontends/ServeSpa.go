package frontends

import (
	"axo/axo"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// Utility function for serving Single Page Applications (SPA).
// It will run development server and reverse proxy for the SPA in development.
// For production, it will serve the static files from the dist directory.

// SPA Handler for client-side routing
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// ServeSPA sets up a Single Page Application (SPA) server with support for both
// development and production modes. In development mode, it starts a development
// server (e.g., Vite) and proxies requests to it. In production mode, it serves
// static files from a specified distribution directory.
//
// Parameters:
//
//	router     - *http.ServeMux: The HTTP router to handle incoming requests.
//	route      - string: The URL path prefix for the SPA (e.g., "/").
//	devCommand - string: The command to start the development server (e.g., "npm run dev").
//	port       - string: The port number for the development server to listen on.
//	sitePath   - string: The root directory of the project, used to run the development server.
//	distPath   - string: The directory containing the production build of the SPA.
func ServeSPA(router *http.ServeMux, route string, port string, sitePath string, distPath string, devCommand [2]string, buildCommands []string) {
	// Website Route
	// Check if in production mode
	if os.Args[len(os.Args)-1] == "--prod" {
		// Production mode
		// Run build commands if provided
		if len(buildCommands) > 0 {
			fmt.Println("Building production assets...")
			for _, cmd := range buildCommands {
				buildCmd := exec.Command("sh", "-c", cmd)
				buildCmd.Dir = sitePath
				buildCmd.Stdout = os.Stdout
				buildCmd.Stderr = os.Stderr

				fmt.Printf("Running: %s\n", cmd)
				if err := buildCmd.Run(); err != nil {
					log.Fatalf("Build command failed: %v", err)
				}
			}
			fmt.Println("Build completed successfully")
		}

		// SPA HANDLER for production mode
		spa := spaHandler{staticPath: distPath, indexPath: "index.html"}
		router.Handle(route, spa)
	} else {
		// Development mode starting vite development server and proxying to it
		// Check if node_modules exist
		if _, err := os.Stat(filepath.Join(sitePath, "node_modules")); os.IsNotExist(err) {
			// If not, run the install command
			installCmd := exec.Command("sh", "-c", devCommand[0])
			installCmd.Dir = sitePath
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			fmt.Println("Running: " + devCommand[0])
			if err := installCmd.Run(); err != nil {
				log.Fatalf("Install command failed: %v", err)
			}
			fmt.Println("Install completed successfully")
		}
		go func() {
			cmd := exec.Command("sh", "-c", fmt.Sprintf("%s -- --port %s", devCommand[1], port))
			// Change directory to the root of the project
			cmd.Dir = sitePath
			if os.Getenv("verbose") == "true" {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}
		}()
		// Proxy to vite development server
		router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			// Proxy to vite development server on specified port
			proxyURL := "http://localhost:" + port
			//reverse proxy to vite development server
			axo.ReverseProxy(w, r, proxyURL)
		})
	}
}
