package dao

import (
	"goodsManagement/model"
	"goodsManagement/utils"
)

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
        (select purchase_order.*, warehouse.name as warehouse_name from purchase_order left join	warehouse on warehouse.id = purchase_order.warehouse_id )
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

func Review(id int, status string) int64 {
	db := utils.GetDb()
	order := model.PurchaseOrder{Status: status}
	update, err := db.Id(id).Update(&order)
	if err != nil {
		panic(err)
	}
	return update
}

func PurchaseOrderGetById(id int) *model.PurchaseOrder {
	db := utils.GetDb()
	//purchaseOrder := new(model.PurchaseOrder)
	purchaseOrder := &model.PurchaseOrder{}
	_, err := db.Id(id).Get(purchaseOrder)
	if err != nil {
		panic(err)
	}
	return purchaseOrder
}
