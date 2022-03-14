package main

import (
	"fmt"
	"github.com/mzcoding/translate_speaker/pkg/models/tasks"
	"net/http"
)

const portNumber = ":9090"

func main() {

	_, err := tasks.NewDb()
	if err != nil {
		fmt.Printf("Error connect to databse")
	}

	//routes
	http.HandleFunc("/task", CreateTask)
	_ = http.ListenAndServe(portNumber, nil)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func showTask() {

}

func updateTask() {

}

func destroyTask() {

}
