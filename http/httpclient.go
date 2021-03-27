package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
)


func fileCommand(filename string) []byte{
	jsonFile, err := os.Open("../tcp/" + filename)
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
	client := &http.Client{}

	var body bytes.Buffer


	var method string

	fmt.Println("Method: ")
	fmt.Scanln(&method)

	var cmd string
	fmt.Print("(data/update/delete/read    .json): ")
	fmt.Scanln(&cmd)


	var text = fileCommand(cmd)



	fmt.Println(text)



	body.Write([]byte(text))

	req, err := http.NewRequest(method, "http://localhost:8080/", &body)

	resp, err := client.Do(req)
	fmt.Printf("%+v\n", resp)
	if err != nil {panic(err) }

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }


}
