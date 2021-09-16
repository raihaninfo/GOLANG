package model

import "GOLANG/RestAPI/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
