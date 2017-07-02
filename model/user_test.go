package model

import (
	"testing"
)

// Generates a new user and stores him in the database
// if the user could not be created the underlying test fails
func generateUser(t *testing.T) *User {
	var newUser = &User{
		Email:      "konstantinos_tsakas@brown.edu",
		Password:   "kostakis74757",
		Sex:        "boy",
		TalkingTo:  "girls",
		University: "Brown University",
	}

	var err = newUser.Create()
	if err != nil {
		t.Fatal("Failed to create test user.", err)
	}

	return newUser
}

func TestUserByEmail(t *testing.T) {
	var newUser = generateUser(t)
	var user, err = UserByID(newUser.ID)
	if err != nil {
		t.Fatal("Could not find previously created user by email.")
	} else if newUser.ID != user.ID {
		t.Fatal("The user ID found is different from the one created.")
	}
}

func TestUserByID(t *testing.T) {
	var newUser = generateUser(t)
	var user, err = UserByEmail(newUser.Email)
	if err != nil {
		t.Fatal("Could not find previously created user by id.")
	} else if newUser.ID != user.ID {
		t.Fatal("The user ID found is different from the one created.")
	}
}
