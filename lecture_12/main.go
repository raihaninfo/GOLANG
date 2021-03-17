package main

import "fmt"

func main(){
	fmt.Print("Enter your age : ")
	var age int
	fmt.Scanf("%d", &age)

	if age <3 {
		fmt.Println("your are infant")
	} else if age >3 && age <18{
		fmt.Println("Your are baby")
	}else if age >18 && age <21{
		fmt.Println("Your are adult")
	}

}