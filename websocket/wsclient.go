
package main

import (
	"log"
	"time"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"reflect"
)

func readMessages(c *websocket.Conn){
	for {



		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("%s", strings.Join(strings.Split(string(message), "%%%%"), ": "))
		time.Sleep(1*time.Second)
	}
}

func main() {
	var name string


	fmt.Print("your name: ")
	fmt.Scanf("%s", &name)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	fmt.Println(reflect.TypeOf(c))

	if err != nil {
		log.Fatal("dial:", err)
	}
	// go readMessages(c)
	defer c.Close()

	for {
		var msg string
		fmt.Print(":: ")
		fmt.Scanf("%s", &msg)

		c.WriteMessage(websocket.TextMessage, []byte(name + "%%%%" + msg))

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("%s", strings.Join(strings.Split(string(message), "%%%%"), ": "))
		time.Sleep(1*time.Second)
	}


}
