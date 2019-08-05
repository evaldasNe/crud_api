package repository

import (
	"log"

	"github.com/evaldasNe/rest_api/database"
	"github.com/evaldasNe/rest_api/entity"
)

// GetAllAuthors func will find all authors in database
// and returns it
func GetAllAuthors() []entity.Author {
	//Init books var as a slice Book struct
	var authors []entity.Author

	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var author entity.Author
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(&author.ID, &author.Firstname, &author.Lastname)
		if err != nil {
			log.Fatal(err.Error())
		}
		authors = append(authors, author)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return []entity.Author(authors)
}

// GetAuthor func will find author in database by id
// and returns that author
func GetAuthor(ID int) entity.Author {
	db := database.DB

	// Execute the query
	rows, err := db.Query("SELECT * FROM authors WHERE id = ?", ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var author entity.Author

	for rows.Next() {
		// get data
		err = rows.Scan(&author.ID, &author.Firstname, &author.Lastname)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return entity.Author(author)
}

// CreateAuthor func will add new author in database
// returns inserted author's ID (int)
func CreateAuthor(author entity.Author) int {
	db := database.DB
	sqlStatement := `
	INSERT INTO authors (firstname, lastname)
	VALUES (?, ?)`
	res, err := db.Exec(sqlStatement, author.Firstname, author.Lastname)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	return int(id)
}

// UpdateAuthor func will update author data
func UpdateAuthor(author entity.Author) {
	db := database.DB
	sqlStatement := `
	UPDATE authors SET firstname = ?, lastname = ? 
	WHERE id = ?`
	var err error
	_, err = db.Exec(sqlStatement, author.Firstname, author.Lastname, author.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// DeleteAuthor func will delete author from database by ID
func DeleteAuthor(ID int) {
	db := database.DB
	sqlStatement := `DELETE FROM authors WHERE id = ?`
	_, err := db.Exec(sqlStatement, ID)
	if err != nil {
		log.Fatal(err.Error())
	}
}
