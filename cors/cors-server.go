package main

import (
	"net/http"
	"io"
	"fmt"
	// "crypto/sha256"
	// "strings"
	"time"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")


type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`

	jwt.StandardClaims
}


var users = map[string]string{
	"root": "password#12345",
	"user2": "password2",
}


func Handler(w http.ResponseWriter, req *http.Request) {
	// db := map[string]string{ // checksum = sha256(username + " " + password)
	//   "root": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855с",
	// }

	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// w.WriteHeader(http.StatusOK)

	// fmt.Printf("%+v\n", req)
	// fmt.Println(req.Method)

	data, err := io.ReadAll(req.Body)

	var raw map[string]interface{}
    if err := json.Unmarshal(data, &raw); err != nil {
        panic(err)
    }

	var method_name = raw["method_name"]
	mySigningKey := []byte("AllYourBase")



	if method_name == "auth" {
		claims := UserClaims{
			raw["username"],
			raw["password"],

			jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "test",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		fmt.Printf("%v %v", ss, err)

	}
	fmt.Println("\n\n")
	// for name, values := range req.Header {
	//     // Loop over all values for the name.
	//     for _, value := range values {
	//         fmt.Println(name, value)
	//     }
	// }

	fmt.Println(req.Header.Get("Authorization"))



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
