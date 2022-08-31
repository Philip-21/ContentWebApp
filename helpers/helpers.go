package helpers

import (
	"github.com/Philip-21/Content/config"
	"github.com/gin-gonic/gin"
)

var app *config.AppConfig

func IsAuthenticated(c *gin.Context) bool {
	exist := app.Session.Exists(c.Request.Context(), "email")
	return exist
}
