package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"bytes"
)

func main() {
	client := &http.Client{}

	var body bytes.Buffer



	body.Write([]byte("Hello server!"))

	req, err := http.NewRequest("PUT", "http://localhost:8080/", &body)

	resp, err := client.Do(req)
	fmt.Printf("%+v\n", resp)
	if err != nil {panic(err) }

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }


}
