package chat

import "encoding/json"

// Define message types
const (
	UserMsgType          = "userMessage"
	QueuePositionMsgType = "queuePosition"
	AuthMsgType          = "authorization"
)

// QueuePositionMessage is a notification of the queue position sent to the user.
type QueuePositionMessage struct {
	Type     string `json:"type"`
	Position int    `json:"position"`
}

// ConnectMessage is a notification that a user has joined a channel.
type ConnectMessage struct {
	Type    string `json:"connection"`
	Channel int    `json:"channel"`
}

// UserMessage is a chat message sent by a user.
type UserMessage struct {
	Type    string `json:"type"`
	Channel int    `json:"channel"`
	Text    string `json:"text"`
}

// AuthMessage is an authentication message
// sent right after a connection is established.
type AuthMessage struct {
	Type     string `json:"type"`
	JwtToken string `json:"jwt_token"`
}

// GetMessageType returns the type of the message
func GetMessageType(jsonStr []byte) string {
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonStr, &jsonMap)
	return jsonMap["type"].(string)
}

// MessageFromString returns a message struct from a json string
func MessageFromString(jsonStr []byte) {
	return
}
