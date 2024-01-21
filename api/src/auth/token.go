package auth

import (
	"errors"
	"fmt"
	"net/http"
	"social-car/src/config"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken generates a token for the given user ID.
//
// Parameters:
// - userId: The ID of the user.
//
// Returns:
// - string: The generated token.
// - error: An error if the token generation fails.
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken validates the token extracted from the given HTTP request.
//
// It takes in a *http.Request object as a parameter and returns an error.
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Invalid token")
}

// GetUserID retrieves the user ID from the given HTTP request.
//
// It takes a pointer to an http.Request as a parameter and returns the user ID as a uint64 and an error.
func GetUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint64(permissions["userId"].(float64))
		return userID, nil
	}
	return 0, errors.New("Invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected Signing method: %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
