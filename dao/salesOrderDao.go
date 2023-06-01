package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

func SalesOrderGetDetail() {
	//db := utils.GetDb()
	//sql := `select * from saleOrder`
	//db.SQL()
}

func CreateSalesOrder(salesOrder model.SalesOrder) int64 {
	db := utils.GetDb()
	insert, err := db.Insert(salesOrder)
	if err != nil {
		panic(err)
	}
	return insert
}
