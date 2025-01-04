package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	ـ "github.com/jinzhu/gorm/dialects/mysql"
	"bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}

