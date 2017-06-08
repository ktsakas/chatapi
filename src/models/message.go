package model

import "time"

// Message model
type Message struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel"`
	Content   string `json:"content"`
	CreatedAt time.Time
}
