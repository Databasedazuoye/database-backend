package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"goodsManagement/utils"
	"strconv"
	"time"
)

func CreateOrder(c *gin.Context) {
	supplierId, _ := strconv.ParseInt(c.Query("supplierId"), 10, 64)
	goodsId, _ := strconv.ParseInt(c.Query("goodsId"), 10, 64)
	warehouseId, _ := strconv.ParseInt(c.Query("warehouseId"), 10, 64)
	num, _ := strconv.ParseInt(c.Query("num"), 10, 64)

	goods := dao.GoodsGetById(goodsId)

	order := &model.PurchaseOrder{
		Id:          0,
		SupplierId:  supplierId,
		GoodsId:     goodsId,
		WarehouseId: warehouseId,
		Num:         num,
		Price:       goods.PurchasingPrice,
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		Status:      "未审核",
	}

	num = dao.CreateOrder(order)
	if num == 0 {
		c.JSON(400, gin.H{
			"msg": "insert failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "入库成功",
	})

}

func QueryPurchaseOrderDetail(c *gin.Context) {
	detail := dao.PurchaseOrderGetDetail()
	for idx, item := range detail {
		detail[idx].Date = utils.FormatTimeString(item.Date)
	}
	c.JSON(200, gin.H{
		"data": detail,
	})
}

func Review(c *gin.Context) {
	param := c.Query("status")
	id, _ := strconv.Atoi(c.Query("id"))

	var s string

	purchaseOrder := dao.PurchaseOrderGetById(id)
	fmt.Println(purchaseOrder)
	if purchaseOrder.Status != "未审核" {
		c.JSON(400, gin.H{
			"msg": "审核失败，请勿审核已审核的订单",
		})
		return
	}

	if param == "1" {
		s = "审核通过"
	} else {
		s = "拒绝通过"
	}
	i := dao.Review(id, s)
	if i == 0 {
		c.JSON(400, gin.H{
			"msg": "审核失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "审核成功",
	})

}

func PurchaseOrderDeleteById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	i := dao.PurchaseOrderDeleteById(id)
	if i == 0 {
		c.JSON(400, gin.H{
			"msg": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})

}
