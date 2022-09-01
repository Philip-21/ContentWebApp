package helpers

import (
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) bool {
	exist := app.Session.Exists(c.Request.Context(), "email")
	return exist
}
