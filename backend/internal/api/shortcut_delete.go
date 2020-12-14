package api

import (
	"github.com/gin-gonic/gin"
	"firlus.dev/firl.us/internal/common"
)

// ShortcutDelete handles the endpoint DELETE /api/shortcuts/:path
func ShortcutDelete(c *gin.Context) {
	if !common.ValidatePassword(c.Request.Header["Authorization"][0]) {
		c.AbortWithStatus(403)
	}
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
