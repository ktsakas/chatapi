package test

import (
	"net/http"
	"testing"

	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

// Valid user id for testing
var testUserID string

func TestBadCreateUserRequest(t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Post("http://localhost:8080/user").
		Type("form").
		SendMap(map[string]string{
			"email":    "user@test.com",
			"password": "testing9",
		}).
		End()

	// Check for bad request status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("POST: /user should have bad request status but did not")
	}
}

func TestValidCreateUserRequest(t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Post("http://localhost:8080/user").
		Type("form").
		SendMap(map[string]string{
			"email":      "user@test.com",
			"password":   "testing9",
			"sex":        "male",
			"talking_to": "girls",
		}).
		End()

	// Check for valid status code
	if resp.StatusCode != http.StatusOK {
		t.Error("POST: /user failed with status ", resp.StatusCode)
	}

	// Decode json
	var jsonMap map[string]string
	var dec = json.NewDecoder(resp.Body)
	dec.Decode(&jsonMap)

	testUserExists(jsonMap["id"], t)
}

// Test that a user with the given ID exists.
func testUserExists(userID string, t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Get("http://localhost:8080/user/" + userID).End()

	if resp.StatusCode != http.StatusOK {
		t.Error("User with id " + userID + " does not exist")
	}
}
