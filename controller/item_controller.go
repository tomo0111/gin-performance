package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo0111/gin-performance/entity"
	"net/http"
	"github.com/tomo0111/gin-performance/common"
)

func ItemController (c *gin.Context) {
	var items []entity.Item
	c.JSON(http.StatusOK, common.Db.Find(&items))
}