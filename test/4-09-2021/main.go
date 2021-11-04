package main

import "fmt"

func main() {
	// search "M"
	// pragraph := "my name is raihan, i'm from bangladesh"
	// for _, value := range pragraph {
	// 	valueString := string(value)
	// 	if valueString == "m" {
	// 		fmt.Println(valueString)
	// 	}
	// }

	// array of sum
	numbers := [...]float64{1, 2, 3}
	result := 0.0
	for _, value := range numbers {
		result += value
	}
	fmt.Println(result)

}
