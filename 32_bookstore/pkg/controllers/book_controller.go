package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ramgopalsiddh/bookstore/pkg/models"
	"github.com/ramgopalsiddh/bookstore/pkg/utils"
)

var NewBook models.Book

func getBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetAllBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatBook(w http.Response, r *http.Request){
	CreatBook := &models.Book{}
	utils.ParseBody(r, CreatBook)
	b:=CreatBook.CreatBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}