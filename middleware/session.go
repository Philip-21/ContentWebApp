package middleware

import (
	"net/http"

	"github.com/Philip-21/proj1/config"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
)

var session *scs.SessionManager
var app *config.AppConfig

//adding a middleware that tells the webserver to remember a state using sessions
// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

//Allows a particular User to Access a certain page
func IsAuthenticated(c *gin.Context) bool {
	exist := app.Session.Exists(c.Request.Context(), "user_id")
	return exist
}

//makes routes secure by only allowing logged in users to have access to certain parts,pages(routes) of the application
func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if !IsAuthenticated(c) {
			session.Put(c.Request.Context(), "error", "Log in First!") //puts the request in a context and write an error
			c.Redirect(http.StatusSeeOther, "user/login")
			return
		}
		c.Next()
	})

}
