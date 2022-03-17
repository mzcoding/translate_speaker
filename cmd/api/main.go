package main

import (
	"fmt"
	"github.com/mzcoding/translate_speaker/pkg/config"
	"github.com/mzcoding/translate_speaker/pkg/handlers"
	"github.com/mzcoding/translate_speaker/pkg/models/tasks"
	"log"
	"net/http"
)

const portNumber = ":9091"

func main() {

	var app config.AppConfig

	_, err := tasks.NewDb()
	if err != nil {
		fmt.Printf("Error connect to databse")
	}

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//routes
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
