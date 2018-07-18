package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/tomoyane/gin-performance/entity"
	"github.com/tomoyane/gin-performance/common"
)

func ItemController (c *gin.Context) {
	var items []entity.Item
	common.Db.Find(&items)
	c.JSON(http.StatusOK, &items)
}