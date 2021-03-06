package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var s string
	fmt.Scan(&s)

	conn.Write([]byte(s))

	buf := make([]byte, 2000)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
