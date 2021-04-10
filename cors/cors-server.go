package main

import (
	"net/http"
	"io"
	"fmt"
	// "encoding/json"
)



func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)

	fmt.Printf("%+v\n", req)
	fmt.Println(req.Method)
	// fmt.Println(req.Body)
	data, err := io.ReadAll(req.Body)
	if req.Method == "POST" {

		req.Body.Close()
		if err != nil {
			return
		}

		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else if req.Method == "GET" {

		req.Body.Close()
		if err != nil {return }

		// fmt.Printf("%s\n", data)
		io.WriteString(w, "successful get")
	} else {
		w.WriteHeader(405)
	}
	if req.Method == "OPTIONS" {
		w.WriteHeader(403)
		io.WriteString(w, "Error: not allowed method")
		return
	}



	io.WriteString(w, "You are on page1")
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are on page1")
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/page1", Handler2)

	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
