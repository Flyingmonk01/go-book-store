package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Flyingmonk01/go-book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")

	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
