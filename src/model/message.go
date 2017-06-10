package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Message model
type Message struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel"`
	Content   string `json:"content"`
	CreatedAt time.Time
}

// BeforeCreate sets the UUID before message creation
func (message *Message) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
