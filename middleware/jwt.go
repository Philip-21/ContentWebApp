package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/Philip-21/proj1/models"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("FDr1VjVQiSiybYJrQZNt8Vfd7bFEsKP6vNX1brOSiWl0mAIVCxJiR4/T3zpAlBKc2/9Lw2ac4IwMElGZkssfj3dqwa7CQC7IIB+nVxiM1c9yfowAZw4WQJ86RCUTXaXvRX8JoNYlgXcRrK3BK0E/fKCOY1+izInW3abf0jEeN40HJLkXG6MZnYdhzLnPgLL/TnIFTTAbbItxqWBtkz6FkZTG+dkDSXN7xNUxlg==")

type authClaims struct {
	jwt.StandardClaims
	ID uint `json:"Id"`
}

func GenerateJwt() (string, error) {
	var user models.ContentUser
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expiresAt,
		},
		ID: user.ID,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken() verifys the JWT token.
func VerifyJwt(tokenString string) (uint, string, error) {
	var claims authClaims
	//to decrypt the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	id := claims.ID
	email := claims.Subject
	return id, email, nil
}
