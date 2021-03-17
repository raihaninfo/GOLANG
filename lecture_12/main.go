package main

import "fmt"

func main() {
	fmt.Print("Enter your age : ")
	var age int
	fmt.Scanf("%d", &age)

	if age < 3 {
		fmt.Println("you are infant")
	} else if age > 3 && age < 18 {
		fmt.Println("You are baby")
	} else if age > 18 && age < 21 {
		fmt.Println("You are adult")
	} else if age > 21 && age < 50 {
		fmt.Println("You are middle age")
	} else if age > 50 && age < 100 {
		fmt.Println("you are Old man")
	}

}
