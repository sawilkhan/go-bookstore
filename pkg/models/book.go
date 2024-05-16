package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sawilkhan/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book{
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) *Book{
	var fetchedBook Book
	db.Where("ID=?", Id).Find(fetchedBook)
	return &fetchedBook
}

func DeleteBook(Id int64) Book{
	var deletedBook Book
	db.Where("ID=?",Id).Delete(deletedBook)
	return deletedBook
}