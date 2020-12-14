package api

import (
	"firlus.dev/firl.us/internal/model"
	"github.com/gin-gonic/gin"
	"firlus.dev/firl.us/internal/common"
)

// ShortcutPut handles the endpoint PUT /api/shortcuts/:path
func ShortcutPut(c *gin.Context) {
	if common.ValidatePassword(c.Request.Header["Authorization"][0]) {
		c.AbortWithStatus(403)
	}
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
