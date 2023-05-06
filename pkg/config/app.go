package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(sql.Open("test.db"), &gorm.Config{}) //Open("gorm:gorm@tcp(localhost:9910)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
