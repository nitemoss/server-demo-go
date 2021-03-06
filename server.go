package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= 20; i++ {
		// fmt.Println("new conn")
		handleConnection(listener, i)
	}
	termbox.Close()

}
func handleConnection(con *net.UDPConn, i int) {

	buf := make([]byte, 4000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n*i])

	var data struct {
		X int32
		Y int32
	}

	err = binary.Read(buff, binary.LittleEndian, &data)
	// fmt.Print(data) // проблема здесь - сервер получает одно и то же сообщение, хотя клиент каждый раз отправляет разные


	if err != nil {
		return
	}

	termbox.SetCell(int(data.X), int(data.Y), '*', termbox.ColorRed, termbox.ColorDefault)
	termbox.Flush()
	// buff.Reset(make([]byte, 4000))


	// fmt.Println(data)
}
