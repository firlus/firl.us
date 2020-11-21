package api

import (
	"firlus.dev/firl.us/internal/model"
	"github.com/gin-gonic/gin"
)

// ShortcutPut handles the endpoint PUT /api/shortcuts/:path
func ShortcutPut(c *gin.Context) {
	path := c.Param("path")
	url := c.PostForm("url")
	if !Storage.ShortcutExists(path) {
		c.AbortWithStatus(404)
	}
	shortcut := model.Shortcut{
		Path: path,
		URL:  url}
	if !shortcut.IsValid() {
		c.AbortWithStatus(400)
	} else {
		err := Storage.UpdateShortcut(&shortcut)
		if err != nil {
			c.AbortWithStatus(500)
		} else {
			c.JSON(200, gin.H{})
		}
	}
}
