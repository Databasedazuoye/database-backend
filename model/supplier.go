package model

type Supplier struct {
	Id      int `xorm:"pk autoincr"`
	Name    string
	Phone   string
	Address string
}
