package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	allTodos, _ := json.Marshal(todos)
	w.Write(allTodos)
}

var todos = []Todo{
	{uuid.FromStringOrNil("B5ADEF4C-A3AA-46A8-91D8-D6AB553F75B5"), "fix that shit", false},
	{uuid.FromStringOrNil("18B0E287-791E-4A59-A098-C76D3BF452F1"), "properly use Echo", false},
	{uuid.FromStringOrNil("7432E6CB-D42D-418E-B2B5-1680486C8056"), "Write a good API", false},
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", todosHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(rice.MustFindBox("client/build").HTTPBox())))

	log.Println("Listening at port 3002")
	loggedRouter := handlers.RecoveryHandler()(handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":3002", loggedRouter)

}

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}
