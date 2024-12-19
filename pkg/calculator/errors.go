package calculator

import "errors"

var (
	ErrInvalidExpression = errors.New(`{"error": "Invalid expression"}`)
	ErrDivisionByZero    = errors.New(`{"error": "Division by zero"}`)
)
