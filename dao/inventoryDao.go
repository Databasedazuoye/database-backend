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
