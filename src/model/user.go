package model

import (
	"time"

	"github.com/jinzhu/gorm"
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
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Create adds the user struct in the database
// and returns the stored record.
func (user *User) Create() error {
	return db.Create(&user).Error
}

// UserByID finds the user given his id.
func UserByID(id string) User {
	var user = User{
		ID: id,
	}

	db.First(&user)
	return user
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
