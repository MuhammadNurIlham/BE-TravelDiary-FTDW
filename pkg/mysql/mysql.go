package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

var DB *gorm.DB

// For Connection Database
func DatabaseInit() {
	var err error
	dsn := "root:@tcp(localhost:3306)/thejourney-backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database Success!")
}
