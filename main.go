package main

import (
	"app/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	r := gin.New()

	r.GET("/", controllers.HomeIndex)

	http.Handle("/", r)
}
