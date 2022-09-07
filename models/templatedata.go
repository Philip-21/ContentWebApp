package models

import (
	"github.com/Philip-21/Content/forms"
)

type TemplateData struct {
	CSRFToken       string
	Data            map[string]interface{}
	Form            *forms.Form
	Warning         string
	Error           string
	Flash           any
	IsAuthenticated int
	StringMap       map[string]string
}

// func AddData(td *TemplateData, c *gin.Context) *TemplateData {
// 	td.Flash = app.Session.PopString(c.Request.Context(), "flash")
// 	return td
// }
