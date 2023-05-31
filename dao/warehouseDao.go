package dao

import (
	"fmt"
	"goodsManagement/model"
	"goodsManagement/utils"
	"strconv"
)

func WarehouseInsert(warehouse model.Warehouse) int64 {
	db := utils.GetDb()
	warehouse.Id = 0
	insert, err := db.Insert(warehouse)
	if err != nil {
		panic(err)
	}
	return insert
}

func WarehouseUpdate(warehouse model.Warehouse) int64 {
	db := utils.GetDb()
	fmt.Println("id:" + strconv.Itoa(warehouse.Id))
	update, err := db.Id(warehouse.Id).Update(warehouse)
	if err != nil {
		panic(err)
	}
	return update
}

func WarehouseDeleteById(id int64) int64 {
	db := utils.GetDb()
	warehouse := model.Warehouse{}
	i, err := db.Id(id).Delete(warehouse)
	if err != nil {
		panic(err)
	}
	return i
}

func WarehouseSelectAll() []model.Warehouse {
	db := utils.GetDb()
	warehouseList := make([]model.Warehouse, 0)
	db.Find(&warehouseList)
	return warehouseList
}
