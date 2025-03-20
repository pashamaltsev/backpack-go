package backpackgo

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAPIKeyOrSecret = errors.New("invalid API key or secret")
)

type BackpackError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *BackpackError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}
