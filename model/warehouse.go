package model

type Warehouse struct {
	Id      int `xorm:"pk autoincr"`
	Name    string
	Phone   string
	Address string
}
