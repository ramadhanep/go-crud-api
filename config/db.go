package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("mysql", "root:Romadhanep17;17@/go_crud_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	DB.LogMode(true)

	fmt.Println("Database Connected")
}
