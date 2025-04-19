package routes

import (
	"axo/auth"
	"axo/axo"
	"axo/database"
	"axo/models"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var user models.User
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	if err := auth.Register(user); err != nil {
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
