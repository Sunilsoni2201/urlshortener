package errors

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func NewError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError, // default error code is Internal Server Error
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

func NewInvalidInputError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusUnprocessableEntity,
	}
}
