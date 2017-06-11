package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// User model
type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Sex        string    `json:"sex"`
	TalkingTo  string    `json:"talking_to"`
	University string    `json:"university"`
	Contacts   []User    `json:"contacts"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
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
func (user *User) Create() {
	db.Create(&user)
}

// UserByID finds the user given his id.
func UserByID(id string) User {
	var user = User{
		ID: id,
	}

	db.First(&user)
	return user
}
