package httpx

import (
	"context"
	"fmt"
	"net/http"
)

type Router struct {
	mux    *http.ServeMux
	prefix string
}

func NewRouter(mux *http.ServeMux, prefix string) *Router {
	return &Router{mux: mux, prefix: prefix}
}

func (r *Router) Handle(method string, path string, handler AppHandler) {
	r.mux.Handle(
		fmt.Sprintf("%s %s%s", method, r.prefix, path),

		Adapt(handler),
	)
}

func (r *Router) Include(fn func(*Router)) {
	fn(r)
}

func RegisterRoutes(mux *http.ServeMux, prefix string, routes ...func(*Router)) {
	router := NewRouter(mux, prefix)

	for i := 0; i < len(routes); i++ {
		router.Include(routes[i])
	}

}

type AppHandler func(http.ResponseWriter, *http.Request) error

const ErrorKey contextKey = "handler_error"

func Adapt(h AppHandler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(11)
			if err := h(w, r); err != nil {
				ctx := context.WithValue(r.Context(), ErrorKey, err)
				fmt.Println(ctx)
				*r = *r.WithContext(ctx)
			}
		},
	)
}
