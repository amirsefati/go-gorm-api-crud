package controllers

import (
	"bookstore/pkg/models"
	"encoding/json"
	"net/http"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBook := models.GetAllBooks()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("content", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
