package booking

import (
	"net/http"
	"ticket-booking/domain/handler/api/helper"
)

func Routes(r *helper.Router) {
	r.Handle(http.MethodPost, "/book", getBooking)
}
