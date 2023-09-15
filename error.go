package lastfm

import (
	"fmt"
	"strings"
)

// APIError is an error sent by the JSON API.
type APIError struct {
	Message    string `json:"message"`
	Code       int    `json:"error"`
	HTTPStatus int    `json:"-"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	var sb strings.Builder

	if e.Code != 0 {
		sb.WriteString(
			fmt.Sprintf("%d", e.Code),
		)
	}

	if e.Message != "" {
		if e.Code != 0 {
			sb.WriteString(": ")
		}

		sb.WriteString(e.Message)
	}

	return sb.String()
}
