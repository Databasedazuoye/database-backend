package main

import (
	"goodsManagement/router"
	utils "goodsManagement/util"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()

	r.Run()
}
