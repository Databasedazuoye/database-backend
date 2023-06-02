package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"goodsManagement/utils"
	"time"
)

func SalesOrderInsert(c *gin.Context) {
	warehouseId := utils.StringToInt64(c.Query("warehouseId"))
	goodsId := utils.StringToInt64(c.Query("goodsId"))
	num := utils.StringToInt64(c.Query("num"))
	inventory := dao.InventoryGetByWarehouseIdAndGoodsId(warehouseId, goodsId)
	if inventory.Stock < int(num) {
		c.JSON(400, gin.H{
			"msg": "库存不足，出库失败",
		})
	}
	goods := dao.GoodsGetById(goodsId)

	salesOrder := model.SalesOrder{
		GoodsId:     int(goodsId),
		WarehouseId: int(warehouseId),
		Num:         int(num),
		Price:       goods.PurchasingPrice,
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		Status:      "未审核",
	}

	insert := dao.SalesOrderInsert(salesOrder)
	if insert == 0 {
		c.JSON(400, gin.H{
			"msg": "出库失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"msg": "出库成功",
		})
	}

}

func SalesOrderGetAll(c *gin.Context) {
	list := dao.SalesOrderGetAll()
	c.JSON(200, gin.H{
		"data": list,
	})
}
