package main

import "fmt"

func main(){

	// fmt.Println("Enter your name & age : ")
	// var name string
	// var age int
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("your name %s & age is %d", name, age)

	fmt.Println("Enter your First name : ")
	var fname string
	fmt.Scanln(&fname)

	fmt.Println("Enter your last name : ")
	var lname string
	fmt.Scanln(&lname)

	fmt.Print("Your full name :")
	fmt.Print(fname + " " + lname)

}
