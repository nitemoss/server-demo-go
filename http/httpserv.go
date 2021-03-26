package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	/*if req.Method == "GET" {
		io.WriteString(w, `<a href="/page1">GOTO PAGE1</a>`)
		//w.Write()
	} */
	fmt.Printf("%+v\n", req)
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }

		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else {
		w.WriteHeader(405)
	}

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
