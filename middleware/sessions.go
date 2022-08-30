package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager

func SesssionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
