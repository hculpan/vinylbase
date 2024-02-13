package handlers

import (
	"log"
	"net/http"

	"github.com/hculpan/vinylbase/cmd/web/templates"
	"github.com/hculpan/vinylbase/pkg/entities"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Render the login page using Templ
	err := templates.LoginTemplate().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := entities.FetchUser(username)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	} else if user == nil || !user.ComparePassword(password) {
		log.Default().Printf("redirecting to error")
		http.Redirect(w, r, "/error?msg=Failed_to_login._Invalid_username_or_password.", http.StatusSeeOther)
		return
	}

	sessionManager.Put(r.Context(), "username", user.Username)
	sessionManager.Put(r.Context(), "realname", user.Realname)
	sessionManager.Put(r.Context(), "userid", user.Id())
	http.Redirect(w, r, "/mycollection", http.StatusSeeOther)
}
