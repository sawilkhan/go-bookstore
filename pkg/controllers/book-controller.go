package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sawilkhan/go-bookstore/pkg/models"
	"github.com/sawilkhan/go-bookstore/pkg/utils"
)


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(200)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing json")
	}
	bookDetails := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(200)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := models.CreateBook(createBook)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(200)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0,0)
	if err != nil{
		fmt.Println("Error while parsing json")
	}
	bookDetails := models.DeleteBook(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(200)
	w.Write(res)
}