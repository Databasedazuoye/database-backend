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

func SalesOrderInsert(salesOrder model.SalesOrder) int64 {
	db := utils.GetDb()
	insert, err := db.Insert(salesOrder)
	if err != nil {
		panic(err)
	}
	return insert
}

func SalesOrderGetAll() []model.SalesOrderDTO {
	db := utils.GetDb()
	sql := `select s.*, warehouse.name as warehouse_name from 
		(select sales_order.*, goods.name as goods_name from sales_order left join goods on sales_order.goods_id = goods.id) s 
		left join warehouse on warehouse.id = s.warehouse_id
		`
	list := make([]model.SalesOrderDTO, 0)
	err := db.SQL(sql).Find(&list)
	if err != nil {
		panic(err)
	}
	return list
}

func SalesOrderUpdateById(id int, status string) int64 {
	db := utils.GetDb()

	//sql := `update sales_order set status = ? where id = ?`
	//session := db.SQL(sql, status, id)
	salesOrder := model.SalesOrder{Status: status}
	update, err := db.Id(id).Update(&salesOrder)
	if err != nil {
		panic(err)
	}
	return update
}

func SaleOrderGetById(id int) (model.SalesOrder, bool) {
	db := utils.GetDb()
	salesOrder := model.SalesOrder{}
	get, err := db.Id(id).Get(&salesOrder)
	if err != nil {
		panic(err)
	}
	return salesOrder, get
}

func SalesOrderUpdateBillById(id int, billId int) {
	db := utils.GetDb()
	sql := `update sales_order set bill_id = ? where id = ?`
	_, err := db.Exec(sql, billId, id)
	if err != nil {
		panic(err)
	}
}
