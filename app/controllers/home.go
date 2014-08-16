package controllers

import (
  "github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
  c.String(200, "gin!")
}
