package controller

import (
	"GOLANG/RestAPI/model"
	"GOLANG/RestAPI/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//POST Method
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet { //GET Method
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
