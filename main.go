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
	Name   string
	Author string
}

func getAllBooks(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	DB.Find(&books)
	fmt.Fprintln(writer, "Getting Book")
	fmt.Fprintln(writer, "ID Name Author --------")

	for i := 0; i < len(books); i++ {
		fmt.Fprintf(writer, "%d %s %s\n", books[i].ID, books[i].Name, books[i].Author)
	}

}

func createBook(writer http.ResponseWriter, request *http.Request) {
	DB.Create(&Book{Name: "Sumit", Author: "Sumit"})
	fmt.Fprintf(writer, "Creating Book")

}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	DB.Delete(&Book{}, id)
	fmt.Fprintf(writer, "Deleting Book")

	getAllBooks(writer, request)
}

func updataBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	DB.Model(&Book{}).Where("id = ?", id).Update("name", "Me")

	fmt.Fprintf(writer, "Updating Book")

	getAllBooks(writer, request)

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getAllBooks).Methods("GET")
	router.HandleFunc("/create", createBook)
	router.HandleFunc("/delete/{id}", deleteBook)
	router.HandleFunc("/update/{id}", updataBook)

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
