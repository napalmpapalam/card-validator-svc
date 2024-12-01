package problems

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s\n", e.Code, e.Message)
}

func BadRequest(err error) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}
