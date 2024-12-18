package calculator

import "errors"

var (
	ErrInvalidExpression     = errors.New(`json:"error: invalid expression"`)
	ErrDivisionByZero        = errors.New(`json:"error: division by zero"`)
	ErrMismatchedParentheses = errors.New(`json:"error: mismatched parentheses"`)
	ErrInvalidCharacter      = errors.New(`json:"error: invalid character"`)
)
