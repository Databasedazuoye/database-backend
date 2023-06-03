package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"strconv"
	"strings"
)

func GoodsInsert(c *gin.Context) {

	var goods model.Goods
	c.BindJSON(&goods)
	insertNum := dao.GoodsInsert(goods)
	if insertNum == 0 {
		c.JSON(400, gin.H{
			"msg": "插入失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "插入成功",
		})
	}

}

func GoodsUpdate(c *gin.Context) {

	var goods model.Goods
	c.BindJSON(&goods)
	updateNum := dao.GoodsUpdate(goods)
	if updateNum == 0 {
		c.JSON(400, gin.H{
			"msg": "更新失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "更新成功",
		})
	}
}

func GoodsSelectAll(c *gin.Context) {

	all := dao.GoodsSelectAll()
	for idx, item := range all {
		all[idx].ExpirationDate = strings.Split(item.ExpirationDate, "T")[0]
		all[idx].ManufactureDate = strings.Split(item.ManufactureDate, "T")[0]
		all[idx].PurchaseDate = strings.Split(item.PurchaseDate, "T")[0]
	}

	c.JSON(200, gin.H{
		"data": all,
	})
}

func GoodsDeleteById(c *gin.Context) {
	param := c.Param("id")
	i, _ := strconv.ParseInt(param, 10, 64)
	num := dao.GoodsDeleteById(i)
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

func GoodsGetByNameLike(c *gin.Context) {
	name := c.Query("name")
	dao.GoodsGetByNameLike(name)
	c.JSON(200, gin.H{})
}
