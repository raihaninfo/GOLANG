package main

import "fmt"

func main(){
	var x int 
	var y *int 
	// memory address
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(&y)
}