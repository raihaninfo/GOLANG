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

	conn.Write([]byte("This is my first message.\n"))

	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(string(bs[:n]))
	conn.Close()
}
