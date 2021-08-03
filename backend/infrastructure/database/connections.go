package database

import (
	"fmt"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/gateway"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ref:https://gorm.io/docs/connecting_to_the_database.html

var db *gorm.DB

func Connect() {
	fmt.Println("Start Connecting.....")
	dsn := "sugar:sugar29fa038@tcp(127.0.0.1:3306)/sugar-webapp?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(
		&gateway.User{},
	)
	fmt.Println("Connected!!")
	db = gormDB
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}
