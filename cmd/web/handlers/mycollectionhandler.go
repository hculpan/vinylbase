package handlers

import (
	"net/http"

	"github.com/hculpan/vinylbase/cmd/web/templates"
)

func MyCollectionHandler(w http.ResponseWriter, r *http.Request) {
	username := sessionManager.GetString(r.Context(), "username")
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Render the hello page with the username using Templ
	err := templates.HelloTemplate(username).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
