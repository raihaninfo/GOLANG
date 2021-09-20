package model

import "fmt"

// Insert database query
func CreateTodo(name, todo string) error {
	insertQuery, err := con.Query("INSERT INTO todo VALUES(?, ?)", name, todo)
	defer insertQuery.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
