package api

import (
	"github.com/gin-gonic/gin"
)

// RedirectToShortcut handles the endpoint GET /:path
func RedirectToShortcut(c *gin.Context) {
	path := c.Param("path")
	shortcut, _ := Storage.ReadShortcut(path)
	if !shortcut.IsValid() {
		c.AbortWithStatus(404)
	} else {
		c.Redirect(301, shortcut.URL)
		c.Abort()
	}
}
