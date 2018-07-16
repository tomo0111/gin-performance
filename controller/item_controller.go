package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo0111/gin-performance/entity"
	"net/http"
	"github.com/tomo0111/gin-performance/common"
)

func ItemController (c *gin.Context) {
	var items []entity.Item
	common.Db.Find(&items)
	c.JSON(http.StatusOK, &items)
}