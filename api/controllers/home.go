package controllers

import (
	"github.com/gin-gonic/gin"
)

// HomeIndex /api
func HomeIndex(c *gin.Context) {
	c.String(200, "Gin for API!")
}
