package handlers

import (
	"encoding/json"
	"example2_crud/helpers"
	"example2_crud/models"
	"io"
	"net/http"
	"strconv"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParseParamId(r)

	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, ok := models.GetBookById(id)

	if !ok {
		helpers.SendError(w, "Book with ID: "+strconv.FormatUint(uint64(id), 10)+" not found", http.StatusNotFound)
		return
	}

	helpers.SendJSON(w, book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	helpers.SendJSON(w, models.GetAllBooks())
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		helpers.SendError(w, "Not valid http body", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(bytes, &book); err != nil {
		helpers.SendError(w, "Provided book has wrong format", http.StatusBadRequest)
		return
	}

	models.CreateBook(&book)
	helpers.SendJSON(w, &book)
}

func PutBook(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParseParamId(r)

	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	book := models.Book{}
	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		helpers.SendError(w, "Not valid http body", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(bytes, &book); err != nil {
		helpers.SendError(w, "Provided book has wrong format", http.StatusBadRequest)
		return
	}

	models.PutBook(&book, id)
	helpers.SendJSON(w, &book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParseParamId(r)

	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.DeleteBook(id)

	helpers.SendJSON(w, nil)
}
