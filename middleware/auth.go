package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token, ok := GetToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	id, email, err := VerifyJwt(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	c.Set("id", id)
	c.Set("email", email)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}
func GetToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
func GetSession(c *gin.Context) (uint, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return 0, "", false
	}
	email, ok := c.Get("email")
	if !ok {
		return 0, "", false
	}
	return id.(uint), email.(string), true
}
