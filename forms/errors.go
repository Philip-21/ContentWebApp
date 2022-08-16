package forms

type errors map[string][]string

//Add error message to the given form field
func (e errors) Add(field, message string) {
	//putting the error message in an array of strings
	e[field] = append(e[field], message)
}

//Get the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
