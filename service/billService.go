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

func BillGetAll(c *gin.Context) {
	list := dao.BillSelectAll()
	c.JSON(200, gin.H{
		"data": list,
	})
}

func BillGetById(c *gin.Context) {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	fmt.Println(id)
	bill := dao.BillGetById(id)
	c.JSON(200, gin.H{
		"data": bill,
	})
}
