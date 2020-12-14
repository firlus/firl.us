package api

import (
	"github.com/gin-gonic/gin"
	"firlus.dev/firl.us/internal/common"
)

// ShortcutGet handles the endpoint GET /api/shortcuts/:path
func ShortcutGet(c *gin.Context) {
	if !common.ValidatePassword(c.Request.Header["Authorization"][0]) {
		c.AbortWithStatus(403)
	}
	path := c.Param("path")
	shortcut, _ := Storage.ReadShortcut(path)
	if !shortcut.IsValid() && c.Param("api") == "api" && c.Param("shortcuts") == "shortcuts" { // Gin's router is shit so this mess is necessary
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, gin.H{
			"path": shortcut.Path,
			"url":  shortcut.URL})
	}

}
