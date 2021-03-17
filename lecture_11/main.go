package main

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
	price  string
	pages  int
}

func main() {
	// structs
	var b1 Book
	b1.Title = "GOLANG book"
	b1.Author = "CALEB DOXY"
	b1.price = "$200"
	b1.ISBN = "52528522552"
	b1.pages = 300
	fmt.Println(b1)
	fmt.Println(b1.Title)
	fmt.Println(b1.Author)
	b3 := struct {
		Title  string
		Author string
		Price  string
		Pages  int
	}{
		Title:  "Golang book",
		Author: "Caleb doxy",
		Price:  "100$",
		Pages:  200,
	}
	fmt.Println(b3)

}
