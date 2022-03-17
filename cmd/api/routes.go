package main

import (
	"github.com/bmizerany/pat"
	"github.com/mzcoding/translate_speaker/pkg/config"
	"github.com/mzcoding/translate_speaker/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Post("/task", http.HandlerFunc(handlers.Repo.Task))
	//update task
	//delete task

	return mux
}
