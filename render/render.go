package render

import (
	"github.com/Philip-21/proj1/config"
	"github.com/gin-gonic/gin"
)

var app *config.AppConfig

type Data struct {
	flash   string
	errors  string
	warning string
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

type FlashData interface {
	Flash(d *Data, c *gin.Context) *Data
}

type ErrorData interface {
	Error(d *Data, c *gin.Context) *Data
}
type WarningData interface {
	Warning(d *Data, c *gin.Context) *Data
}
