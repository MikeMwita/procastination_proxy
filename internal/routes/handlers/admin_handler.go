package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Block a domain

func (h Handler) PostAdminBlockDomain(c *gin.Context) {
	domain := c.Param("domain")
	err != h.ProxyService.BlockDomain(domain)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Domain '%s' blocked successfully", domain),
	})
}
