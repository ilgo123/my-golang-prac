package main

import (
	"log"
	"net/http"
	"web-native/config"
	"web-native/controllers/categorycontroller"
	homecontroller "web-native/controllers/homeController"
	"web-native/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running")
	http.ListenAndServe(":2000", nil)
}