package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo0111/gin-performance/controller"
	"github.com/tomo0111/gin-performance/common"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r.GET("/", controller.HelloController)
	r.GET("/items", controller.ItemController)
	r.Run(":8080")
}

