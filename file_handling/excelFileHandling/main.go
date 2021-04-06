package main

import (
	"fmt"
	"os"
)

func main() {

	fi, err := os.Stat("raihan.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fi.IsDir())

}
