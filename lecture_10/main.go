package main

import "fmt"

func main() {
<<<<<<< Updated upstream
	// map, struct
	x := make(map[string]string)
	x["Name"] = "Raihan"
	x["Age"] = "21"
	x["Job"] = "Student"
	delete(x, "Job")
	fmt.Println(x)
	// fmt.Println(x["Name"])
=======
	// array, slice, map, struct
	var student [3]string
	student[0] = "Raihan"
	student[1] = "Hasan"
	student[2] = "Ruhul"
	fmt.Println(student[2])
	var age [3]int
	age[0] = 12
	age[1] = 15
	age[2] = 20
	fmt.Println(age[2])
	// fmt.Println(len(student))
	// short hand way
	students := [...]string{"Raihan", "hasan", "Ruhul-Amin", "monirul", "Rabbi"}
	fmt.Println(students)

>>>>>>> Stashed changes
}
