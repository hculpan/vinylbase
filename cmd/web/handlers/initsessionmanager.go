package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

var sessionManager *scs.SessionManager

func InitSessionManager(r *chi.Mux) {
	// Initialize session manager
	sessionManager = scs.New()
	sessionManager.Cookie.Name = "sessionid"
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	// Middleware for session management
	r.Use(sessionManager.LoadAndSave)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the "username" property exists in the session
		username := sessionManager.GetString(r.Context(), "username")
		if username == "" {
			// User is not authorized, redirect to the error page with a message
			http.Redirect(w, r, "/error?msg=unauthorized", http.StatusFound)
			return
		}

		// User is authorized, proceed with the request
		next.ServeHTTP(w, r)
	})
}
