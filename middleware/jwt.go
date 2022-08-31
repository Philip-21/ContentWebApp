package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Philip-21/Content/config"
	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

var connect *config.Envconfig

// var SECRET_KEY = fmt.Sprintf("SECRET_KEY=%s", connect.SecretKey)
var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(email string) (signedToken string, signedRefreshToken string, err error) {
	//generate new token
	claims := &SignedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	//gets a new token if initial token has expired
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	//call the jwt
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

// confirms the token to be used in the middlewre
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("invalid token ")
		msg = err.Error()
		return
	}
	//additinal checks
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
