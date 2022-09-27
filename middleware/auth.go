package middleware

import (
	"log"
	"net/http"

	"github.com/Philip-21/Content/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := helpers.ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[email]: ", claims["email"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
