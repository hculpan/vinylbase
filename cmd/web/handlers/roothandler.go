package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	username := sessionManager.GetString(r.Context(), "username")
	if username == "" {
		// User is not authorized, redirect to the error page with a message
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/mycollection", http.StatusSeeOther)
}
