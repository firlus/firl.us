package api

import (
	"net/http"
	"strings"

	"firlus.dev/firl.us/internal/common"
	"firlus.dev/firl.us/internal/model"
)

// ShortcutPut handles the endpoint PUT /api/shortcuts/:path
func ShortcutPut(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("Authorization")
	if !common.ValidatePassword(password) {
		http.Error(w, "Wrong password", http.StatusForbidden)
	} else {
		pathSlice := strings.Split(r.URL.Path, "/")
		path := pathSlice[len(pathSlice)-1]
		url := r.FormValue("url")

		if !Storage.ShortcutExists(path) {
			http.Error(w, "Shortcut not found", http.StatusNotFound)
		} else {
			shortcut := model.Shortcut{
				Path: path,
				URL:  url}
			if !shortcut.IsValid() {
				http.Error(w, "Invalid shortcut", http.StatusBadRequest)
			} else {
				err := Storage.UpdateShortcut(&shortcut)
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}

			}
		}
	}
}
