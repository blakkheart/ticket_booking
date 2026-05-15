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

func (r *Router) Handle(method string, path string, handler AppHandler) {
	r.mux.Handle(
		fmt.Sprintf("%s %s%s", method, r.prefix, path),

		Adapt(handler),
	)
}

func (r *Router) Include(fn func(*Router)) {
	fn(r)
}

func RegisterRoutes(
	mux *http.ServeMux,
	prefix string,
	routes ...func(*Router),
) {
	router := NewRouter(mux, prefix)

	for _, r := range routes {
		router.Include(r)
	}

}

type AppHandler func(http.ResponseWriter, *http.Request) (Response, error)

const ErrorKey contextKey = "handler_error"

func Adapt(h AppHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := h(w, r)
		if err != nil {

			httpErr := ResolveHTTPError(err)
			errResp := NewResponse(map[string]string{
				"code":    httpErr.Code,
				"message": httpErr.Message,
			}, httpErr.Status)
			errResp.WriteJson(w)

			return
		}

		resp.WriteJson(w)

	})
}
