package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contract", contract)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `My Name is raihan, i'm web developer`)
}

func contract(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("contract.gohtml")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
