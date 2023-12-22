package middleware

import (
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validUsername := "super_user"
		validPassword := "admin_123"

		username, password, ok := r.BasicAuth()

		if ok && username == validUsername && password == validPassword {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
