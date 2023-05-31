package model

type PurchaseOrder struct {
	Id          int64 `xorm:"pk autoincr"`
	SupplierId  int64
	GoodsId     int64
	WarehouseId int64
	Num         int64
	Price       float64
	Date        string
	Status      string
}

type PurchaseOrderView struct {
	Id            int64
	SupplierId    int64
	GoodsId       int64
	Warehouse     int64
	Num           int64
	Price         float64
	Date          string
	Status        string
	SupplierName  string
	GoodsName     string
	WarehouseName string
}
