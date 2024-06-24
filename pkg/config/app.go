package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(){
	dsn := fmt.Sprintf("root:1111@tcp(localhost:3306)/book_store?charset=utf8&parseTime=True&loc=Local")
	
	d, err := gorm.Open("mySql", dsn)

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
