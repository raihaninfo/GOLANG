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
	n, err := posf.Write([]byte("My name is Raihan"))

	fmt.Println(n, err)

	posf.Close()
}
