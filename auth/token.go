package auth

import (
	"axo/axo"
	"axo/database"
	"axo/models"
	"os"
	"time"
)

// Token Operations
type TokenResponse struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

func ClearRefTokens(user models.User) {
	// Find all refresh tokens for this user
	var tokens []models.RefreshToken
	database.DB.Where("user_id = ?", user.ID).Find(&tokens)

	// Only delete expired tokens
	now := time.Now()
	for _, token := range tokens {
		if token.Exp.Before(now) {
			database.DB.Delete(&token)
		}
	}
}

func GenerateAccesToken(user models.User) (TokenResponse, error) {
	// 15min
	var exp time.Time = time.Now().Add(15 * time.Minute)
	var accesMap map[string]any = map[string]any{
		"id":      user.ID,
		"email":   user.Email,
		"role_id": user.RoleID,
		"role":    user.Role.Name,
	}
	token, err := axo.GenerateToken(
		os.Getenv("JWT_SECRET"),
		exp,
		accesMap,
	)
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		Token: token,
		Exp:   exp,
	}, nil
}

func GenerateRefreshToken(user models.User) (TokenResponse, error) {
	// 30 days
	var exp time.Time = time.Now().Add(30 * 24 * time.Hour)
	var refreshMap map[string]any = map[string]any{
		"id": user.ID,
	}
	token, err := axo.GenerateToken(
		os.Getenv("JWT_SECRET"),
		exp,
		refreshMap,
	)
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		Token: token,
		Exp:   exp,
	}, nil
}
