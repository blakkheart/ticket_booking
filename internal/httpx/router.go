package httpx

import (
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

func (r *Router) Handle(method string, path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(fmt.Sprintf("%s %s%s", method, r.prefix, path), func(w http.ResponseWriter, r *http.Request) { handler(w, r) })
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
