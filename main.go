package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/evaldasNe/crud_api/database"
	"github.com/evaldasNe/crud_api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	var err error
	database.DB, err = sql.Open("mysql", "root:@/api")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer database.DB.Close()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", model.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", model.GetBook).Methods("GET")
	r.HandleFunc("/api/books", model.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", model.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", model.DeleteBook).Methods("DELETE")

	r.HandleFunc("/api/authors", model.GetAuthors).Methods("GET")
	r.HandleFunc("/api/authors/{id}", model.GetAuthor).Methods("GET")
	r.HandleFunc("/api/authors", model.CreateAuthor).Methods("POST")
	r.HandleFunc("/api/authors/{id}", model.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/api/authors/{id}", model.DeleteAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
