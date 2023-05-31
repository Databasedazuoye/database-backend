package service

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/dao"
)

func Test(c *gin.Context) {
	dao.PurchaseOrderGetDetail()
}
