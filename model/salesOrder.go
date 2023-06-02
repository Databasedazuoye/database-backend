package model

type SalesOrder struct {
	Id          int `xorm:"pk autoincr"`
	BillId      int
	GoodsId     int
	WarehouseId int
	Num         int
	Price       float64
	Date        string
	Status      string
}

type SalesOrderDTO struct {
	Id            int
	BillId        int
	GoodsId       int
	WarehouseId   int
	Num           int
	Price         float64
	Date          string
	Status        string
	GoodsName     string
	WarehouseName string
}
