package server

import (
	"net/http"

	"ticket-booking/domain/handler/api/helper"
)

func RegisterRoutes(mux *http.ServeMux, prefix string, routes ...func(*helper.Router)) {
	router := helper.NewRouter(mux, prefix)

	for i := 0; i < len(routes); i++ {
		router.Include(routes[i])
	}

}
