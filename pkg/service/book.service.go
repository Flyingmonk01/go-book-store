package service

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"strconv"

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
	slog.Info("Fetching book", "id", id)

	var book models.Book
	err := s.db.QueryRow(context.Background(), "SELECT id, name, author, published_year FROM book WHERE id =$1", id).Scan(&book.ID, &book.Name, &book.Author, &book.PublishedYear)

	if err != nil {
		slog.Error("Failed to fetch book", "error", err, "id", id)
		return book
	}

	slog.Info("Book fetched successfully", "id", book.ID, "name", book.Name, "author", book.Author)
	return book
}

func (s *BookService) CreateBookService(book models.Book) models.Book {
	slog.Info("Creating book", "name", book.Name, "author", book.Author, "year", book.PublishedYear)

	book.ID = strconv.Itoa(rand.IntN(10000000))

	_, err := s.db.Exec(context.Background(), "INSERT INTO book (id, name, author, published_year) VALUES ($1, $2, $3, $4)", book.ID, book.Name, book.Author, book.PublishedYear)
	if err != nil {
		slog.Error("Failed to create book", "error", err, "book_id", book.ID)
		return book
	}

	slog.Info("Book created successfully", "id", book.ID, "name", book.Name)
	return book
}

func (s *BookService) GetBooksService() []models.Book {
	slog.Info("Fetching all books")

	rows, err := s.db.Query(context.Background(), "SELECT id, name, author, published_year FROM book")
	if err != nil {
		slog.Error("Failed to query books", "error", err)
		return []models.Book{}
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.PublishedYear)
		if err != nil {
			slog.Error("Failed to scan book row", "error", err)
			continue
		}
		books = append(books, book)
	}

	slog.Info("Books fetched successfully", "count", len(books))
	return books
}

func (s *BookService) UpdateBookService(bookId string) models.Book {
	slog.Info("Updating book", "id", bookId)

	var book models.Book
	err := s.db.QueryRow(context.Background(), "SELECT id, name, author, published_year FROM book WHERE id =$1", bookId).Scan(&book.ID, &book.Name, &book.Author, &book.PublishedYear)

	if err != nil {
		slog.Error("Failed to fetch book for update", "error", err, "id", bookId)
		return book
	}

	// For demonstration, let's just update the name
	book.Name = book.Name + " (Updated)"

	_, err = s.db.Exec(context.Background(), "UPDATE book SET name=$1, author=$2, published_year=$3 WHERE id=$4", book.Name, book.Author, book.PublishedYear, book.ID)
	if err != nil {
		slog.Error("Failed to update book", "error", err, "id", book.ID)
		return book
	}

	slog.Info("Book updated successfully", "id", book.ID, "name", book.Name)
	return book
}

func (s *BookService) DeleteBookService(bookId string) bool {
	slog.Info("Deleting book", "id", bookId)
	_, err := s.db.Exec(context.Background(), "DELETE FROM book WHERE id=$1", bookId)
	if err != nil {
		slog.Error("Failed to delete book", "error", err, "id", bookId)
		return false
	}
	slog.Info("Book deleted successfully", "id", bookId)
	return true
}
