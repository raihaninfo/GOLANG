package main

import (
	"GOLANG/RestAPI/controller"
	"GOLANG/RestAPI/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000", mux))
}
