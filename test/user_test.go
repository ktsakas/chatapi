package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"../config"
	"../migration"

	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

// Valid user id for testing
var testUserID string
var appURL = "http://localhost:" + config.Get("DevPort")

func TestMain(m *testing.M) {
	migration.Rebuild()

	retCode := m.Run()

	os.Exit(retCode)
}

func requestFailMsg(route string, resp gorequest.Response) string {
	return fmt.Sprintf("%s: %s failed with status %d, BODY: %s", resp.Request.Method, route, resp.StatusCode, resp.Body)
}

func TestBadCreateUserRequest(t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Post(appURL + "/user").
		Type("form").
		SendMap(map[string]string{
			"email":    "user@test.com",
			"password": "testing9",
		}).
		End()

	// Check for bad request status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("POST: /user should have bad request status but did not")
	}
}

func TestValidCreateUserRequest(t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Post(appURL + "/user").
		Type("form").
		SendMap(map[string]string{
			"email":     "user@test.com",
			"password":  "testing9",
			"sex":       "male",
			"talkingTo": "girls",
		}).
		End()

	// Check for valid status code
	if resp.StatusCode != http.StatusOK {
		t.Fatal(requestFailMsg("/user", resp))
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
	resp, _, _ := request.Get(appURL + "/user/" + userID).End()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("User with id " + userID + " does not exist")
	}
}

func TestUserLogin(t *testing.T) {
	request := gorequest.New()
	resp, _, _ := request.Post(appURL + "/user").
		Type("form").
		SendMap(map[string]string{
			"email":     "user@login.com",
			"password":  "loginpass9",
			"sex":       "male",
			"talkingTo": "girls",
		}).
		End()

	// Check for valid status code
	if resp.StatusCode != http.StatusOK {
		t.Fatal(requestFailMsg("/user", resp))
	}

	var loginReq = gorequest.New()
	loginResp, _, _ := loginReq.Post(appURL + "/login").
		Type("form").
		SendMap(map[string]string{
			"email":    "test@login.com",
			"password": "loginpass9",
		}).
		End()

	// Check for valid login
	if loginResp.StatusCode != http.StatusOK {
		t.Fatal(requestFailMsg("/user", resp))
	}

	// TODO: test that we got token back
}
