package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/utils"
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
	//salesOrder := model.SalesOrder{
	//	GoodsId: goodsId,
	//
	//}

}
