package auth

import (
	"axo/axo"
	"axo/database"
	"axo/flags"
	"axo/mail"
	"axo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

// Auth Routes
func LoginRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var email = r.FormValue("email")
	var password = r.FormValue("password")
	user, err := Login(email, password)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Creating the refresh token
	accesToken, err := GenerateAccesToken(user)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Clearing old tokens
	ClearRefTokens(user)
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
		"user":         user,
	})
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
func LogoutRoute(w http.ResponseWriter, r *http.Request) {
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
func RefreshRoute(w http.ResponseWriter, r *http.Request) {
	ref_token, err := axo.GetCookie(r, "axo_auth_ref")
	if err != nil {
		ref_token = r.URL.Query().Get("axo_auth_ref")
	}
	if ref_token == "" {
		axo.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}
	//Check if the token is valid
	_, err = axo.VerifyToken(os.Getenv("JWT_SECRET"), ref_token)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accesToken, err := Refresh(ref_token)
	if err != nil {
		axo.Error(w, err.Error(), http.StatusBadRequest)
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
	fmt.Fprint(w, accesToken.Token)
}

// Verification System
var verifications = expirable.NewLRU[string, string](0, nil, time.Minute*2)

func VerifyRoute(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")
	if email == "" || !axo.Unwrap(axo.RegexTest(email, models.MailRegex)) {
		axo.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		axo.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if user.Verified {
		axo.Error(w, "User already verified", http.StatusBadRequest)
		return
	}
	//?Verification code
	var sentCode = r.URL.Query().Get("code")
	if sentCode != "" {
		// Check if the code is valid
		if code, ok := verifications.Get(user.Email); ok {
			if sentCode == code {
				user.Verified = true
				database.DB.Save(&user)
				verifications.Remove(user.Email)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "User verified successfully",
				})
				return
			}
			axo.Error(w, "Invalid verification code", http.StatusBadRequest)
			return
		}
		axo.Error(w, "Verification code expired", http.StatusBadRequest)
		return
	}

	//[1] Generate a verification code
	code := axo.GenerateMemCode(4)
	//[2] Store the code in the cache
	verifications.Add(user.Email, code)
	//[3] Send the code to the user
	template, err := mail.LoadTemplate("verification_code")
	if err != nil {
		axo.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template = axo.MultiReplace(
		template,
		map[string]string{
			"{{.login_code}}":  code,
			"{{.base_url}}":    os.Getenv("BASE_URL"),
			"{{.title}}":       "Please verify your email.",
			"{{.description}}": "To verify your email, please enter the code below.",
			"{{.sub_text}}":    "This code is valid for 2 minutes.",
			"{{.warning}}":     "If you did not request this, please ignore this email.",
		},
	)
	mail.Send(
		user.Email,
		"Axo Verify",
		template,
		true,
	)
}
