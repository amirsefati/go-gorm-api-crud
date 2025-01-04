package routes

import (
    "github.com/gorilla/mux"
    "bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/books", controllers.getBooks).Methods("GET")
    router.HandleFunc("/book/{book_id}", controllers.getBookById).Methods("GET")
    router.HandleFunc("/book/{book_id}", controllers.updateBookById).Methods("PUT")
    router.HandleFunc("/book/{book_id}", controllers.deleteBookById).Methods("DELETE")

}