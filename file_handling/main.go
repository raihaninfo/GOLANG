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
	text:= "My name is raihan, i'm web Developer"
	createFile("mdaburaihan.txt", text)
	
}
func createFile(fileName, text string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer posf.Close()
	posf.Write([]byte(text))
	return true
}
