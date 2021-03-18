package main

import "fmt"

func update(a*int){
	*a=*a+10
}

func main(){
	var x int 
	var y *int 	
	x=10
	y=&x
	// memory address
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(&y)

	update(&x)
	fmt.Println(x)
}