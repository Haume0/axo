package axo

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

// Error : makes default http error response prettier
func Error(w http.ResponseWriter, message string, status int) {
	// Creating JSON response
	res, err := json.Marshal(errorResponse{Error: message})
	if err != nil {
		http.Error(w, "error marshalling data", http.StatusInternalServerError)
		return
	}

	// Setting header
	w.Header().Set("Content-Type", "application/json")

	// Writing response
	w.WriteHeader(status)
	w.Write(res)
}

// StaticFileHandler serves static files from the given directory
func StaticFileHandler(directory string) http.Handler {
	return http.StripPrefix("/", http.FileServer(http.Dir(directory)))
}
