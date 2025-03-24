package routes

import (
	"axo/axo"
	"net/http"
)

func GetError(w http.ResponseWriter, r *http.Request) {
	axo.Error(w, "Hello, Error", http.StatusInternalServerError)
}
