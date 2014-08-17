package main

import (
	"net/http"

	"github.com/anoside/anoside/api/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()

	r.GET("/api", controllers.HomeIndex)

	http.Handle("/", r)
}
