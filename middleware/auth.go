package middleware

import (
	"github.com/Philip-21/Content/helpers"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader

		helpers.ValidateToken(tokenString)

		//helpers.IsAuthenticted(&gin.Context{})
	}
}
