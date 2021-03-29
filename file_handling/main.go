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

	createFile("mdaburaihan.txt")
}
func createFile(fileName string)bool{
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer posf.Close()
	text := "My name is raihan"
	_, err = posf.Write([]byte(text))
	//fmt.Println(n, err)
	if err !=nil{
		return false
	}

	return true
}
