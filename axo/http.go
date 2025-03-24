package axo

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

// ReverseProxy is the reverse proxy for the API
func ReverseProxy(w http.ResponseWriter, r *http.Request, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("URL parse error: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	// ReverseProxy
	proxy.ServeHTTP(w, r)
}
