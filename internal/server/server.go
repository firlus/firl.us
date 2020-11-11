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
	r.PUT("/api/shortcuts/:path", func(c *gin.Context) {
		path := c.Param("path")
		url := c.PostForm("url")
		if !store.ShortcutExists(path) {
			c.AbortWithStatus(404)
		}
		shortcut := model.Shortcut{
			Path: path,
			URL:  url}
		if !shortcut.IsValid() {
			c.AbortWithStatus(400)
		} else {
			err := store.UpdateShortcut(&shortcut)
			if err != nil {
				c.AbortWithStatus(500)
			} else {
				c.JSON(200, gin.H{})
			}
		}
	})
	r.DELETE("/api/shortcuts/:path", func(c *gin.Context) {
		path := c.Param("path")
		if !store.ShortcutExists(path) {
			c.AbortWithStatus(404)
		} else {
			err := store.DeleteShortcut(path)
			if err != nil {
				c.AbortWithStatus(500)
			} else {
				c.JSON(200, gin.H{})
			}
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
