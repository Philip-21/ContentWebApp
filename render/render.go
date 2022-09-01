package render

import (
	"github.com/Philip-21/Content/config"
	"github.com/Philip-21/Content/models"
	"github.com/gin-gonic/gin"
)

var app *config.AppConfig

func AddData(td *models.TemplateData, c *gin.Context) *models.TemplateData {

	return td
}
