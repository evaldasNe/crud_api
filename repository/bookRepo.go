package repository

import (
	"errors"

	"github.com/evaldasNe/crud_api/database"
	"github.com/evaldasNe/crud_api/entity"
)

// GetAllBooks func will find all books in database
// and returns books and errors
func GetAllBooks() ([]entity.Book, error) {
	//Init books var as a slice Book struct
	var books []entity.Book

	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return books, err
	}

	defer rows.Close()

	var book entity.Book
	// Fetch rows
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}

	return books, err
}

// GetBook func will find book in database by id
// and returns that book and errors
func GetBook(ID string) (entity.Book, error) {
	db := database.DB
	var book entity.Book

	// Execute the query
	rows, err := db.Query("SELECT * FROM books WHERE id = ?", ID)
	if err != nil {
		return book, err
	}

	defer rows.Close()

	for rows.Next() {
		// get data
		err = rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Author)
		if err != nil {
			return book, err
		}
	}
	if err = rows.Err(); err != nil {
		return book, err
	}

	return book, err
}

// CreateBook func will add new book in database
// and returns errors
func CreateBook(book entity.Book) error {
	db := database.DB
	sqlStatement := `
	INSERT INTO books (id, isbn, title, author)
	VALUES (?, ?, ?, ?)`
	var err error
	_, err = db.Exec(sqlStatement, book.ID, book.Isbn, book.Title, book.Author)

	return err
}

// UpdateBook func will update book data
// returns errors
func UpdateBook(book entity.Book) error {
	db := database.DB
	sqlStatement := `UPDATE books SET isbn = ?, title = ?, author = ? WHERE id = ?`
	res, err := db.Exec(sqlStatement, book.Isbn, book.Title, book.Author, book.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No changes was made")
	}

	return err
}

// DeleteBook func will delete book from database by id
// returns errors
func DeleteBook(ID string) error {
	db := database.DB
	sqlStatement := `DELETE FROM books WHERE id = ?`
	res, err := db.Exec(sqlStatement, ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Book not found")
	}

	return err
}
