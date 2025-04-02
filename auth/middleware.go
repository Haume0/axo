package auth

import "net/http"

// Middleware is a middleware function that checks if the user is authenticated.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: will be implemented in the future
		// Check if the user is authenticated
	})
}
