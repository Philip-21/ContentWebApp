package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/Philip-21/Content/config"
	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email     string
	Uid       string
	User_type string
	jwt.StandardClaims
}

var connect *config.Envconfig

var SECRET_KEY = fmt.Sprintf("SECRET_KEY=%s", connect.SecretKey)

func GetAllToken(email string, uid string, user_type string) (signedToken string, signedRefreshToken string, err error) {
	//generate new token
	claims := &SignedDetails{
		Email:     email,
		Uid:       uid,
		User_type: user_type,
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
