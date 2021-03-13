package main

import (
	"net"
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


type Data struct {
	N string
	Ch chan int
}

var dat Data

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	dat.Ch = make(chan int, 1)
	dat.Ch <- 1

	for {
		fmt.Println("New connection")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}


	<- dat.Ch
	var act Action
	var readableData map[string]interface{}

	json.Unmarshal(buf[:n], &readableData)

	fmt.Println("Received", readableData["action"], " command")


	err = json.Unmarshal(buf[:n], &act)
	if err != nil {
		fmt.Println(err)
		return
	}


	var obj GeneralObject
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	}

	var toDo DefinedAction
	switch act.Action {
		case "create":
			toDo = obj.GetCreateAction()
		case "update":
			toDo = obj.GetUpdateAction()
		case "read":
			toDo = obj.GetReadAction()
	}

	toDo.GetFromJSON(buf[:n])

	toDo.Process()

	data := []byte("Received")
	conn.Write(data)

	// fmt.Println("response sent")

	dat.Ch <- 1

	conn.Close()
}
