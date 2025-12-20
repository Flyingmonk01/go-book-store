package routes

import (
	"github.com/Flyingmonk01/go-book-store/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
