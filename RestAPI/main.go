package main

import (
	"GOLANG/RestAPI/controller"
	"GOLANG/RestAPI/model"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	var port string = ":8000"
	fmt.Println("Listening to port", port)
	log.Fatal(http.ListenAndServe(port, mux))

}
