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
func GenerateToken(userID int64, email string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * TokenExpiedTime).Unix(),
	}).SignedString([]byte(JWTSecret))

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// VerifyToken validates the JWT token and returns the user
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}
	if !parsedToken.Valid {
		return 0, fmt.Errorf("Token is invalid")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to parse claims")
	}
	if claims["exp"] == nil {
		return 0, fmt.Errorf("token is invalid")
	}
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return 0, fmt.Errorf("token is expired")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil
}
