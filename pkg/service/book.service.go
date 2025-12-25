package service

import (
	"context"
	"fmt"
	"os"

	"github.com/Flyingmonk01/go-book-store/pkg/models"
	"github.com/jackc/pgx/v5"
)

type BookService struct {
	db *pgx.Conn
}

func NewBookService(db *pgx.Conn) *BookService {
	return &BookService{db: db}
}

func (s *BookService) GetBookService(id string) models.Book {

	fmt.Println("id:", id)
	var book models.Book
	err := s.db.QueryRow(context.Background(), "SELECT id, name, author, published_year FROM book WHERE id =$1", id).Scan(&book.ID, &book.Name, &book.Author, &book.PublishedYear)
	fmt.Println("book:", book)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)

	}
	return book
}
