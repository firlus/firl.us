package api

import (
	"firlus.dev/firl.us/internal/model"
	"github.com/gin-gonic/gin"
)

// ShortcutPost handles the endpoint POST /api/shortcuts
func ShortcutPost(c *gin.Context) {
	path := c.PostForm("path")
	url := c.PostForm("url")
	// TODO validate parameters
	shortcut := model.Shortcut{
		Path: path,
		URL:  url}
	if Storage.ShortcutExists(path) {
		c.AbortWithStatus(409)
	} else if !shortcut.IsValid() {
		c.AbortWithStatus(400)
	} else {
		Storage.CreateShortcut(&shortcut)
		c.JSON(200, gin.H{})
	}
}
