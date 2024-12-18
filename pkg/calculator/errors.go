package calculation

import "errors"

var (
	InvalidExpression = errors.New("invalid expression")
	DivisionByZero    = errors.New("division by zero")
)
