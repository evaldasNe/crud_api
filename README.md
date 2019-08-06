# crud_api
To get started
1. Install golang if you haven't already https://golang.org/dl/
2. Then move to $GOPATH/src 
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
5. Import database structure from file ```db.sql``` to your local database
6. Modify database connection in main.go at line 19
```go
database.DB, err = sql.Open("mysql", "username:password@/db_name")
```
7. 
```sh 
go run crud_api
```
