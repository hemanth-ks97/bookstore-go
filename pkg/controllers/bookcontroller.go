package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hemanth-ks97/bookstore-go/pkg/models"
	"github.com/hemanth-ks97/bookstore-go/pkg/utils"
)

var NewBook models.Book

func GetBookHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Write(res)
}

func GetBookByIdHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request){
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	sId := vars["bookID"]
	Id, err := strconv.ParseInt(sId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBookById(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func UpdateBookByIdHandler(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	sId := vars["bookID"]
	Id, err := strconv.ParseInt(sId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookdetails, db := models.GetBookById(Id)
	if updateBook.Author != ""{
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Title != ""{
		bookdetails.Title = updateBook.Title
	}
	if updateBook.Publisher != ""{
		bookdetails.Publisher = updateBook.Publisher
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}