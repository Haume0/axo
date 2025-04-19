package routes

import (
	"axo/auth"
	"axo/axo"
	"axo/database"
	"axo/flags"
	"axo/models"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var email = r.FormValue("email")
	var password = r.FormValue("password")
	user, err := auth.Login(email, password)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Creating the refresh token
	accesToken, err := auth.GenerateAccesToken(user)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := auth.GenerateRefreshToken(user)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Clearing old tokens
	auth.ClearRefTokens(user)
	//Saving the refresh token to the database
	database.DB.Create(&models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken.Token,
		Exp:       refreshToken.Exp,
		CreatedAt: time.Now(),
	})

	//Setting the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "axo_auth_ref",
		Value:    refreshToken.Token,
		Path:     "/",
		Expires:  refreshToken.Exp,
		MaxAge:   int(refreshToken.Exp.Sub(time.Now()).Seconds()), // Calculate from time difference
		Secure:   *flags.IsProduction,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"access_token": accesToken.Token,
		"exp":          accesToken.Exp,
		"user":         user,
	})
}
