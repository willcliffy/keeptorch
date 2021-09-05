package transport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/jsonapi"
)

type Success struct {
	Data     interface{} `json:"data"` 
}

func SendResponseToClient(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.Header().Set("Cache-Control", "private, max-age=60")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(Success{Data: data}); err != nil {
		log.Println("error marshalling payload: ", err)
	}
}
