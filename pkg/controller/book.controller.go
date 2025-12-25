package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Flyingmonk01/go-book-store/pkg/models"
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
	res := c.bookService.GetBooksService()
	json.NewEncoder(w).Encode(res)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
	}
	res := c.bookService.CreateBookService(book)

	json.NewEncoder(w).Encode(res)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)

	bookId := params["id"]
	res := c.bookService.UpdateBookService(bookId)

	json.NewEncoder(w).Encode(res)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// db := config.GetDB()
	params := mux.Vars(r)
	fmt.Println(params)
	bookId := params["id"]
	c.bookService.DeleteBookService(bookId)
	json.NewEncoder(w).Encode(map[string]any{"message": "Book deleted successfully", "success": true})

}
