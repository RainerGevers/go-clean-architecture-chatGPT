package main

import (
	"RainerGevers/go-clean-architecture-chatGPT/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the routes and handlers
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	// Start the server
	http.ListenAndServe(":8000", router)
}
