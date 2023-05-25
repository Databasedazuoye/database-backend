package router

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/service"
	utils "goodsManagement/util"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Cors())

	users := r.Group("/users")
	{
		users.POST("/register", service.Register)
		users.POST("/login", service.Login)
	}

	return r
}
