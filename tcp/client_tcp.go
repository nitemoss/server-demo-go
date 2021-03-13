package main

import (
	"net"
	"fmt"
	// "encoding/json"
	"io/ioutil"
)

func fileCommand(filename string){
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	return text
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer conn.Close()



	fmt.Println("Sending 'create' command...")
	var text = fileCommand("json/data.json")


	bytes := []byte(text)
	fmt.Println(bytes)

	conn.Write(bytes)

	buf := make([]byte, 2000)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(n, buf)

}
