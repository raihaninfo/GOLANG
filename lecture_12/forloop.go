package main

import "fmt"

func main(){
	// for loop

	for i := 1; i <=10; i++ {
		fmt.Println(i)
	}

	students := []string{"Raihan", "hasan", "Ruhul-amin"}
	for i, std := range students{
		fmt.Println(i, std)
	}

	i:=0
	for i<10 {
		fmt.Println("raihan")
		i++
	}
	
}