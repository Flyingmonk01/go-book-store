package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Flyingmonk01/go-book-store/pkg/config"
	"github.com/Flyingmonk01/go-book-store/pkg/controller"
	"github.com/Flyingmonk01/go-book-store/pkg/routes"
	"github.com/Flyingmonk01/go-book-store/pkg/service"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")
	// 1. Connect to database
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	d := config.ConnectToDB()

	// 2. Create service with DB
	bookService := service.NewBookService(d)

	// 3. Create controller with service
	bookController := controller.NewBookController(bookService)

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r, bookController)

	http.Handle("/", r)
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
