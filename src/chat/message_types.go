package chat

import "encoding/json"

type QueuePositionMessage struct {
	msgType  string `json:"type"`
	position int
}

type UserMessage struct {
	msgType string `json:"type"`
	channel int
	text    string
}

// GetMessageType returns the type of the message
func GetMessageType(jsonStr []byte) string {
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonStr, &jsonMap)
	return jsonMap["type"].(string)
}
