package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	id     int
	name   string
	author string
}

func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	var book Book
	DB.Find(&book)
	fmt.Fprintf(writer, "Getting Book")
	fmt.Printf(book.author)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	DB.Create(&Book{id: 1, name: "Sumit", author: "Sumit"})
	fmt.Fprintf(writer, "Creating Book")

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getAllBooks).Methods("GET")
	router.HandleFunc("/create", createBook)

	dataSource := "root:sumit@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	bookstoreDB, error := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	}

	bookstoreDB.AutoMigrate(&Book{})
	DB = bookstoreDB

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}

}
