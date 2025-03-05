package main

import (
	"log"

	application "github.com/arniknz/calc_go/internal/app/agent"
)

func main() {
	agent := application.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}
