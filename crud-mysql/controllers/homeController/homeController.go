package homecontroller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"web-native/entities"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

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
