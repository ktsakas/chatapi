package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Channel model
type Channel struct {
	ID    string `json:"id"`
	Email string `json:"organization"`
}

// BeforeCreate sets the UUID before channel creation
func (channel *Channel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
