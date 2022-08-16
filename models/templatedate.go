package models

type TemplateData struct {
	IsAuthenticated int
	CSRFToken       string //a security token that handles forms
}
