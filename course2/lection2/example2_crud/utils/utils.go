package utils

import (
	"example2_crud/handlers"
	"github.com/gorilla/mux"
)

func BuildBooksResource(router *mux.Router, url string) {
	router.HandleFunc(url, handlers.GetAllBooks).Methods("GET")
	router.HandleFunc(url, handlers.CreateBook).Methods("POST")
	router.HandleFunc(url+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(url+"/{id}", handlers.PutBook).Methods("PUT")
	router.HandleFunc(url+"/{id}", handlers.DeleteBook).Methods("DELETE")
}
