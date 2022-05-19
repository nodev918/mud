package main

import (
	"fmt"
	"net"
	"os"
)

func sender(conn net.Conn) {
	words := "hello world!"
	conn.Write([]byte(words))
	fmt.Println("send over")
}

func main() {
	// server := "localhost:1024"
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)
}
