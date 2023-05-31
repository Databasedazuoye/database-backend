package model

type Goods struct {
	Id              int `xorm:"pk autoincr"`
	Name            string
	Category        string
	Unit            string
	ManufactureDate string
	ExpirationDate  string
	PurchaseDate    string
	PurchasingPrice float64
	Num             int
}
