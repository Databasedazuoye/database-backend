package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
)

func Test(c *gin.Context) {
	inventory := dao.InventoryGetByWarehouseIdAndGoodsId(1, 1241)
	c.JSON(200, inventory)
}
