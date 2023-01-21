package usecases

import (
	"RainerGevers/go-clean-architecture-chatGPT/models"
	"RainerGevers/go-clean-architecture-chatGPT/repository"
)

func GetUsers() []models.User {
	// Call the repository to get the list of users
	users := repository.GetUsers()

	// Apply business rules to the users
	// ...

	return users
}

func GetUser(id int) models.User {
	// Call the repository to get the user
	user := repository.GetUser(id)

	// Apply business rules to the user
	// ...

	return user
}

func CreateUser(user models.User) {
	// Validate the user
	// ...

	// Call the repository to create the user
	repository.CreateUser(user)
}
