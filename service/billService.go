package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	"goodsManagement/utils"
	"time"
)

func BillCreate(c *gin.Context) {
	var arr utils.Array
	err := c.ShouldBindJSON(&arr)
	if err != nil {
		panic(err)
	}
	array := arr.Array

	var sum float64

	billId := generateID()
	for i := range array {
		salesOrder, b := dao.SaleOrderGetById(i)
		if !b {
			c.JSON(400, gin.H{
				"msg": "不存在此记录",
			})
			return
		}
		dao.SalesOrderUpdateBillById(i, billId)
		sum += salesOrder.Price

	}
	//value, _ := c.Get("id")

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

func BillGetAll(c *gin.Context) {
	list := dao.BillSelectAll()
	c.JSON(200, gin.H{
		"data": list,
	})
}
