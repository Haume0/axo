package routes

import (
	"axo/axo"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

func GetError(w http.ResponseWriter, r *http.Request) {
	axo.Error(w, "Hello, Error", http.StatusInternalServerError)
}
