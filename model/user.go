package model

type User struct {
	Id       int
	Username string
	Password string
}

func (User) TableName() string {
	return "psi_user"
}
