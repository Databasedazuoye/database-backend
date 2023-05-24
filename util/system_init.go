package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

var db *xorm.Engine

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))

}

func InitMySQL() *xorm.Engine {
	fmt.Println(viper.GetString("mysql.dns"))
	var err error
	db, err = xorm.NewEngine("mysql", viper.GetString("mysql.dns"))
	if err != nil {
		panic(err)
	}
	return db
}

func GetDb() *xorm.Engine {
	return db
}
