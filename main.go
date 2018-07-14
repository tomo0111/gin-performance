package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo0111/gin-performance/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.HelloController)
	r.Run(":8080")
}

