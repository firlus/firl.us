package api

import (
	"github.com/gin-gonic/gin"
	"firlus.dev/firl.us/internal/common"
)

// AllShortcutsGet handles the endpoint GET /api/shortcuts
func AllShortcutsGet(c *gin.Context) {
	if !common.ValidatePassword(c.Request.Header["Authorization"][0]) {
		c.AbortWithStatus(403)
	}
	shortcuts, _ := Storage.ReadShortcuts()
	c.JSON(200, shortcuts) 
}
