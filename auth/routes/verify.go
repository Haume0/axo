package routes

import (
	"axo/axo"
	"axo/database"
	"axo/mail"
	"axo/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

var verifications = expirable.NewLRU[string, string](0, nil, time.Minute*2)

func Verify(w http.ResponseWriter, r *http.Request) {
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
