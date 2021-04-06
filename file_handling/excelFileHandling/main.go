package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("index.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	file.Write([]byte("My name is raihan"))
	defer file.Close()
	fmt.Println("raihan")
}
