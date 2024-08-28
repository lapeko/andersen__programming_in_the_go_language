package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if _, err := fmt.Fprintf(response, "<h1>Hello from server</h1>"); err != nil {
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
