package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Flyingmonk01/go-book-store/pkg/service"
	"github.com/gorilla/mux"
)

type BookController struct {
	bookService *service.BookService
}

func NewBookController(bookService *service.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	res := c.bookService.GetBookService(id)
	if res.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()

	w.Write([]byte("Get Books"))
}
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Create Book"))
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Update Book"))
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	w.Write([]byte("Delete Book"))
}
