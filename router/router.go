package router

import (
	"github.com/gin-gonic/gin"
	"goodsManagement/middleware"
	"goodsManagement/service"
	utils "goodsManagement/utils"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Cors())
	r.Use(gin.Logger())

	users := r.Group("/users")
	{
		users.POST("/register", service.Register)
		users.POST("/login", service.Login)
		users.GET("/permission", middleware.AuthMiddleware(), service.GetPermission)
	}

	goods := r.Group("/goods")
	{
		goods.POST("", service.GoodsInsert)
		goods.PUT("", service.GoodsUpdate)
		goods.DELETE("/:id", service.GoodsDeleteById)
		goods.GET("", service.GoodsSelectAll)

	}

	warehouse := r.Group("/warehouses")
	{
		warehouse.GET("", service.WarehouseSelectAll)
		warehouse.POST("", service.WarehouseInsert)
		warehouse.PUT("", service.WarehouseUpdate)
		warehouse.DELETE("/:id", service.WarehouseDeleteById)
	}

	supplier := r.Group("/suppliers")
	{
		supplier.GET("", service.SupplierSelectAll)
		supplier.POST("", service.SupplierInsert)
		supplier.PUT("", service.SupplierUpdate)
		supplier.DELETE("/:id", service.SupplierDeleteById)
	}

	purchaseOrder := r.Group("/purchase")
	{
		purchaseOrder.GET("/create", service.CreateOrder)
		purchaseOrder.GET("", service.QueryPurchaseOrderDetail)
		purchaseOrder.GET("/review", service.Review)
	}

	r.GET("/test", service.Test)

	return r
}
