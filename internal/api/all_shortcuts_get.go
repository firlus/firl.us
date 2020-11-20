package api

import (
	"github.com/gin-gonic/gin"
)

// AllShortcutsGet handles the endpoint GET /api/shortcuts
func AllShortcutsGet(c *gin.Context) {
	shortcuts, _ := Storage.ReadShortcuts()
	c.JSON(200, shortcuts) 
}
