package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"firlus.dev/firl.us/internal/api"
	"firlus.dev/firl.us/internal/store"
	"github.com/gorilla/mux"
)

// Storage instance to manipulate databse

// Setup setup
func Setup(port string, storage store.Store) {
	api.Storage = storage

	router := mux.NewRouter()

	// Routes: Resolve Shortcut
	router.Methods("GET").Path("/:shortcut").HandlerFunc(api.RedirectToShortcut)

	// Routes: API
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Methods("GET").Path("/shortcuts").HandlerFunc(api.AllShortcutsGet)
	apiRouter.Methods("POST").Path("/shortcuts").HandlerFunc(api.ShortcutPost)
	apiRouter.Methods("GET").Path("/shortcuts/{path}").HandlerFunc(api.ShortcutGet)
	apiRouter.Methods("PUT").Path("/shortcuts/{path}").HandlerFunc(api.ShortcutPut)
	apiRouter.Methods("DELETE").Path("/shortcuts/{path}").HandlerFunc(api.ShortcutDelete)

	addr := fmt.Sprintf("0.0.0.0:%v", port)

	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
