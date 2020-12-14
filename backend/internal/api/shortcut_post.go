package api

import (
	"firlus.dev/firl.us/internal/model"
	"github.com/gin-gonic/gin"
	"firlus.dev/firl.us/internal/common"
)

// ShortcutPost handles the endpoint POST /api/shortcuts
func ShortcutPost(c *gin.Context) {
	if !common.ValidatePassword(c.Request.Header["Authorization"][0]) {
		c.AbortWithStatus(403)
	}
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
