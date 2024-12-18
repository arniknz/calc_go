package main

import (
	application "github.com/arniknz/calc_go/internal/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/calculate", application.HandleRequest).Methods("POST")

	application.StartServer()
}
