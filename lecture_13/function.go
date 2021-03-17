package main

import "fmt"

/*func add(x int, y int) int{
	// function body
	r := x+y
	return r

}
/*
func add(x, y int) int{
	// function body
	r := x+y
	return r

 }*/
func add(x, y int) int {
	// function body
	r := x + y
	return r

}

func main() {
	// code

	x := add(20, 20)
	fmt.Println(x)

	// anonimus functin 
	a := func(r, n int) int{
		r=r+n
		return r
	}
	fmt.Println(a(50,10))
}
