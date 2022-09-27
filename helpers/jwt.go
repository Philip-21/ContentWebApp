package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(email string, admin bool) (string, error) {
	claims := &jwtCustomClaims{
		email,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	//create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//generate token using secret key
	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	return t, nil

}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(SECRET_KEY), nil
	})

}
