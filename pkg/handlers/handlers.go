package handlers

import (
	"encoding/json"
	"github.com/mzcoding/translate_speaker/pkg/config"
	"log"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

// Task is create the task
func (m *Repository) Task(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "Task Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

	/*err := tasks.DB.Create({
		Word "dfsfs",
		Translate string
		Frequency int
		IsActive  bool
		Time      time.Time
	})
	if err != nil {

	} */
}
