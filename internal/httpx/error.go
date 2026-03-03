package httpx

import (
	"errors"
	"net/http"
)

type HTTPError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Err     error
}

func (e *HTTPError) Error() string {
	return e.Message
}

var (
	ErrInvalidRequestBody    error
	ErrInternalServerProblem error
)

var errorMap = map[error]HTTPError{
	ErrInvalidRequestBody: {
		Status:  http.StatusBadRequest,
		Message: "invalid request body",
		Code:    "INVALID_REQUEST_BODY",
	},
	ErrInternalServerProblem: {
		Status:  http.StatusInternalServerError,
		Message: "internal error",
		Code:    "INTERNAL_ERROR",
	},
}

func ResolveHTTPError(err error) *HTTPError {
	if err == nil {
		return &HTTPError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_ERROR",
			Message: "internal error",
		}
	}

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr
	}

	for target, mapped := range errorMap {
		if errors.Is(err, target) {
			return &mapped
		}
	}

	return &HTTPError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "internal error",
	}
}
