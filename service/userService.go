package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
	"goodsManagement/model"
	utils "goodsManagement/utils"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	var user model.User

	c.ShouldBindJSON(&user)

	db := utils.GetDb()

	var userList []model.User
	db.Where("username=?", user.Username).Find(&userList)

	if len(userList) > 0 {
		c.JSON(409, gin.H{
			"msg": "此账号已存在",
		})
		return
	}

	user.Id = 0
	user.Password = utils.SHA256Hash(user.Password)
	if insert, err := db.Insert(user); err != nil {
		panic(err)
		fmt.Println(insert)
	}

	dao.SetRoleByUsername(user.Username, 3)
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

func Login(c *gin.Context) {
	user := model.User{}
	c.ShouldBindJSON(&user)

	fmt.Println(user.Username)
	fmt.Println(utils.SHA256Hash(user.Password))

	userList := dao.GetUser(user.Username, utils.SHA256Hash(user.Password))

	if len(userList) == 0 {
		c.JSON(400, gin.H{
			"msg": "账号或密码输入有误",
		})
		return
	}

	token := utils.GetToken(user.Username, userList[0].Id)

	c.JSON(200, gin.H{
		"data":     token,
		"username": user.Username,
	})

	permissions := dao.GetPermissions(userList[0].Id)

	ctx := context.Background()
	if err := utils.GetRedisHelper().Set(ctx, token, strings.Join(permissions, "|"), time.Hour*24).Err(); err != nil {
		panic(err)
	}

}

func GetPermission(c *gin.Context) {
	value, _ := c.Get("id")
	id, _ := value.(int)
	permissions := dao.GetPermissions(id)
	c.JSON(200, gin.H{
		"msg":  "获取权限成功",
		"data": permissions,
	})
}

func LogOut(c *gin.Context) {
	ctx := context.Background()
	token := c.GetHeader("Authorization")
	del := utils.GetRedisHelper().Del(ctx, token)
	result, _ := del.Result()
	if result == 1 {
		c.JSON(200, gin.H{
			"msg": "退出成功",
		})
		return
	} else {
		c.JSON(400, gin.H{
			"msg": "退出失败",
		})
	}

}
