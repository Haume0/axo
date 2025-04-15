package auth

import (
	"axo/axo"
	"axo/database"
	"axo/models"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// User operations
func Auth(token string) (models.User, error) {
	return models.User{}, nil
}
func Login(email string, password string) (models.User, error) {
	//Check user.Mail with MailRegex
	if !axo.Unwrap(axo.RegexTest(email, models.MailRegex)) {
		return models.User{}, fmt.Errorf("BAD_MAIL_FORMAT")
	}

	//Check user.Password with PasswordRegex
	if !axo.Unwrap(axo.RegexTest(password, models.PasswordRegex)) {
		return models.User{}, fmt.Errorf("BAD_PASSWORD_FORMAT")
	}

	//sha256 hash the password
	sha := sha256.New()
	sha.Write([]byte(password))
	password = fmt.Sprintf("%x", sha.Sum(nil))

	var user models.User
	database.DB.Preload("Role").Where("email = ? AND password = ?", email, password).First(&user)
	if user.ID == 0 {
		return models.User{}, fmt.Errorf("USER_NOT_FOUND")
	}
	if !user.Active {
		return models.User{}, fmt.Errorf("USER_NOT_ACTIVE")
	}
	if !user.Verified {
		return models.User{}, fmt.Errorf("USER_NOT_VERIFIED")
	}
	return user, nil
}
func Register(user models.User) error {
	//Check user.Mail with MailRegex
	if !axo.Unwrap(axo.RegexTest(user.Email, models.MailRegex)) {
		return fmt.Errorf("BAD_MAIL_FORMAT")
	}

	//Check user.Password with PasswordRegex
	if !axo.Unwrap(axo.RegexTest(user.Password, models.PasswordRegex)) {
		return fmt.Errorf("BAD_PASSWORD_FORMAT")
	}

	//Checking if user exists,
	var searchUsers []models.User
	database.DB.Where("email = ?", user.Email).Find(&searchUsers)
	if len(searchUsers) > 0 {
		return fmt.Errorf("USER_ALREADY_EXISTS")
	}

	//sha256 hash the password
	sha := sha256.New()
	sha.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", sha.Sum(nil))

	//Creating user
	database.DB.Create(&user)
	return nil
}
func Refresh(ref_token string) (TokenResponse, error) {
	var user models.User
	var token models.RefreshToken
	database.DB.Where("token = ?", ref_token).First(&token)
	if token.ID == 0 {
		return TokenResponse{}, fmt.Errorf("TOKEN_NOT_FOUND")
	}
	database.DB.Where("id = ?", token.UserID).First(&user)
	if user.ID == 0 {
		return TokenResponse{}, fmt.Errorf("USER_NOT_FOUND")
	}
	if !user.Active {
		return TokenResponse{}, fmt.Errorf("USER_NOT_ACTIVE")
	}
	if !user.Verified {
		return TokenResponse{}, fmt.Errorf("USER_NOT_VERIFIED")
	}
	//Check if the token is expired
	if token.Exp.Before(time.Now()) {
		return TokenResponse{}, fmt.Errorf("TOKEN_EXPIRED")
	}
	accesToken, err := GenerateAccesToken(user)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("ERROR_GENERATING_ACCESS_TOKEN")
	}
	return accesToken, nil
}

// User operations
func CheckAuth() (bool, error) {
	// TODO: Will check if user is authenticated or not.
	return false, nil
}

func GetUserJWT(token string) (models.User, error) {
	userMap, err := axo.VerifyToken(
		os.Getenv("JWT_SECRET"),
		token,
	)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	database.DB.Preload("Role").Where("id = ?", userMap["id"]).First(&user)
	if user.ID == 0 {
		return models.User{}, fmt.Errorf("USER_NOT_FOUND")
	}
	return user, nil
}
func GetUser(r *http.Request) (models.User, error) {
	// Try to get token from cookie first
	if token, err := axo.GetCookie(r, "axo_auth_acc"); err == nil {
		user, err := GetUserJWT(token)
		if err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	// Fall back to authorization header
	authHeader := r.Header.Get("Authorization")
	var token string

	// Handle both "Bearer token" and plain token formats
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else if authHeader != "" {
		token = authHeader
	} else {
		return models.User{}, fmt.Errorf("NO_AUTH_TOKEN")
	}

	user, err := GetUserJWT(token)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
