package model

type Inventory struct {
	GoodsId     int
	WarehouseId int
	Stock       int
}

type InventoryDTO struct {
	GoodsId     int
	WarehouseId int
	Stock       int
	GoodsName   string
}
