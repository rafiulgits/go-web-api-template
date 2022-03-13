package dtos

import (
	"errors"
	"webapi/models"

	"gorm.io/gorm"
)

type ErrorDto struct {
	Message   string      `json:"message"`
	Fields    interface{} `json:"fields,omitempty"`
	ErrorCode int         `json:"-"` //HTTP Status Code to indicate error type
}

func NewErrorDto(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
	}
}

func NewValidationError(err error) *ErrorDto {
	return &ErrorDto{
		Message:   "data validation failed",
		Fields:    models.ParseModelErrors(err),
		ErrorCode: 400,
	}
}

func NewDatabaseError(err error) *ErrorDto {
	code := getErrorDatabaseErrorCode(err)
	var msg string
	switch code {
	case 404:
		msg = "no such record found"
	default:
		msg = "something went wrong while querying database"
	}
	return &ErrorDto{
		Message:   msg,
		ErrorCode: getErrorDatabaseErrorCode(err),
	}
}

func NewDatabaseErrorWithMessage(err error, message string) *ErrorDto {
	return &ErrorDto{
		Message:   message,
		ErrorCode: getErrorDatabaseErrorCode(err),
	}
}

func getErrorDatabaseErrorCode(err error) int {
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return 404
	}
	return 500
}
