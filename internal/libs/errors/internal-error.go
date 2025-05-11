package errors

import (
	"errors"
	"log/slog"
	"net/http"
)

const (
	InternalErrorMessage = "service temporary unavailable"
)

type internalAPIError struct {
	error
	debugMessage string
	userMessage  string
	status       int
}

func (e *internalAPIError) Unwrap() error {
	return e.error
}

func (e *internalAPIError) Error() string {
	errMessage := e.debugMessage

	if e.status != 0 {
		errMessage += " (returned " + http.StatusText(e.status) + " with reason " + e.userMessage + ")"
	}

	if e.error != nil {
		errMessage += " called by: " + e.error.Error()
	}

	return errMessage
}

func (e *internalAPIError) UserMessage() (message string) {
	return e.userMessage
}

func (e *internalAPIError) Status() (status int) {
	return e.status
}

func New(errs error, debugMessage, userMessage string, status int) error {
	return &internalAPIError{error: errs, debugMessage: debugMessage, userMessage: userMessage, status: status}
}

func NewError(l *slog.Logger, status int, debugMessage, userMessage string, errs ...error) error {
	err := errors.Join(errs...)
	l.Debug(debugMessage, slog.Any("err", err))

	return &internalAPIError{error: err, debugMessage: debugMessage, userMessage: userMessage, status: status}
}

func NewInternalError(l *slog.Logger, message string, errs ...error) error {
	err := errors.Join(errs...)
	l.Error(message, slog.Any("err", err))

	return &internalAPIError{error: err, debugMessage: message, userMessage: InternalErrorMessage, status: http.StatusInternalServerError}
}

func NewValidationError(l *slog.Logger, internalMessage, userMessage string, errs ...error) error {
	return NewError(l, http.StatusBadRequest, internalMessage, userMessage, errs...)
}

func NewNotFoundError(l *slog.Logger, internalMessage, userMessage string, errs ...error) error {
	return NewError(l, http.StatusNotFound, internalMessage, userMessage, errs...)
}

func NewUnauthorizedError(l *slog.Logger, internalMessage, userMessage string, errs ...error) error {
	return NewError(l, http.StatusUnauthorized, internalMessage, userMessage, errs...)
}
