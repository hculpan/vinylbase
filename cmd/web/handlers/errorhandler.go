package handlers

import (
	"net/http"
	"strings"

	"github.com/hculpan/vinylbase/cmd/web/templates"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	errorMsg := r.URL.Query().Get("msg")
	errorMsg = strings.ReplaceAll(errorMsg, "_", " ")

	err := templates.ErrorTemplate(errorMsg).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
