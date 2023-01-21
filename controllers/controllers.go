package controllers

import (
	"RainerGevers/go-clean-architecture-chatGPT/models"
	"RainerGevers/go-clean-architecture-chatGPT/usecases"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get the list of users from the use case
	users := usecases.GetUsers()

	// Write the users as JSON to the response
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get the user from the use case
	user := usecases.GetUser(id)

	// Write the user as JSON to the response
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a user struct
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Create the user using the use case
	usecases.CreateUser(user)

	// Write a success message to the response
	w.Write([]byte("User created"))
}
