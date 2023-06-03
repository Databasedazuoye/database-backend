package dao

import (
	"fmt"
	"goodsManagement/model"
	"goodsManagement/utils"
)

func GoodsInsert(goods model.Goods) int64 {
	db := utils.GetDb()
	goods.Id = 0
	insert, err := db.Insert(goods)
	if err != nil {
		panic(err)
	}
	return insert
}

func GoodsUpdate(goods model.Goods) int64 {
	db := utils.GetDb()
	update, err := db.Id(goods.Id).Update(goods)
	if err != nil {
		panic(err)
	}
	return update
}

func GoodsDeleteById(id int64) int64 {
	db := utils.GetDb()
	goods := model.Goods{}
	i, err := db.Id(id).Delete(goods)
	if err != nil {
		panic(err)
	}
	return i
}

func GoodsSelectAll() []model.Goods {
	db := utils.GetDb()
	goodsList := make([]model.Goods, 0)
	db.Find(&goodsList)
	return goodsList
}

func GoodsGetById(id int64) *model.Goods {
	db := utils.GetDb()
	goods := new(model.Goods)
	get, _ := db.Id(id).Get(goods)
	if !get {
		panic("不存在此商品")
	}
	return goods
}

func GoodsGetByNameLike(name string) []model.Goods {
	db := utils.GetDb()
	list := make([]model.Goods, 0)
	err := db.Where("name like ?", fmt.Sprintf("%%%s%%", name)).Find(&list)
	if err != nil {
		panic(err)
	}
	return list

}
