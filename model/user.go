package model

import (
	"time"

	"../config"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Sex        string    `json:"sex"`
	TalkingTo  string    `json:"talkingTo"`
	University string    `json:"university"`
	Contacts   []User    `json:"contacts"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// BeforeCreate sets the UUID before user creation
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Create adds the user struct in the database
// and returns the stored record.
func (user *User) Create() error {
	var err = db.Create(user).Error

	if err != nil {
		var _, lookupErr = UserByEmail(user.Email)

		if lookupErr == nil {
			return config.ErrRecordExists
		}

		return err
	}

	return nil
}

// Update updates a user record
func (user *User) Update() error {
	return db.Update(&user).Error
}

// UserByID finds the user given his id.
func UserByID(id string) (*User, error) {
	var user = User{
		ID: id,
	}

	var err = db.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByEmail finds the user given his email.
func UserByEmail(email string) (*User, error) {
	var user = User{
		Email: email,
	}

	var err = db.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ValidateUserCredentials checks the username and password against the databse
func ValidateUserCredentials(email, password string) (*User, bool) {
	var user = User{}
	var recordNotFound = db.Where("email = ? AND password = ?", email, password).First(&user).RecordNotFound()

	if recordNotFound {
		return nil, false
	}

	return &user, true
}

// Delete deletes the given user
// ** this should only be used for testing purposes **
func (user *User) Delete() error {
	if config.Get("Environment") != "DEV" {
		panic("User.Delete should never be called in production!")
	}

	return db.Delete(user).Error
}
