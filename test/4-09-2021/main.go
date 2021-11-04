package main

import "fmt"

 

func main() {
	pragraph := "my name is raihan, i'm from bangladesh"

	for _, value := range pragraph {
		valueString := string(value)
		if valueString == "m" {
			fmt.Println(valueString)
		}
	}
}
