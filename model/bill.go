package model

type Bill struct {
	Id     int `xorm:"pk"`
	UserId int
	Total  float64
	Date   string
}
