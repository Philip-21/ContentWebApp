package models

import (
	"github.com/Philip-21/Content/forms"
)

type TemplateData struct {
	CSRFToken       string
	SessionData     map[string]interface{}
	User            map[string]string
	Form            *forms.Form
	Warning         any
	Error           map[string]interface{}
	Message         map[string]interface{}
	IsAuthenticated int
	StringMap       map[string]string
}

// func AddData(td *TemplateData, c *gin.Context) *TemplateData {
// 	td.Flash = app.Session.PopString(c.Request.Context(), "flash")
// 	return td
// }
