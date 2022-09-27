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
