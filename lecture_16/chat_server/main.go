package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close() //await
	log.Printf("server started on :8888")

	for {

		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//continue
		}

		bs := make([]byte, 1024)
		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}

		reqstr := string(bs[:n])
		fmt.Println(reqstr)
		recvTime := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf(`Your message: %s, received at %s`, reqstr, recvTime)
		conn.Write([]byte(msg))

	}

}
