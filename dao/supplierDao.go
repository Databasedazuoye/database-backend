package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

func SupplierInsert(supplier model.Supplier) int64 {
	db := utils.GetDb()
	supplier.Id = 0
	insert, err := db.Insert(supplier)
	if err != nil {
		panic(err)
	}
	return insert
}

func SupplierUpdate(supplier model.Supplier) int64 {
	db := utils.GetDb()
	update, err := db.Id(supplier.Id).Update(supplier)
	if err != nil {
		panic(err)
	}
	return update
}

func SupplierDeleteById(id int64) int64 {
	db := utils.GetDb()
	supplier := model.Supplier{}
	i, err := db.Id(id).Delete(supplier)
	if err != nil {
		panic(err)
	}
	return i
}

func SupplierSelectAll() []model.Supplier {
	db := utils.GetDb()
	supplierList := make([]model.Supplier, 0)
	db.Find(&supplierList)
	return supplierList
}
