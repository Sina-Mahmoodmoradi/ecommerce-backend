package apperrors

import (
    "errors"
    "net/http"
)

type Code string

const (
    CodeNotFound        Code = "NOT_FOUND"
    CodeEmailExists     Code = "EMAIL_EXISTS"
    CodeInvalidArgument Code = "INVALID_ARGUMENT"
    CodeInternal        Code = "INTERNAL"
)

type AppError struct {
    Code    Code
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Err
}

// Factory helpers
func New(code Code, msg string, err error) *AppError {
    return &AppError{Code: code, Message: msg, Err: err}
}

func Is(err error, code Code) bool {
    var appErr *AppError
    if errors.As(err, &appErr) {
        return appErr.Code == code
    }
    return false
}

func HTTPStatus(code Code) int {
    switch code {
    case CodeNotFound:
        return http.StatusNotFound
    case CodeEmailExists:
        return http.StatusConflict
    case CodeInvalidArgument:
        return http.StatusBadRequest
    default:
        return http.StatusInternalServerError
    }
}
