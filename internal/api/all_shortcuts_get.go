package api

import (
	"encoding/json"
	"net/http"

	"firlus.dev/firl.us/internal/common"
)

// AllShortcutsGet handles the endpoint GET /api/shortcuts
func AllShortcutsGet(w http.ResponseWriter, r *http.Request) {
	password := r.Header.Get("Authorization")
	if !common.ValidatePassword(password) {
		http.Error(w, "Wrong password", http.StatusForbidden)
	} else {
		shortcuts, _ := Storage.ReadShortcuts()
		shortcutsBytestream, _ := json.Marshal(shortcuts)
		w.Write(shortcutsBytestream)
	}
}
