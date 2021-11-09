package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Learn with Raihan")
	contant := "My name is Raihan, i'm Web Developer"
	file, err := os.Create("./newfile.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	length, err := io.WriteString(file, contant)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)

	defer file.Close()

	readFiles("./newfile.txt")
}

func readFiles(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}
