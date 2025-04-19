package routes

import (
	"axo/axo"
	"axo/database"
	"axo/mail"
	"axo/models"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

var resetCodes = expirable.NewLRU[string, string](0, nil, time.Minute*2)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var email = r.FormValue("email")
	if email == "" {
		axo.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if !axo.Unwrap(axo.RegexTest(email, models.MailRegex)) {
		axo.Error(w, "BAD_MAIL_FORMAT", http.StatusBadRequest)
		return
	}

	// Get verification code from query parameters
	var sentCode = r.URL.Query().Get("code")
	var newPassword = r.FormValue("new_password")

	if sentCode != "" {
		// Verify password complexity
		//Check user.Password with PasswordRegex
		if !axo.Unwrap(axo.RegexTest(newPassword, models.PasswordRegex)) {
			axo.Error(w, "BAD_PASSWORD_FORMAT", http.StatusBadRequest)
			return
		}

		// Check if the code is valid
		if code, ok := resetCodes.Get(email); ok {
			if sentCode == code {
				var user models.User
				if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
					axo.Error(w, "User not found", http.StatusNotFound)
					return
				}

				// Hash the password securely
				sha := sha256.New()
				sha.Write([]byte(newPassword))
				user.Password = fmt.Sprintf("%x", sha.Sum(nil))

				// Save the new password
				if err := database.DB.Save(user).Error; err != nil {
					axo.Error(w, "Failed to update password", http.StatusInternalServerError)
					return
				}

				// Successfully updated password
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message": "Password updated successfully"}`))
				return
			}
			axo.Error(w, "Invalid code", http.StatusBadRequest)
			return
		}
		axo.Error(w, "Code expired", http.StatusBadRequest)
		return
	}

	// Generate a verification code (consider making this longer than 4 digits)
	code := axo.GenerateMemCode(6) // Increased from 4 to 6 digits for better security

	// Store the code in the cache
	resetCodes.Add(email, code)

	// Rate limit check could be added here

	// Send the code to the user
	template, err := mail.LoadTemplate("verification_code")
	if err != nil {
		axo.Error(w, "Internal server error", http.StatusInternalServerError) // Don't expose specific error details
		return
	}

	template = axo.MultiReplace(
		template,
		map[string]string{
			"{{.login_code}}":  code,
			"{{.base_url}}":    os.Getenv("BASE_URL"),
			"{{.title}}":       "Axo Recover",
			"{{.description}}": "Please use this code to reset your password.",
			"{{.sub_text}}":    "",
			"{{.warning}}":     "If you did not request this, please ignore this email.",
		},
	)

	mail.Send(
		email,
		"Axo Recover",
		template,
		true,
	)

	// Return success without exposing whether the email exists in the system
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Verification code sent"}`))
}
