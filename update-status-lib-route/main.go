package main

import (
	"log"
	"net/http"
	"update-status/config"
	destnumber "update-status/controllers/dest-number"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/update-status", destnumber.UpdateStatus)

	log.Println("Server running")
	http.ListenAndServe(":2000", nil)	
}