package model

import (
	"encoding/json"
	"net/http"

	"github.com/evaldasNe/crud_api/entity"
	"github.com/evaldasNe/crud_api/repository"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

// GetBooks func will return all books
// from database in JSON format
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := repository.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook func will return one book
// by ID from database in JSON format
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get params
	book, err := repository.GetBook(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if book.ID == "" {
		http.Error(w, "Book NOT FOUND by this ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook func will add new book
// in database and returns that book in JSON format
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = xid.New().String() // generate id
	err := repository.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook func will update book's data
// and returns that book's data in JSON format
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := repository.GetBook(params["id"]) //get current book data
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if book.ID == "" {
		http.Error(w, "Book NOT FOUND by this ID", http.StatusBadRequest)
		return
	}

	var bookUpdates entity.Book //get book updates
	_ = json.NewDecoder(r.Body).Decode(&bookUpdates)

	if bookUpdates.Isbn != "" {
		book.Isbn = bookUpdates.Isbn
	}
	if bookUpdates.Title != "" {
		book.Title = bookUpdates.Title
	}
	if bookUpdates.Author != 0 {
		book.Author = bookUpdates.Author
	}

	repository.UpdateBook(book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// DeleteBook func will delete book
// from database and returns all books
// from database in JSON format
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repository.DeleteBook(params["id"])
	books, err := repository.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
