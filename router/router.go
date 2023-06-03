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
		users.DELETE("", middleware.AuthMiddleware(), service.LogOut)
	}

	goods := r.Group("/goods")
	{
		goods.POST("", service.GoodsInsert)
		goods.PUT("", service.GoodsUpdate)
		goods.DELETE("/:id", service.GoodsDeleteById)
		goods.GET("", service.GoodsSelectAll)
		goods.GET("/nameLike", service.GoodsGetByNameLike)

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
		purchaseOrder.DELETE("/:id", service.PurchaseOrderDeleteById)
	}

	inventory := r.Group("/inventories")
	{
		inventory.GET("/:id", service.InventoryGetByWarehouseId)
	}

	salesOrder := r.Group("/sale")
	{
		salesOrder.GET("/create", service.SalesOrderInsert)
		salesOrder.GET("", service.SalesOrderGetAll)
		salesOrder.GET("/review", service.SaleOrderReview)

	}

	bill := r.Group("/bill")
	{
		bill.POST("", middleware.AuthMiddleware(), service.BillCreate)
		bill.GET("", service.BillGetAll)
	}

	r.GET("/test", service.Test)

	return r
}
