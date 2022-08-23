package render

import (
	"github.com/Philip-21/proj1/config"
	"github.com/Philip-21/proj1/models"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

var app *config.AppConfig

func AddData(td *models.TemplateData, c *gin.Context) *models.TemplateData {
	td.Flash = app.Session.PopString(c.Request.Context(), "flash")
	td.Error = app.Session.PopString(c.Request.Context(), "error")
	td.Warning = app.Session.PopString(c.Request.Context(), "warning")
	td.CSRFToken = csrf.GetToken(c)
	if app.Session.Exists(c.Request.Context(), "user_id") {
		td.IsAuthenticated = 1
	}

	return td
}

type Data struct {
	flash           string
	errors          string
	warning         string
	csrf            string
	isAuthenticated int
}

// type FlashData struct{}
// type ErrorData struct{}
// type WarningData struct{}

func Flash(d *Data, c *gin.Context) *Data {
	d.flash = app.Session.PopString(c.Request.Context(), "flash")
	return d

}

func Error(d *Data, c *gin.Context) *Data {
	d.errors = app.Session.PopString(c.Request.Context(), "error")
	return d
}

func Warning(d *Data, c *gin.Context) *Data {
	d.warning = app.Session.PopString(c.Request.Context(), "warning")
	return d
}

func Csrf(d *Data, c *gin.Context) {
	d.csrf = csrf.GetToken(c)

}
func IsAuth(d *Data, c *gin.Context) {
	if app.Session.Exists(c.Request.Context(), "user_id") {
		d.isAuthenticated = 1
	}
}

type FlashData interface {
	Flash(d *Data, c *gin.Context) *Data
}

type ErrorData interface {
	Error(d *Data, c *gin.Context) *Data
}
type WarningData interface {
	Warning(d *Data, c *gin.Context) *Data
}

type CSRFToken interface {
	Csrf(d *Data, c *gin.Context)
}

type IsAuthenticate interface {
	IsAuth(d *Data, c *gin.Context)
}
