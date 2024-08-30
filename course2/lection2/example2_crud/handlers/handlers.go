package handlers

import (
	"example2_crud/helpers"
	"example2_crud/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, ok := vars["id"]

	if !ok {
		helpers.SendError(w, "ID not found in path", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseUint(idString, 10, 0)

	if err != nil {
		helpers.SendError(w, "ID is not valid", http.StatusBadRequest)
		return
	}

	book, ok := models.GetBookById(uint(id))

	if !ok {
		helpers.SendError(w, "Book with ID: "+strconv.FormatUint(id, 10)+" not found", http.StatusNotFound)
		return
	}

	helpers.SendJSON(w, book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	helpers.SendJSON(w, models.GetAllBooks())
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	//helpers.SetJsonHeader(w)
}

func PutBook(w http.ResponseWriter, r *http.Request) {
	//helpers.SetJsonHeader(w)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//helpers.SetJsonHeader(w)
}
