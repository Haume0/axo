package routes

import (
	"axo/auth"
	"axo/axo"
	"axo/database"
	"axo/models"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	ref_token, err := axo.GetCookie(r, "axo_auth_ref")
	if err != nil {
		ref_token = r.URL.Query().Get("axo_auth_ref")
	}
	if ref_token == "" {
		axo.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}
	//Check if the token is valid
	claims, err := axo.VerifyToken(os.Getenv("JWT_SECRET"), ref_token)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accesToken, err := auth.Refresh(ref_token)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	if err := database.DB.Where("id = ?", claims["id"]).First(&user).Error; err != nil {
		axo.Error(w, "User not found", http.StatusNotFound)
		return
	}
	//Setting the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "axo_auth_acc",
		Value:    accesToken.Token,
		Path:     "/",
		Expires:  accesToken.Exp,
		MaxAge:   int(accesToken.Exp.Sub(time.Now()).Seconds()), // Calculate from time difference
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	})
	// End
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"access_token": accesToken.Token,
		"exp":          accesToken.Exp,
	})
}
