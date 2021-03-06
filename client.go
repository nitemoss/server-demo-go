package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"time"
	"math/rand"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	w, h := termbox.Size()

	var data struct {
		X int32
		Y int32
	}

	var buf bytes.Buffer

	for {

		data.X = int32(rand.Intn(w))
		data.Y = int32(rand.Intn(h))

		fmt.Println(data)
		err = binary.Write(&buf, binary.LittleEndian, data)
		time.Sleep(100 * time.Millisecond)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		buf.Reset()
	}
	conn.Close()
}
