package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// User model
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Sex        string `json:"sex"`
	TalkingTo  string `json:"talking_to"`
	University string `json:"university"`
	Contacts   []User `json:"contacts"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// BeforeCreate sets the UUID before user creation
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
