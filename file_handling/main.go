package main

import (
	"fmt"
	"os"
)

func main() {
	// golang file handiling
	text := "My name is raihan, i'm web Developer, I Love Golang"
	fileName:="raihan.txt"
	createFile(fileName, text)

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
