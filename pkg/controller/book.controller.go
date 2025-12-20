package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Flyingmonk01/go-book-store/pkg/models"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	book := models.Book{
		ID:   id,
		Name: "Sample Book",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()

	w.Write([]byte("Get Books"))
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Create Book"))
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Update Book"))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Delete Book"))
}
