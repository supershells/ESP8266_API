package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewInternalServerError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewConflictError(message string) error {
	return AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

func NewForbiddenError(message string) error {
	return AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewNotImplementedError(message string) error {
	return AppError{
		Code:    http.StatusNotImplemented,
		Message: message,
	}
}

func NewServiceUnavailableError(message string) error {
	return AppError{
		Code:    http.StatusServiceUnavailable,
		Message: message,
	}
}

func NewGatewayTimeoutError(message string) error {
	return AppError{
		Code:    http.StatusGatewayTimeout,
		Message: message,
	}
}

func NewBadGatewayError(message string) error {
	return AppError{
		Code:    http.StatusBadGateway,
		Message: message,
	}
}

func NewTooManyRequestsError(message string) error {
	return AppError{
		Code:    http.StatusTooManyRequests,
		Message: message,
	}
}

func NewInsufficientStorageError(message string) error {
	return AppError{
		Code:    http.StatusInsufficientStorage,
		Message: message,
	}
}
