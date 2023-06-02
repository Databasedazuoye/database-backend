package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

func BillInsert(bill model.Bill) {
	db := utils.GetDb()
	db.Insert(bill)
}

func BillSelectAll() []model.Bill {
	db := utils.GetDb()
	billList := make([]model.Bill, 0)
	db.Find(&billList)
	return billList
}
