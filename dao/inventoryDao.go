package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

func InventoryGetByWarehouseId(id int64) []model.InventoryDTO {
	db := utils.GetDb()
	sql := `select i.*, goods.name as goods_name from 
                           (select * from inventory where warehouse_id = ?) i 
                               join goods on i.goods_id = goods.id`
	list := make([]model.InventoryDTO, 0)
	err := db.SQL(sql, id).Find(&list)
	if err != nil {
		panic(err)
	}
	return list
}

func InventoryGetByWarehouseIdAndGoodsId(warehouseId int, goodsId int) *model.Inventory {
	db := utils.GetDb()
	inventory := &model.Inventory{}
	get, err := db.Where("warehouse_id = ? and goods_id = ?", warehouseId, goodsId).Get(inventory)
	if err != nil {
		panic(err)
	}
	if !get {
		panic("unable to find specified record")
	}

	return inventory
}
