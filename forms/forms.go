package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	//url.Values maps a string key to a list of values,
	//for query parameters and form values
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//create a new form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for required fields
func (f *Form) Required(fields ...string) { //n.b (... called ellipses) denotes range of items or argument
	for _, field := range fields {
		//Get(url.values) :gets the first value associated with the given key.
		//If there are no values associated with the key,
		// Get returns the empty string
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field Cannot be Blanl")
		}
	}
}

//MinLength checks for string minimum length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length { //if len is < the actual lenght of the field
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

//Enail validation
func (f *Form) ValidEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) { //url.Values
		f.Errors.Add(field, "Invalid email Address")
	}
}
