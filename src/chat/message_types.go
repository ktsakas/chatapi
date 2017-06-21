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
	Position int
}

// ConnectMessage is a notification that a user has joined a channel.
type ConnectMessage struct {
	Type    string `json:"connection"`
	Channel int
}

// UserMessage is a chat message sent by a user.
type UserMessage struct {
	Type    string `json:"type"`
	Channel int
	Text    string
}

// AuthMessage is an authentication message
// sent right after a connection is established.
type AuthMessage struct {
	Type     string `json:"type"`
	JwtToken string
}

// GetMessageType returns the type of the message
func GetMessageType(jsonStr []byte) string {
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonStr, &jsonMap)
	return jsonMap["type"].(string)
}
