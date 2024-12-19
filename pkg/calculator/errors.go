package calculator

import "errors"

var (
	ErrInvalidExpression = errors.New(`{"error": "invalid expression"}`)
	ErrDivisionByZero    = errors.New(`{"error": "division by zero"}`)
	ErrInvalidCharacter  = errors.New(`{""error": "invalid character"}`)
)
