package api

import (
	"net/http"

	"firlus.dev/firl.us/internal/common"
	"firlus.dev/firl.us/internal/model"
)

// ShortcutPost handles the endpoint POST /api/shortcuts
func ShortcutPost(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("Authorization")
	if !common.ValidatePassword(password) {
		http.Error(w, "Wrong password", http.StatusForbidden)
	} else {
		path := r.FormValue("path")
		url := r.FormValue("url")
		shortcut := model.Shortcut{
			Path: path,
			URL:  url}
		if Storage.ShortcutExists(path) {
			http.Error(w, "Shortcut already exists", http.StatusConflict)
		} else if !shortcut.IsValid() {
			http.Error(w, "Invalid shortcut", http.StatusBadRequest)
		} else {
			Storage.CreateShortcut(&shortcut)
		}
	}
}
