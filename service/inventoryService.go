package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"strconv"
)

func InventoryGetByWarehouseId(c *gin.Context) {
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	list := dao.InventoryGetByWarehouseId(i)
	c.JSON(200, gin.H{
		"msg":  "获取成功",
		"data": list,
	})

}
