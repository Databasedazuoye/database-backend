package dao

import (
	"goodsManagement/model"
	utils "goodsManagement/utils"
)

func GetUser(username string, password string) []model.User {
	db := utils.GetDb()
	userList := make([]model.User, 0)
	sql := "username = ? and password = ?"
	db.Where(sql, username, password).Find(&userList)
	return userList
}

func GetPermissions(userId int) []string {
	db := utils.GetDb()
	sql := "select name from permission where id in (select permission_id from role_permission where role_id = (select role_id from user_role where user_id = ?)) "

	var permissionList []string

	db.SQL(sql, userId).Find(&permissionList)

	return permissionList
}

func SetRoleByUsername(username string, roleId int) {
	db := utils.GetDb()
	sql := "insert into user_role (user_id, role_id) values ((select id from user where username = ?), ?)"
	db.Exec(sql, username, roleId)
}
