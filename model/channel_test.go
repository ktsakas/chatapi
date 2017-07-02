package model

import (
	"os"
	"testing"

	"../migration"
)

func TestMain(m *testing.M) {
	Connect("collegechat")

	migration.Rebuild()

	retCode := m.Run()

	os.Exit(retCode)
}

func TestChannelByDomain(t *testing.T) {
	var err error
	var channel = &Channel{
		Name:    "Test Channel",
		Domain:  "test.edu",
		IsGroup: true,
	}

	err = channel.Create()
	if err != nil {
		t.Fatal("Failed to create test channel.", err)
	}

	channel, err = ChannelByDomain("test.edu")
	if err != nil {
		t.Fatal("Could not find previously created channel by domain.")
	}
}

func TestFindPrivateChannel(t *testing.T) {
	var err error
	var userA = &User{
		Email:      "test_userA@test.com",
		Password:   "testing9",
		Sex:        "boy",
		TalkingTo:  "girls",
		University: "Brown University",
	}
	err = userA.Create()
	if err != nil {
		t.Error("Failed to create user.", err)
		return
	}

	var userB = &User{
		Email:      "test_userB@test.com",
		Password:   "testing9",
		Sex:        "girl",
		TalkingTo:  "boys",
		University: "Brown University",
	}
	err = userB.Create()
	if err != nil {
		t.Fatal("Failed to create user.", err)
	}

	var newChannel = &Channel{
		Name:    "Private Channel Between 2 Users",
		Domain:  "",
		IsGroup: false,
		Members: []User{*userA, *userB},
	}
	err = newChannel.Create()

	if err != nil {
		t.Fatal("Failed to create channel.", err)
	}

	var privateChannel, findErr = FindPrivateChannel(userA, userB)
	if findErr != nil {
		t.Fatal("Failed to find the private channel that was created.", findErr)
	}

	if newChannel.ID != privateChannel.ID ||
		newChannel.Name != privateChannel.Name ||
		newChannel.Domain != privateChannel.Domain {
		t.Fatal("Private channel found does not match.")
	}
}
