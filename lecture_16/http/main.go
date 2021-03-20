package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

		nl, err := net.Listen("tcp", ":8888")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		bs := make([]byte, 1024)

		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}
		fmt.Println(n)
		reqstr := string(bs)
		fmt.Println(reqstr)

		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>This is my first golang web<br>
		<h2>Md Abu Raihan</h2>
		<h3>Web Developer</h3>
		</strong></body></html>`

		resp := "HTTP/1.1 200 OK\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n%s"
		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)
		conn.Write([]byte(msg))
	//conn.Close()

}
