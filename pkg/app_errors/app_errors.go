package app_errors

import "net/http"

type AppError struct {
	StatusCode int
	Message    string
	Cause      error
}

func (aE *AppError) Error() string {
	return aE.Message
}

func (aE *AppError) Unwrap() error {
	return aE.Cause
}

func ErrorNotFound(msg string) *AppError {
	return &AppError{StatusCode: http.StatusNotFound, Message: msg}
}

func ErrorConflict(msg string) *AppError {
	return &AppError{StatusCode: http.StatusConflict, Message: msg}
}

func ErrorInternal(cause error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    "erro interno no servidor",
		Cause:      cause,
	}
}

func ErrorBadRequest(msg string) *AppError {
	return &AppError{StatusCode: http.StatusBadRequest, Message: msg}
}

func ErrorUnauthorized(msg string) *AppError {
	return &AppError{StatusCode: http.StatusUnauthorized, Message: msg}
}
