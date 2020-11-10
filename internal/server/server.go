package server

import (
	"firlus.dev/firl.us/internal/model"
	"firlus.dev/firl.us/internal/store"
	"github.com/gin-gonic/gin"
)

// Setup setup
func Setup(port int, store store.Store) {
	r := gin.Default()
	r.POST("/api/shortcuts", func(c *gin.Context) {
		path := c.PostForm("path")
		url := c.PostForm("url")
		// TODO validate parameters
		shortcut := model.Shortcut{
			Path: path,
			URL:  url}
		if store.ShortcutExists(path) {
			c.AbortWithStatus(409)
		} else if !shortcut.IsValid() {
			c.AbortWithStatus(400)
		} else {
			store.CreateShortcut(&shortcut)
			c.JSON(200, gin.H{})
		}
	})
	r.GET("/api/shortcuts/:path", func(c *gin.Context) {
		path := c.Param("path")
		shortcut, _ := store.ReadShortcut(path)
		if !shortcut.IsValid() {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, gin.H{
				"path": shortcut.Path,
				"url":  shortcut.URL})
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
