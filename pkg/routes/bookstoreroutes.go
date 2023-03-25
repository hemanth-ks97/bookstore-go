package routes

import (
	"github.com/gorilla/mux"
	"github.com/hemanth-ks97/bookstore-go/pkg/controllers"
)

func BookStoreRoutes(r *mux.Router){
	r.HandleFunc("/book", controllers.CreateBookHandler).Methods("POST")
	r.HandleFunc("/book", controllers.GetBookHandler).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.GetBookByIdHandler).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.UpdateBookByIdHandler).Methods("PUT")
	r.HandleFunc("/book/{bookID}", controllers.DeleteBookByIdHandler).Methods("DELETE")
}