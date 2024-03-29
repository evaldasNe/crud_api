package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/evaldasNe/crud_api/entity"
	"github.com/evaldasNe/crud_api/repository"
	"github.com/gorilla/mux"
)

// GetAuthors func will return all authors
// from database in JSON format
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// GetAuthor func will return one author
// by ID from database in JSON format
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get params
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author, err := repository.GetAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if author.ID == 0 {
		http.Error(w, "Author not found by this ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// CreateAuthor func will add new author
// in database and returns that auhtor in JSON format
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author entity.Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	id, err := repository.CreateAuthor(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author.ID = id
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// UpdateAuthor func will update author's data
// and returns that auhtor's data in JSON format
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	auhtor, err := repository.GetAuthor(id) //get current author data
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if auhtor.ID == 0 {
		http.Error(w, "Author not found by this ID", http.StatusBadRequest)
		return
	}

	var authorUpdates entity.Author //get author updates
	_ = json.NewDecoder(r.Body).Decode(&authorUpdates)

	if authorUpdates.Firstname != "" {
		auhtor.Firstname = authorUpdates.Firstname
	}
	if authorUpdates.Lastname != "" {
		auhtor.Lastname = authorUpdates.Lastname
	}

	err = repository.UpdateAuthor(auhtor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auhtor)
}

// DeleteAuthor func will delete author
// from database and returns all auhtors
// from database in JSON format
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repository.DeleteAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authors, err := repository.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}
