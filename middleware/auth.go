package middleware

import (
	"fmt"
	"net/http"

	"github.com/Philip-21/Content/helpers"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
)

var session *scs.SessionManager

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header Provided ")})
			c.Abort()
			return
		}
		//validate token
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		// store a new key/value pair exclusively for this context
		c.Set("email", claims.Email)
		c.Next()

		if !helpers.IsAuthenticated(c) {
			session.Put(c.Request.Context(), "error", "log in first!")
			c.JSON(http.StatusBadRequest, gin.H{"error": "log in first"})
			http.Redirect(c.Writer, c.Request, "/login", http.StatusSeeOther)
			return

		}
		c.Next()

	}

}
