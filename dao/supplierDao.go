package dao

import (
	"fmt"
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

	sql := `select * from purchase_order where supplier_id = ?`
	list := make([]model.PurchaseOrder, 0)
	err := db.SQL(sql, id).Find(&list)
	if len(list) != 0 {
		return 0
	}

	supplier := model.Supplier{}
	i, _ := db.Id(id).Delete(supplier)

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

func SupplierGetByNameLike(name string) []model.Supplier {
	db := utils.GetDb()
	list := make([]model.Supplier, 0)
	err := db.Where("name like ?", fmt.Sprintf("%%%s%%", name)).Find(&list)
	if err != nil {
		panic(err)
	}
	return list

}
