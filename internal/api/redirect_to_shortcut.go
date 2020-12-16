package api

import (
	"net/http"
	"strings"
)

// RedirectToShortcut handles the endpoint GET /:path
func RedirectToShortcut(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	shortcut, _ := Storage.ReadShortcut(path)
	if !shortcut.IsValid() {
		http.Error(w, "Shortcut not found", http.StatusNotFound)
	} else {
		http.Redirect(w, r, shortcut.URL, http.StatusPermanentRedirect)
	}
}
