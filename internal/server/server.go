package server

import (
	"firlus.dev/firl.us/internal/api"
	"firlus.dev/firl.us/internal/store"
	"github.com/gin-gonic/gin"
)

// Storage instance to manipulate databse

// Setup setup
func Setup(port string, storage store.Store) {
	api.Storage = storage
	router := gin.Default()
	apiRoute := router.Group("/api")
	router.GET("/:api", api.RedirectToShortcut)
	apiRoute.POST("/shortcuts", api.ShortcutPost)
	apiRoute.GET("/shortcuts", api.AllShortcutsGet)

	// This route is different (due to Gin's annoying router.)
	// Should look like that:
	// api.GET("/shortcuts/:path", func(c *gin.Context) {
	router.GET("/:api/:shortcuts/:path", api.ShortcutGet)
	apiRoute.PUT("/shortcuts/:path", api.ShortcutPut)
	apiRoute.DELETE("/shortcuts/:path", api.ShortcutDelete)

	router.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
