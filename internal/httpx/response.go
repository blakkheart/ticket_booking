package httpx

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
