package auth

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
)

func Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/login", login)
	r.Handle(http.MethodPost, "/auth", authorize)
}
