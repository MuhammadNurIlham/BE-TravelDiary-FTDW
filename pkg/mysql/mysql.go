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

	// ==== if using my sql ====
	dsn := "root:@tcp(localhost:3306)/thejourney-backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// ================ setup using postgresql ================== //
	// var DB_HOST = os.Getenv("DB_HOST")
	// var DB_USER = os.Getenv("DB_USER")
	// var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	// var DB_NAME = os.Getenv("DB_NAME")
	// var DB_PORT = os.Getenv("DB_PORT")

	// // ==== if using postgresql ====== //
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database Success!")
}
