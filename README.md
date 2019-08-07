# CRUD API
## To get started
1. Install golang if you haven't already https://golang.org/dl/
2. Then move to `$GOPATH/src` 
```sh
cd $GOPATH/src
```
3. 
```sh 
git clone https://github.com/evaldasNe/crud_api.git
```
4. 
```sh 
cd crud_api
```
5. Import database structure from file `db.sql` to your local database
6. Modify database connection in `main.go` at line 19
```go
database.DB, err = sql.Open("mysql", "username:password@/db_name")
```
7. 
```sh 
go run crud_api
```


# Introduction
With this API you can create, read, update and delete two data tables (books and authors) in your database

# Error Codes
If something went wrong you will get `400 Bad Request` response status and error message.

# Success Codes
If you add new row to the table you will get `201 Created` response status.
All other requests will give `200 OK` response status.

# `POST` Post author
```url
http://localhost:8000/api/authors
```
## Headers
`Content-Type application/json`
## Body 
###### raw (application/json)
```JSON
{
	"firstname":"John",
	"lastname":"Smith"
}
```
# `GET` Get all authors
```url
http://localhost:8000/api/authors
```
# `GET` Get author
```url
http://localhost:8000/api/authors/{id}
```
# `PUT` Update author
```url
http://localhost:8000/api/authors/{id}
```
## Headers
`Content-Type application/json`
## Body 
###### raw (application/json)
```JSON
{
	"firstname":"Changed",
	"lastname":"Name"
}
```
# `DELETE` Delete author
```url
http://localhost:8000/api/authors/{id}
```

# `POST` Post book
```url
http://localhost:8000/api/books
```
## Headers
`Content-Type application/json`
## Body 
###### raw (application/json)
```JSON
{
	"isbn":"0101",
	"title":"New book",
	"author": 7
}
```
# `GET` Get all books
```url
http://localhost:8000/api/books
```
# `GET` Get book
```url
http://localhost:8000/api/books/{id}
```
# `PUT` Update book
```url
http://localhost:8000/api/books/{id}
```
## Headers
`Content-Type application/json`
## Body 
###### raw (application/json)
```JSON
{
	"isbn":"555",
	"title":"Updated book"
}
```
# `DELETE` Delete book
```url
http://localhost:8000/api/books/{id}
```
