package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"strconv"
)

func SupplierInsert(c *gin.Context) {
	var supplier model.Supplier
	c.BindJSON(&supplier)
	insert := dao.SupplierInsert(supplier)
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

func SupplierUpdate(c *gin.Context) {
	var supplier model.Supplier
	c.BindJSON(&supplier)
	update := dao.SupplierUpdate(supplier)
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

func SupplierSelectAll(c *gin.Context) {
	all := dao.SupplierSelectAll()
	c.JSON(200, gin.H{
		"data": all,
	})
}

func SupplierDeleteById(c *gin.Context) {
	param := c.Param("id")
	i, _ := strconv.ParseInt(param, 10, 64)
	num := dao.SupplierDeleteById(i)
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
