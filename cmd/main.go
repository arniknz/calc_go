package main

import (
	application "github.com/arniknz/calc_go/internal/app"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/calculate", application.HandleRequest).Methods("POST")

	application.StartServer()
}
