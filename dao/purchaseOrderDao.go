package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

func getOrder() {
	db := utils.GetDb()
	db.SQL("select ")
}

func CreateOrder(order *model.PurchaseOrder) int64 {
	db := utils.GetDb()
	insert, err := db.Insert(order)
	if err != nil {
		panic(err)
	}
	return insert
}

func PurchaseOrderGetDetail() []model.PurchaseOrderView {
	db := utils.GetDb()
	sql := `
select result.*, supplier.name as supplier_name from
    (select result.*, goods.name as goods_name from
        (select purchase_order.*, warehouse.name as warehouse_name from purchase_order left join	supplier on supplier.id = purchase_order.supplier_id )
            result
            left join goods on goods.id = result.goods_id
    ) result
        left join supplier on result.supplier_id = supplier.id
        `
	list := make([]model.PurchaseOrderView, 0)

	err := db.SQL(sql).Find(&list)
	if err != nil {
		panic(err)
	}

	return list
}
