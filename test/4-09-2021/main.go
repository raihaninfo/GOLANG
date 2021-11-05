package main

import (
	"fmt"
	"time"
)

func main() {
	// // search "M"
	// pragraph := "my name is raihan, i'm from bangladesh"
	// for _, value := range pragraph {
	// 	valueString := string(value)
	// 	if valueString == "m" {
	// 		fmt.Println(valueString)
	// 	}
	// }

	// // array of sum
	// numbers := [...]float64{1, 2, 3}
	// result := 0.0
	// for _, value := range numbers {
	// 	result += value
	// }
	// fmt.Println(result)

	todayTime := time.Now()
	fmt.Println(todayTime)
	// fizzBuzz(15)
}

// // FizzBuzz problem solved
// func fizzBuzz(n int32) {
// 	// Write your code here
// 	if n%3 == 0 && n%5 == 0 {
// 		fmt.Println("FizzBuzz")
// 	} else if n%3 == 0 {
// 		fmt.Println("Fizz")
// 	} else if n%5 == 0 {
// 		fmt.Println("Buzz")
// 	} else {
// 		fmt.Println(n)
// 	}

// }
