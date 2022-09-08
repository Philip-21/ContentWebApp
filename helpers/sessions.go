package helpers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// name of the cookie
const sessionName = "session-Cookie"

func GetCookieStore() *sessions.CookieStore {
	sessionKey := os.Getenv("SESSION_KEY") //a random value to match a key in te session
	return sessions.NewCookieStore([]byte(sessionKey))
}

// used as a middleware for the handlers
// func GetCookie(c *gin.Context) {
// 	session, _ := GetCookieStore().Get(c.Request, sessionName)
// 	_, ok := session.Values["user"]
// 	if !ok {
// 		c.Abort()
// 		return

// 	}
// 	session.Save(c.Request, c.Writer)
// 	c.Next()
// }

// Set adds a new message into the cookie storage.
func SetFlash(c *gin.Context, name, value string) {
	session, _ := GetCookieStore().Get(c.Request, sessionName)

	session.AddFlash(name, value)
	session.Save(c.Request, c.Writer)
}

// Get gets flash messages from the cookie storage.
func GetFlash(c *gin.Context, name string) []string {
	session, _ := GetCookieStore().Get(c.Request, sessionName)

	flashMessage := session.Flashes(name)
	//if we have some messages
	if len(flashMessage) > 0 {
		session.Save(c.Request, c.Writer)

		//string slice to return messages

		var flashes []string
		for _, f := range flashMessage {
			///add Message to slice
			flashes = append(flashes, f.(string))
		}
		return flashes
	}
	return nil
}
