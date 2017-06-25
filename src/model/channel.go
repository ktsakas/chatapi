package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Channel model
type Channel struct {
	ID      string `json:"id"`
	IsGroup string `json:"isGroup"`
	Name    string `json:"name"`
	Domain  string `json:"domain"`
}

// BeforeCreate sets the UUID before channel creation
func (channel *Channel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
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

// PrivateChannelFromUsers finds a private channel between two users.
func PrivateChannelFromUsers() {

}
