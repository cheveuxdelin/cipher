package errors

import "errors"

var (
	// received empty string
	ErrEmptyString error = errors.New("empty string")
	//Character not valid or not ASCII
	ErrStringNotASCII error = errors.New("string not ASCII")
	ErrBlankString    error = errors.New("blank string")
)
