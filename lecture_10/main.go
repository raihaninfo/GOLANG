package main

import "fmt"
func main(){
	// array, slice, map, struct
	var student [3] string
	student[0]="Raihan"
	student[1]="Hasan"
	student[2]="Ruhul"
	fmt.Println(student[2])
	var age [3] int
	age[0] = 12
	age[1] = 15
	age[2] = 20
	fmt.Println(age[2])
	// fmt.Println(len(student))
	// short hand way
	students := [3] string{"Raihan", "hasan", "Ruhul-Amin"}
	fmt.Println(students)


}