package model

import (
	"strconv"
	"time"
)

// Channel model
type Channel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsGroup   bool      `json:"isGroup"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"createdAt"`

	Members []User `gorm:"many2many:channel_users;"`
}

// FindOrCreatePrivateChannel tries to find a private channel between two users
// and if no such channel exists it creates it
func FindOrCreatePrivateChannel(userA *User, userB *User) *Channel {
	// Create channel in the database if it does not exist
	var channel, err = FindPrivateChannel(userA, userB)
	if err == nil {
		return channel
	}

	channel = &Channel{
		IsGroup: false,
		Name:    userA.Email + " talking with " + userB.Email,
		Members: []User{*userA, *userB},
	}
	channel.Create()
	return channel
}

// Create stores a new channel in the database
func (channel *Channel) Create() error {
	return db.Create(&channel).Error
}

// FindPrivateChannel finds a private channel between two users if one exists.
func FindPrivateChannel(userA *User, userB *User) (*Channel, error) {
	// Find private channel where both userA and userB belong
	var userAid, _ = strconv.Atoi(userA.ID)
	var userBid, _ = strconv.Atoi(userB.ID)

	var channel Channel
	db.Raw(`SELECT channels.*
		FROM channels, channel_users as memberA, channel_users as memberB
		WHERE memberA.user_id = ?
		AND memberB.user_id = ?
		AND memberA.channel_id = memberB.channel_id
		AND channels.id = memberA.channel_id
		AND channels.is_group = false;`, userAid, userBid).Scan(&channel)

	if db.Error != nil {
		return nil, db.Error
	}

	return &channel, nil
}

// ChannelByDomain finds a channel given a domain name.
func ChannelByDomain(domain string) (*Channel, error) {
	var channel = Channel{
		Domain: domain,
	}

	var err = db.First(&channel).Error
	if err != nil {
		return nil, err
	}

	return &channel, nil
}
