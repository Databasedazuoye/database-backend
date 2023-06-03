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

func BillGetById(id int) *model.Bill {
	db := utils.GetDb()
	bill := new(model.Bill)
	_, err := db.Id(id).Get(bill)
	if err != nil {
		panic(err)
	}

	return bill
}
