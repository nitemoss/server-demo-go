package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)
type Action struct {
	Action string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID string  `json:"id"`
	Salary float64 `json:"salary"`
	Subject string `json:"subject"`
	Classroom []string `json:"classroom"`
	Person struct {
		Name string `json:"name"`
		Surname string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}

type CreateTeacher struct {
	T Teacher `json:"data"`
}
func (action *CreateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action CreateTeacher) Process() {
	//add action.T to slice of data
	fmt.Printf("Processing %T\n...", action)
	fmt.Println(action.T)
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}
func (action *UpdateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action UpdateTeacher) Process() {
	fmt.Printf("Processing %T\n...", action)
}

type ReadTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
func (action *ReadTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action ReadTeacher) Process() {
	fmt.Printf("Processing %T\n...", action)
}

type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
func (action *DeleteTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (action DeleteTeacher) Process() {
	fmt.Printf("Processing %T\n...", action)
}



func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
	/*if req.Method == "GET" {
		io.WriteString(w, `<a href="/page1">GOTO PAGE1</a>`)
		//w.Write()
	} */
	fmt.Printf("%+v\n", req)
	fmt.Println(req.Method)
	fmt.Println(req.Body)
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

	// var act Action
	// var readableData map[string]interface{}
	// json.Unmarshal(data, &readableData)
	// json.Unmarshal(data, &act)
	// fmt.Println("Received", readableData["action"], " command")
	//
	// var obj GeneralObject
	//
	// switch act.ObjName {
	// case "Teacher":
	// 	obj = &Teacher{}
	// }
	//
	// var toDo DefinedAction
	// switch act.Action {
	// 	case "create":
	// 		toDo = obj.GetCreateAction()
	// 	case "update":
	// 		toDo = obj.GetUpdateAction()
	// 	case "read":
	// 		toDo = obj.GetReadAction()
	// }
	// toDo.GetFromJSON(data)
	//
	// toDo.Process()

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
