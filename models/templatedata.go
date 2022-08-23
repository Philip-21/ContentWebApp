package models

import (
	"github.com/Philip-21/proj1/forms"
)

type TemplateData struct {
	CSRFToken       string
	Data            map[string]interface{}
	Form            *forms.Form
	Warning         string
	Error           string
	Flash           string
	IsAuthenticated int
}

// func AddData(td *TemplateData, c *gin.Context) *TemplateData {
// 	td.Flash = app.Session.PopString(c.Request.Context(), "flash")
// 	return td
// }
