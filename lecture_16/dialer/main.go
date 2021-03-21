package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//bs:=make([]byte, 1024)
	conn.Write([]byte("Hello\n"))

	conn.Close()
}
