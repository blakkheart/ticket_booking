package httpx

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	Data() any
	Status() int
	WriteJson(w http.ResponseWriter) error
}

func (r *response) Data() any {
	return r.data
}

func (r *response) Status() int {
	return r.status
}

type response struct {
	data   any
	status int
}

func NewResponse(data any, status int) Response {
	return &response{
		data:   data,
		status: status,
	}
}

func (r *response) WriteJson(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)

	return json.NewEncoder(w).Encode(r.data)
}
