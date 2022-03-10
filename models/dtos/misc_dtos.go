package dtos

import "errors"

// ErrorDto :
type ErrorDto struct {
	Message string `json:"message"`
}

func (e *ErrorDto) Error() error {
	return errors.New(e.Message)
}

func NewErrorDto(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
	}
}
