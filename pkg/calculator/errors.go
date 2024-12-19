package calculator

import "errors"

var (
	ErrInvalidExpression     = errors.New(`{"error": "invalid expression"}`)
	ErrDivisionByZero        = errors.New(`{"error": "division by zero"}`)
	ErrMismatchedParentheses = errors.New(`{"error": "mismatched parentheses"}`)
	ErrInvalidCharacter      = errors.New(`{""error": "invalid character"}`)
)
