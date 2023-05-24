package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/model"
	utils "goodsManagement/util"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model.User
	db := utils.GetDb()

	get, _ := db.Where("username = ?", username).Get(&user)
	if get {
		c.JSON(409, gin.H{
			"msg": "此账号已存在",
		})
		return
	}

	user = model.User{
		Username: username,
		Password: password,
	}

	db.Insert(user)

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	db := utils.GetDb()

	user := &model.User{
		Username: username,
		Password: password,
	}

	get, _ := db.Get(user)
	if get {
		c.JSON(200, gin.H{
			"msg": "登录成功",
		})
		return
	} else {
		c.JSON(403, gin.H{
			"msg": "账号或密码输入有误",
		})
	}

}
