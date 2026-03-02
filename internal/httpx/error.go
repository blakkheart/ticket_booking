package httpx

import (
	"net/http"
	"ticket-booking/internal/user"
)

type HTTPError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

var errorMap = map[error]HTTPError{
	user.ErrEmailAlreadyUsed: {
		Status:  http.StatusConflict,
		Message: "email already used",
		Code:    "EMAIL_ALREADY_USED",
	},
	user.ErrInvalidCredentials: {
		Status:  http.StatusBadRequest,
		Message: "invalid input",
		Code:    "INVALID_INPUT",
	},
}

func ResolveHTTPError(err error) HTTPError {
	if err == nil {
		return HTTPError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_ERROR",
			Message: "internal error",
		}
	}

	// var httpErr HTTPError
	// if errors.As(err, &httpErr) {
	// 	return httpErr
	// }

	// for target, mapped := range errorMap {
	// 	if errors.Is(err, target) {
	// 		return mapped
	// 	}
	// }

	httpErr, ok := errorMap[err]
	if ok {
		return httpErr
	} // TODO might not wrk

	return HTTPError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "internal error",
	}
}
