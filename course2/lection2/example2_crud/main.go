package main

import (
	"example2_crud/utils"
	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	appUrlPrefix = "/api/v1"
	booksUrl     = appUrlPrefix + "/books"
)

var port string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	router := mux.NewRouter()
	utils.BuildBooksResource(router, booksUrl)
	log.Println("Server is running on PORT", port)
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
