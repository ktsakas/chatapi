package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Channel model
type Channel struct {
	ID        string    `json:"id"`
	IsGroup   bool      `json:"isGroup"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"createdAt"`

	Members []User `gorm:"many2many:channel_users;"`
}

// BeforeCreate sets the UUID before channel creation
func (channel *Channel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

// Create stores a new channel in the database
func (channel *Channel) Create() error {
	var err = db.Create(&channel).Error

	if err != nil {
		return err
	}

	return nil
}

// FindPrivateChannel finds a private channel between two users if one exists.
func FindPrivateChannel(userA *User, userB *User) (*Channel, error) {
	var privateChannel *Channel
	// Find private channel where both userA and userB belong
	var err = db.Raw(
		`SELECT channels.*
		FROM channels, channel_users as memberA, channel_users as memberB
		WHERE memberA.user_id = ?
		AND memberB.user_id = ?
		AND memberA.channel_id = memberB.channel_id
		AND channels.id = memberA.channel_id
		AND channels.is_group = false`, userA.ID, userB.ID).Scan(&privateChannel).Error

	if err != nil {
		return nil, err
	}

	return privateChannel, nil
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
