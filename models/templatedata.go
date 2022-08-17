package models

import (
	"github.com/Philip-21/proj1/forms"
	"github.com/Philip-21/proj1/render"
)

type TemplateData struct {
	IsAuthenticated int
	CSRFToken       string //a security token that handles forms
	Data            map[string]interface{}
	Form            *forms.Form
	Flash           *render.FlashData
	Warning         *render.WarningData
	Error           *render.ErrorData
}

// func AddData(td *TemplateData, c *gin.Context) *TemplateData {
// 	td.Flash = app.Session.PopString(c.Request.Context(), "flash")
// 	return td
// }
