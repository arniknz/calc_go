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

	if r.Method == "POST" {
		defer r.Body.Close()
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		err := d.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // 400
			return
		}

		result, err := calculator.Calc(request.Expression)
		if err != nil {
			if errors.Is(err, calculator.ErrInvalidExpression) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity) // 402
			} else if errors.Is(err, calculator.ErrDivisionByZero) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity) // 402
			} else {
				fmt.Fprintf(w, "unknown err")
			}

		} else {
			fmt.Fprintf(w, `{"result": %f}`, result) // 200 OK
		}
	} else {
		http.Error(w, `{"error": "Only POST method is allowed"}`, http.StatusMethodNotAllowed) // 405
	}

}

func StartServer() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Printf("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
