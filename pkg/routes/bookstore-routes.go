package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBookById).Methods("DELETE")

}
