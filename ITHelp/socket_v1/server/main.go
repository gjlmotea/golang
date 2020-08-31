package main

import (
	"fmt"
	"net"
)

func main() {
	addr := "127.0.0.1:5000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Print("error when listen", addr)
	}
	fmt.Println("listen to", addr)
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Print("error when conn", addr)
		}

		buf := make([]byte, 1024)
		strLen, _ := conn.Read(buf)
		fmt.Println(strLen)
		fmt.Println(string(buf))
	}

}
