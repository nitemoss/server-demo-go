package main

import (
	"net/http"
	"io"
	"log"
	"github.com/gorilla/websocket"
	"strings"
)

type Client struct {
	RemoteAddr string
	Conn *websocket.Conn
}

var clients []Client

var upgrader = websocket.Upgrader{}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")

	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }

		log.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else if req.Method == "OPTIONS" {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(405)
	}

}

func Socket(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)

	clients = append(clients, Client{req.RemoteAddr, conn})
	// log.Println(clients)

	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		readable_message := strings.Join(strings.Split(string(p), "%%%%"), ": ")
		log.Println(readable_message)
		if err != nil {
			log.Println(err)
			return
		}


		for _, c := range clients {
			if err := c.Conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}

		}

	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/socket", Socket)

	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
