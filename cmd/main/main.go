package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hemanth-ks97/bookstore-go/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}