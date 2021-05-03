
package main

import (
	"log"
	// "time"
	"fmt"
	"strings"
	"github.com/gorilla/websocket"
)

func wait(c *websocket.Conn){
	for {
		
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Println(strings.Join(strings.Split(string(message), "%%%%"), ": "))
		// time.Sleep(1*time.Second)
	}

}

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	var name string
	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)

	defer c.Close()
	go wait(c)
	for {
		var msg string
		fmt.Print("msg: ")
		fmt.Scanf("%s", &msg)

		c.WriteMessage(websocket.TextMessage, []byte(name + "%%%%" + msg))


		// _, message, err := c.ReadMessage()
		// if err != nil {
		// 	log.Println("read:", err)
		// 	return
		// }
		// log.Printf("recv: %s", message)
		// time.Sleep(1*time.Second)
	}


}
