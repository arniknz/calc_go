package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/arniknz/calc_go/pkg/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var request Request

	if r.Method == "POST" {
		body := r.Body
		defer body.Close()
		decoder := json.NewDecoder(body)
		err := decoder.Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		expression := request.Expression
		result, err := calculator.Calc(expression)

		if err != nil {
			if errors.Is(err, calculator.InvalidExpression) {
				fmt.Fprintf(w, "err: %s", err.Error())
			} else {
				fmt.Fprintf(w, "unknown err")
			}
		} else {
			fmt.Fprintf(w, "result: %f", result)
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
