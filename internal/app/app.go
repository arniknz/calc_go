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

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	d := json.NewDecoder(r.Body)
	err := d.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculator.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculator.ErrInvalidExpression) {
			fmt.Fprintf(w, "err: %s", err.Error())
		} else {
			fmt.Fprintf(w, "unknown err")
		}

	} else {
		fmt.Fprintf(w, "result: %f", result)
	}
}

func StartServer() {
	http.HandleFunc("/", CalcHandler)
	log.Printf("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
