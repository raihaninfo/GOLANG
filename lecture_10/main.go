package main

import "fmt"

func main() {
	// slice
	// var student [3]string
	// student[0] = "Raihan"
	// student[1] = "Hasan"
	// student[2] = "Ruhul"
	//  x := student[0:2]

	// x := make([]string, 0)
	var x [] string
	x = append(x, "Salman", "Farhan", "habib", "Hanas", "Ruhul")

	fmt.Println(x, len(x))
	fmt.Printf("%T", x) //data type chack 

}
