package render

import (
	"html/template"
	"os"
)

func CreateTemplate() {
	const location = "./frontend"
	tmpl, err := template.ParseFiles(location)
	if err != nil {

		return
	}

	err = tmpl.Execute(os.Stdout, location)
	if err != nil {
		return
	}

}
