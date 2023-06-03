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

func SalesOrderInsert(c *gin.Context) {
	warehouseId := utils.StringToInt64(c.Query("warehouseId"))
	goodsId := utils.StringToInt64(c.Query("goodsId"))
	num := utils.StringToInt64(c.Query("num"))
	inventory := dao.InventoryGetByWarehouseIdAndGoodsId(warehouseId, goodsId)
	if inventory.Stock < int(num) {
		c.JSON(400, gin.H{
			"msg": "库存不足，出库失败",
		})
		return
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

func SaleOrderReview(c *gin.Context) {
	value := c.Query("id")
	id, _ := strconv.Atoi(value)
	salesOrder, get := dao.SaleOrderGetById(id)
	if !get {
		c.JSON(400, gin.H{
			"msg": "不存在此记录",
		})
	}
	inventory := dao.InventoryGetByWarehouseIdAndGoodsId(int64(salesOrder.WarehouseId), int64(salesOrder.GoodsId))
	if inventory.Stock-salesOrder.Num < 0 {
		c.JSON(400, gin.H{
			"msg": "库存不足",
		})
		return
	}
	i := dao.SalesOrderUpdateById(id, "审核通过")
	if i == 0 {
		c.JSON(400, gin.H{
			"msg": "请勿重复审核",
		})
		return
	}
	dao.DecreaseStock(salesOrder.GoodsId, salesOrder.WarehouseId, salesOrder.Num)

	////////////////////////////////////////////
	billId := generateID()
	var sum float64
	salesOrder, b := dao.SaleOrderGetById(id)
	if !b {
		c.JSON(400, gin.H{
			"msg": "不存在此记录",
		})
		return
	}
	dao.SalesOrderUpdateBillById(salesOrder.Id, billId)
	sum += salesOrder.Price

	bill := model.Bill{
		Id:     int(billId),
		UserId: 1,
		Total:  sum,
		Date:   time.Now().Format("2006-01-02 15:04:05"),
	}
	dao.BillInsert(bill)
	c.JSON(200, gin.H{
		"msg": "成功生成订单",
	})
}

func generateID() int {
	timestamp := time.Now().Unix()
	// 假设时间戳是 int64 类型，将其转换为字符串并截取后8位
	timestampStr := fmt.Sprintf("%d", timestamp)
	idStr := timestampStr[len(timestampStr)-8:]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		panic(err)
	}
	return id
}
