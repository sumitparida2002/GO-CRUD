package main

import (
	"gorm.io/gorm"
)

// var DB *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// func main() {
// 	dsn := "root:sumit@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
// 	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	database.AutoMigrate(&Product{})

// 	DB = database

// 	DB.Create(&Product{Code: "D42", Price: 100})

// 	var product Product
// 	DB.First(&product, 1) // find product with integer primary key

// }
