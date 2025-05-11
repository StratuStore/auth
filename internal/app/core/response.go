package core

import (
	"github.com/StratuStore/auth/internal/libs/errors"
)

type Response[T any] struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	Body  T      `json:"body,omitempty"`
}

func NewResponseByUserError(err errors.UserError) *Response[any] {
	return &Response[any]{
		OK:    false,
		Error: err.UserMessage(),
	}
}

func NewErrorResponse(msg string) *Response[any] {
	return &Response[any]{
		OK:    false,
		Error: msg,
	}
}

func NewOKResponse[T any](body T) *Response[T] {
	return &Response[T]{
		OK:   true,
		Body: body,
	}
}
