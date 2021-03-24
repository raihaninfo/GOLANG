package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// package level variables
var db *sql.DB
var err error

func init() {
	// Open up our database connection.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/hosting")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished

	//defer db.Close()
	fmt.Println("database connection successfull")
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}
func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)

}

func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func request(w http.ResponseWriter, r *http.Request) {

	// Method 1
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")
	//fmt.Println(name, company, email)
	//fmt.Fprintf(w, `received form %s %s %s`, name, company, email) //response

	//Method 2

	// r.ParseForm()
	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }

	qs := "INSERT INTO `reqest` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
	sql := fmt.Sprintf(qs, name, company, email)
	insert, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Fprintf(w, `OK`)

}
