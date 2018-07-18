package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomoyane/gin-performance/common"
	"github.com/tomoyane/gin-performance/controller"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r.GET("/", controller.HelloController)
	r.GET("/items", controller.ItemController)
	r.Run(":8080")
}

