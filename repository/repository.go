package repository

import (
	"RainerGevers/go-clean-architecture-chatGPT/models"
	"log"
)

func GetUsers() []models.User {
	// Connect to the database
	db := connectToDB()

	// Query the database for the list of users
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the rows and create a slice of user structs
	var users []models.User
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal(err)
		}
		users = append(users, models.User{ID: id, Name: name, Email: email})
	}

	return users
}

func GetUser(id int) models.User {
	// Connect to the database
	db := connectToDB()

	// Query the database for the user
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)

	// Scan the row into a user struct
	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		log.Fatal(err)
	}

	return user
}

func CreateUser(user models.User) {
	// Connect to the database
	db := connectToDB()

	// Insert the user into the database
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		log.Fatal(err)
	}
}
