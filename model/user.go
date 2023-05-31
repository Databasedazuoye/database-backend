package model

import "fmt"

type User struct {
	Id       int `xorm:"pk autoincr"`
	Username string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("id:%v username:%v password:%v", u.Id, u.Username, u.Password)
}

func (User) TableName() string {
	return "user"
}
