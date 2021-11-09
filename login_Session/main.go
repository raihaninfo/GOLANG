package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)


func main()  {
	mux := mux.NewRouter()
	mux.HandleFunc("/", login)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", mux)
}

func login(w http.ResponseWriter, r *http.Request)  {
	tmp,_:=template.ParseFiles("view/index.gohtml")
	tmp.Execute(w, nil)
}
