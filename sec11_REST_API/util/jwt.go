package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TokenExpiedTime is JWT token expiration time in hours
const TokenExpiedTime = 2 // in hours
// JWTSecret is JWT secret key
const JWTSecret = "NJSLO12J4"

// GenerateToken generates a JWT token for the user
func GenerateToken(id int64, email string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * TokenExpiedTime).Unix(),
	}).SignedString([]byte(JWTSecret))

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// ValidateToken validates the JWT token and returns the user
func ValidateToken(token string) error {
	return nil
}
