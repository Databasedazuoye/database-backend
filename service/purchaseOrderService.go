package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
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
		detail[idx].Date = formatTimeString(item.Date)
	}
	c.JSON(200, gin.H{
		"data": detail,
	})
}

func formatTimeString(timeString string) string {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		panic(err)
	}
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
