package helper

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

// type handlerFunc func(w http.ResponseWriter, r *http.Request)

// func addHandler(router *http.ServeMux, method string, path string, handlerFunc handlerFunc) {
// 	router.HandleFunc(fmt.Sprintf("%s %s", method, path), func(w http.ResponseWriter, r *http.Request) { handlerFunc(w, r) })
// }
