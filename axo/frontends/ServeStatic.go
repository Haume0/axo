package frontends

import "net/http"

// ServeStatic sets up a static file server for serving files from a specified directory.
// It maps a given URL path to a directory on the filesystem, allowing files to be accessed
// via the specified path.
//
// Parameters:
//
//	site  - *http.ServeMux: The HTTP router to handle incoming requests.
//	path  - string: The URL path that site will respond to.
//	dist  - string: The filesystem directory containing the static files to be served.
//
// Example:
//
//	ServeStatic(mux, "/", "./public")
func ServeStatic(site *http.ServeMux, path string, dist string) {
	site.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dist))))

}
