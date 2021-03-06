package main

import (
	"net"
	"fmt"
)

type Data struct {
	N string
	Ch chan int
}

var dat Data

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	dat.Ch = make(chan int, 1)
	dat.Ch <- 1

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}

	<- dat.Ch

	fmt.Println(string(buf[:n]))
	dat.N = string(buf[:n])

	data := []byte("Connection great")
	conn.Write(data)

	dat.Ch <- 1

	conn.Close()
}
