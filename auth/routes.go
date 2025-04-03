package auth

import (
	"axo/models"
	"net/http"
)

func LoginRoute(w http.ResponseWriter, r *http.Request) {}
func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var user models.User
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	if err := Register(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{
	"message": "User created successfully."}`))
}
func VerifyRoute(w http.ResponseWriter, r *http.Request)  {}
func LogoutRoute(w http.ResponseWriter, r *http.Request)  {}
func RefreshRoute(w http.ResponseWriter, r *http.Request) {}

// On mail verification i'm planing to hold verification token in lru cache.
