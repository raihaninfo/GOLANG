package main

import "fmt"

func main() {
	// map, struct
	x := make(map[string]string)
	x["Name"] = "Raihan"
	x["Age"] = "21"
	x["Job"] = "Student"
	delete(x, "Job")
	fmt.Println(x)
	// fmt.Println(x["Name"])

}
