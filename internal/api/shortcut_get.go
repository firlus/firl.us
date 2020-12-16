package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"firlus.dev/firl.us/internal/common"
)

// ShortcutGet handles the endpoint GET /api/shortcuts/:path
func ShortcutGet(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("Authorization")
	if !common.ValidatePassword(password) {
		http.Error(w, "Wrong password", http.StatusForbidden)
	} else {
		pathSlice := strings.Split(r.URL.Path, "/")
		path := pathSlice[len(pathSlice)-1]

		shortcut, _ := Storage.ReadShortcut(path)
		if !shortcut.IsValid() { // Gin's router is shit so this mess is necessary
			http.Error(w, "Shortcut does not exist", http.StatusNotFound)
		} else {
			shortcutBytestream, _ := json.Marshal(shortcut)
			w.Write(shortcutBytestream)
		}
	}
}
