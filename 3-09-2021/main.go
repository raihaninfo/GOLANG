package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// Switch
	var city string = "dhaka"

	switch city {
	case "dhaka":
		fmt.Println("I love Dhaka")
	case "kushtia":
		fmt.Println("My home toun is Kushtia")
	default:
		fmt.Println("I love bangladesh")
	}

	// loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var testText string = "this is text"
	for index, lerrer := range testText {
		fmt.Println("index : ", index, "value", string(lerrer))
		// fmt.Println(string(lerrer))

	}

}
