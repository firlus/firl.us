package api

import (
	"fmt"
	"net/http"
	"strings"

	"firlus.dev/firl.us/internal/common"
)

// ShortcutDelete handles the endpoint DELETE /api/shortcuts/:path
func ShortcutDelete(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("Authorization")
	if !common.ValidatePassword(password) {
		http.Error(w, "Wrong password", http.StatusForbidden)
	} else {
		fmt.Println("Something else")
		pathSlice := strings.Split(r.URL.Path, "/")
		path := pathSlice[len(pathSlice)-1]

		if !Storage.ShortcutExists(path) {
			http.Error(w, "Shortcut does not exist", http.StatusNotFound)
		} else {
			err := Storage.DeleteShortcut(path)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}
	}
}
