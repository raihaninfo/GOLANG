package main

import (
	"fmt"
	"os"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(dir)
	text := "My name is raihan, i'm web Developer, I Love Golang"
	fileName:="raihan.txt"
	isErr := createFile(fileName, text)

	fmt.Println(isErr)

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
