package axo

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

// CookieValue : handles getting cookie value and returning base of key
func CookieValue(r *http.Request, key string) string {
	cookie, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	val, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		println(err.Error())
		return ""
	}
	return val
}

// GetLanguage : returns language from header
func GetLanguage(r *http.Request, first ...bool) any {
	// Determine the default behavior for 'first'
	useFirst := true
	if len(first) > 0 {
		useFirst = first[0]
	}

	// Check cookie
	cookie := CookieValue(r, "accept-language")
	if cookie != "" {
		if useFirst {
			return strings.Split(cookie, ",")[0]
		}
		return strings.Split(cookie, ",")
	}

	// Check header
	language := r.Header.Get("Accept-Language")
	languages := strings.Split(language, ",")
	if useFirst {
		lang := languages[0]
		if strings.Contains(lang, "-") {
			lang = strings.Split(lang, "-")[0]
		}
		return lang
	}

	// Return the full list of languages
	for i, lang := range languages {
		if strings.Contains(lang, "-") {
			languages[i] = strings.Split(lang, "-")[0]
		}
	}
	return languages
}
