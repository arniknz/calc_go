package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arniknz/calc_go/pkg/calculator"
)

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type Request struct {
	Expression string `json:"expression"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var request Request
	var response Response

	if r.Method == "POST" {
		body := r.Body
		defer body.Close()
		decoder := json.NewDecoder(body)
		err := decoder.Decode(&request)

		if err != nil {
			response.Error = fmt.Sprintf("Invalid JSON format: %v", err)
			http.Error(w, response.Error, http.StatusBadRequest)
			return
		}

		expression := request.Expression
		result, err := calculator.Calc(expression)

		if err == nil {
			response.Result = fmt.Sprintf("%v", result)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response.Error = "Expression is not valid"
			http.Error(w, response.Error, http.StatusUnprocessableEntity)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func StartServer() {
	log.Printf("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
