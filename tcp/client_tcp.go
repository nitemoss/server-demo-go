package main

import (
	"net"
	"fmt"
	"os"
	// "encoding/json"
	"io/ioutil"
)

func fileCommand(filename string) []byte{
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error", err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(byteValue)
	return byteValue

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		var cmd string
		fmt.Print("(create/update/delete/read): ")
    fmt.Scanln(&cmd)

		var commandFilename = ""
		switch cmd {
			case "create":
				commandFilename = "data.json"
				break
			case "update":
				commandFilename = "update.json"

				break
			case "read":
				commandFilename = "read.json"

				break
			case "delete":
				commandFilename = "delete.json"

				break
			default:
				fmt.Println("Unknown command")
				return
		}
		fmt.Println("cmd", commandFilename)
		fmt.Println("Sending 'create' command...")
		var text = fileCommand(commandFilename)


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



}
