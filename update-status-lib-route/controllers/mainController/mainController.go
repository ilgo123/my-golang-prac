package maincontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"update-status/entities"
)

func Response(w http.ResponseWriter, statusCode int, response entities.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		return
	}

	w.Write(jsonResponse)
}