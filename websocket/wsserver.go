package main

import (
	"net/http"
	"io"
	"log"
	"strings"
	"github.com/gorilla/websocket"
	"fmt"
	"reflect"
)


var current_id = 0

var upgrader = websocket.Upgrader{}

type Peer struct {
	RemoteAddr string

	Conn *websocket.Conn
}


var peers []Peer;


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
	fmt.Println("Started new socket")
	// fmt.Printf("%+v\n", req)
	// fmt.Println(req.RemoteAddr)


	current_id += 1
	var local_id = current_id
	fmt.Println(local_id)

	conn, err := upgrader.Upgrade(w, req, nil)
	fmt.Println(reflect.TypeOf(conn))

	if err != nil {
		log.Println(err)
		return
	}

	peer := Peer{
		 req.RemoteAddr,
		 conn,
	}
	peers = append(peers, peer)
	fmt.Println(peers)
	for {

		messageType, p, err := conn.ReadMessage()
		readable_message := strings.Join(strings.Split(string(p), "%%%%"), ": " + req.RemoteAddr)
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		log.Println(req.RemoteAddr, readable_message)
		// for _, v := range peers {
		// 	fmt.Println(v.RemoteAddr, req.RemoteAddr)
		// 	if v.RemoteAddr != req.RemoteAddr {
		// 		fmt.Println("Need to send to ", v.RemoteAddr)
		// 		fmt.Println(v.RemoteAddr, req.RemoteAddr)
		// 		v.Conn.WriteMessage(messageType, p)
		// 	}
		// }


	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/socket", Socket)

	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
