package frontends

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// ServeStatic sets up a static file server for serving files from a specified directory.
// It maps a given URL path to a directory on the filesystem, allowing files to be accessed
// via the specified path.
//
// Parameters:
//
//	site  - *http.ServeMux: The HTTP router to handle incoming requests.
//	path  - string: The URL path that site will respond to.
//	dist  - string: The filesystem directory containing the static files to be served.
//	buildSteps - [0]: TargetFolder, [1]: BuildCommand, [2]: BuildCommand [...] etc.
//
// Example:
//
//	ServeStatic(mux, "/", "./public")
func ServeStatic(site *http.ServeMux, path string, dist string, buildSteps ...[]string) {
	if len(buildSteps) > 0 {
		fmt.Println("🚀 Building static app...")
		for i, step := range buildSteps[0] {
			if i == 0 {
				// First element is the target folder, skip execution
				continue
			}
			// Execute each build step
			fmt.Printf("Running: %s\n", step)
			cmd := exec.Command("sh", "-c", step)
			cmd.Dir = buildSteps[0][0] // Set the working directory to the target folder
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Build step output: %s", output)
				log.Fatalf("Error executing build step: %v", err)
			}
		}
	}
	site.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dist))))
}
