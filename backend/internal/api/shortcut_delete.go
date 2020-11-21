package api

import (
	"github.com/gin-gonic/gin"
)

// ShortcutDelete handles the endpoint DELETE /api/shortcuts/:path
func ShortcutDelete(c *gin.Context) {
	path := c.Param("path")
	if !Storage.ShortcutExists(path) {
		c.AbortWithStatus(404)
	} else {
		err := Storage.DeleteShortcut(path)
		if err != nil {
			c.AbortWithStatus(500)
		} else {
			c.JSON(200, gin.H{})
		}
	}
}
