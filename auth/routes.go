package auth

import (
	"axo/axo"
	"axo/database"
	"axo/models"
	"encoding/json"
	"net/http"
)

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var email = r.FormValue("email")
	var password = r.FormValue("password")
	user, err := Login(email, password)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var user models.User
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	if err := Register(user); err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var createdUser models.User
	if err := database.DB.Preload("Role").Where("email = ?", user.Email).First(&createdUser).Error; err != nil {
		axo.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
func VerifyRoute(w http.ResponseWriter, r *http.Request)  {}
func LogoutRoute(w http.ResponseWriter, r *http.Request)  {}
func RefreshRoute(w http.ResponseWriter, r *http.Request) {}
