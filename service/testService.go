package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
)

func Test(c *gin.Context) {

	id, _ := dao.SaleOrderGetById(0)

	c.JSON(200, id)
}
