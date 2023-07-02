package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAdminBlocklist(c *gin.Context) {
	// Your custom logic goes here
	// ...

	// Send a response
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Admin Blocklist",
	})
}
