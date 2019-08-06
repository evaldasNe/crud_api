# crud_api
To get started
1. Install golang if you haven't already https://golang.org/dl/
2. Then move to $GOPATH/src 
cd $GOPATH/src
3. git clone https://github.com/evaldasNe/crud_api.git
4. cd crud_api
5. modify database connection in main.go
	database.DB, err = sql.Open("mysql", "username:password@/db_name")
6. go run crud_api
