package helpers

import (
	"net/http"

	"github.com/Philip-21/Content/config"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var session *scs.SessionManager
var app *config.AppConfig

// provides middleware which automatically loads and saves session data for the current request,
// and communicates the session token to and from the client in a cookie
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{ //cookie generated to identify user when they visit a page or website
		HttpOnly: true,
		Path:     "/",              //cookie path which applies to the entire site
		Secure:   app.InProduction, //the app.InProduction refers to the variable defined in the main package,     production use true but in development set it to false
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
