package repository

import (
	"github.com/evaldasNe/crud_api/database"
	"github.com/evaldasNe/crud_api/entity"
)

// GetAllAuthors func will find all authors in database
// and returns all authors and errors
func GetAllAuthors() ([]entity.Author, error) {
	//Init books var as a slice Book struct
	var authors []entity.Author
	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		return authors, err
	}

	defer rows.Close()

	var author entity.Author
	// Fetch rows
	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Firstname, &author.Lastname)
		if err != nil {
			return authors, err
		}
		authors = append(authors, author)
	}
	if err = rows.Err(); err != nil {
		return authors, err
	}

	return authors, err
}

// GetAuthor func will find author in database by id
// and returns that author and errors
func GetAuthor(ID int) (entity.Author, error) {
	db := database.DB
	var author entity.Author

	// Execute the query
	rows, err := db.Query("SELECT * FROM authors WHERE id = ?", ID)
	if err != nil {
		return author, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Firstname, &author.Lastname)
		if err != nil {
			return author, err
		}
	}
	if err = rows.Err(); err != nil {
		return author, err
	}

	return author, err
}

// CreateAuthor func will add new author in database
// returns inserted author's ID (int64) and errors
func CreateAuthor(author entity.Author) (int64, error) {
	db := database.DB
	var id int64
	sqlStatement := `
	INSERT INTO authors (firstname, lastname)
	VALUES (?, ?)`
	res, err := db.Exec(sqlStatement, author.Firstname, author.Lastname)
	if err != nil {
		return id, err
	}
	id, err = res.LastInsertId()

	return id, err
}

// UpdateAuthor func will update author data
// returns errors
func UpdateAuthor(author entity.Author) error {
	db := database.DB
	sqlStatement := `
	UPDATE authors SET firstname = ?, lastname = ? 
	WHERE id = ?`
	var err error
	_, err = db.Exec(sqlStatement, author.Firstname, author.Lastname, author.ID)

	return err
}

// DeleteAuthor func will delete author from database by ID
// returns errors
func DeleteAuthor(ID int) error {
	db := database.DB
	sqlStatement := `DELETE FROM authors WHERE id = ?`
	_, err := db.Exec(sqlStatement, ID)

	return err
}
