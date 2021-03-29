package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)

	posf, err := os.Create("raihan.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()
	text := "My name is raihan"
	n, err := posf.Write([]byte(text))
	fmt.Println(n, err)

	//posf.Close()
}
