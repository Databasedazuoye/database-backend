package main

import (
	"goodsManagement/router"
	utils "goodsManagement/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Use(utils.Cors())
	r.Run()
}
