package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN string

func StartDb() {
	connection, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = connection

}
