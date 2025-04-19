package routes

import (
	"axo/axo"
	"axo/database"
	"axo/flags"
	"axo/models"
	"encoding/json"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	//remove auth tokens
	ref_token, err := axo.GetCookie(r, "axo_auth_ref")
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Deleting the refresh token from the database
	var token models.RefreshToken
	if err := database.DB.Where("token = ?", ref_token).First(&token).Error; err != nil {
		axo.Error(w, "Token not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&token)
	//Clearing the cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "axo_auth_ref",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   *flags.IsProduction,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "axo_auth_acc",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Logged out successfully",
	})
}
