package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO todo VALUES(?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
