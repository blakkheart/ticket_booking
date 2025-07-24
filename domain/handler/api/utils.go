package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, model any) {
	jData, err := json.Marshal(model)
	if err != nil {
		log.Fatal("Error in struct")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
