package model

import (
	"database/sql"
	"fmt"
)

var con *sql.DB
// Database connection function
func Connect() *sql.DB {

	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/restapi-go")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database....")
	con = db
	return db
}
