package httpx

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	Data() any
	Status() int
}

func (r *response) Data() any {
	return r.Data
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

func WriteJsonResponse(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
