package test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

func TestBadCreateUserRequest(t *testing.T) {

}

func TestValidCreateUserRequest(t *testing.T) {
	request := gorequest.New()
	resp, body, _ := request.Post("http://localhost:8080/user").
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

	print(body)
	checkLowercaseJSONKeys(resp.Body, t)
}

// Enusre that all keys in the returned json object are lowercase.
func checkLowercaseJSONKeys(jsonStr io.Reader, t *testing.T) {
	var jsonMap map[string]string
	var dec = json.NewDecoder(jsonStr)
	dec.Decode(&jsonMap)

	for key := range jsonMap {
		if key != strings.ToLower(key) {
			t.Error("Invalid JSON key ", key)
		}
	}
}
