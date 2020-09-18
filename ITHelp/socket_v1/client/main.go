package main

import (
	"fmt"
	"net"
)

func main() {
	addr := "127.0.0.1:5000"
	conn, err := net.Dial("tcp", addr)

	//tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("error when listen", addr)
	}
	defer conn.Close()
	conn.Write([]byte("Something Here"))
}
