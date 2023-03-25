package models

import (
	"github.com/hemanth-ks97/bookstore-go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Title string `gorm:""json:"title"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book,db
}

func DeleteBookById(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(&book)
	return book
}