package userapi

import (
	"errors"
	"net/http"
	"ticket-booking/internal/httpx"
	"ticket-booking/internal/user"
)

var errorMap = map[error]httpx.HTTPError{
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

func ResolveHTTPError(err error) error {
	for target, mapped := range errorMap {
		if errors.Is(err, target) {
			return &mapped
		}
	}

	return err
}
