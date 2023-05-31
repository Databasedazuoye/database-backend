package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"strconv"
)

func WarehouseInsert(c *gin.Context) {
	var warehouse model.Warehouse
	c.BindJSON(&warehouse)
	insert := dao.WarehouseInsert(warehouse)
	if insert == 0 {
		c.JSON(400, gin.H{
			"msg": "插入失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "插入成功",
		})
	}
}

func WarehouseUpdate(c *gin.Context) {
	var warehouse model.Warehouse
	c.BindJSON(&warehouse)
	update := dao.WarehouseUpdate(warehouse)
	if update == 0 {
		c.JSON(400, gin.H{
			"msg": "更新失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "更新成功",
		})
	}
}

func WarehouseSelectAll(c *gin.Context) {
	all := dao.WarehouseSelectAll()
	c.JSON(200, gin.H{
		"data": all,
	})
}

func WarehouseDeleteById(c *gin.Context) {
	param := c.Param("id")
	i, _ := strconv.ParseInt(param, 10, 64)
	num := dao.WarehouseDeleteById(i)
	if num == 0 {
		c.JSON(400, gin.H{
			"msg": "删除失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	}
}
