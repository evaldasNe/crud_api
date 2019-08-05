package repository

import (
	"log"

	"github.com/evaldasNe/rest_api/database"
	"github.com/evaldasNe/rest_api/entity"
)

// GetAllBooks func will find all books in database
// and returns it
func GetAllBooks() []entity.Book {
	//Init books var as a slice Book struct
	var books []entity.Book

	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var book entity.Book
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err.Error())
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return []entity.Book(books)
}

// GetBook func will find book in database by id
// and returns that book
func GetBook(ID string) entity.Book {
	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM books WHERE id = ?", ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var book entity.Book

	for rows.Next() {
		// get data
		err = rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return entity.Book(book)
}

// CreateBook func will add new book in database
func CreateBook(book entity.Book) {
	db := database.DB
	sqlStatement := `
	INSERT INTO books (id, isbn, title, author)
	VALUES (?, ?, ?, ?)`
	var err error
	_, err = db.Exec(sqlStatement, book.ID, book.Isbn, book.Title, book.Author)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// UpdateBook func will update book data
func UpdateBook(book entity.Book) {
	db := database.DB
	sqlStatement := `UPDATE books SET isbn = ?, title = ?, author = ? WHERE id = ?`
	var err error
	_, err = db.Exec(sqlStatement, book.Isbn, book.Title, book.Author, book.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// DeleteBook func will delete book from database by id
func DeleteBook(ID string) {
	db := database.DB
	sqlStatement := `DELETE FROM books WHERE id = ?`
	_, err := db.Exec(sqlStatement, ID)
	if err != nil {
		log.Fatal(err.Error())
	}
}
