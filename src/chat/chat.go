package chat

import (
	"../model"

	"container/list"
	"encoding/json"
)

// Chat is a middleman between incoming messages
// and chat rooms.
type Chat struct {
	// Set of clients that have not been paired with someone to talk.
	unpaired *list.List

	// Map clients to users
	users map[*Client]*model.User

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from the clients.
	unregister chan *Client
}

// New creates a new chat application instance.
func New() *Chat {
	return &Chat{
		unpaired: list.New(),
		register: make(chan *Client),
	}
}

func sameInterests(userA *model.User, userB *model.User) bool {
	return (userA.TalkingTo == userB.TalkingTo && userA.Sex == userB.Sex)
}

func isMatch(userA *model.User, userB *model.User) bool {
	return (userA.TalkingTo == userB.Sex && userA.Sex == userB.TalkingTo)
}

// findMatch finds another client to pair with or returns nil
func (chat *Chat) findMatch(client *Client) *list.Element {
	var user = chat.users[client]
	for e := chat.unpaired.Front(); e != nil; e = e.Next() {
		var possibleMatch = chat.users[e.Value.(*Client)]

		if isMatch(user, possibleMatch) {
			return e
		}
	}

	return nil
}

func (chat *Chat) getPositionInQueue(client *Client) int {
	var user = chat.users[client]
	var queueNum = 1
	for e := chat.unpaired.Front(); e != nil; e = e.Next() {
		var competitor = chat.users[e.Value.(*Client)]

		if sameInterests(user, competitor) {
			queueNum++
		}
	}

	return queueNum
}

func (chat *Chat) updateQueuePositions(paired *Client) {
	var pairedUser = chat.users[paired]
	var queueNum = 1
	for e := chat.unpaired.Front(); e != nil; e = e.Next() {
		var queuedClient = e.Value.(*Client)
		var queuedUser = chat.users[queuedClient]

		if sameInterests(pairedUser, queuedUser) {
			// Update queue number
			var json, _ = json.Marshal(QueuePositionMessage{
				msgType:  "queuePosition",
				position: queueNum,
			})

			queuedClient.send <- json
			queueNum++
		}
	}
}

func (chat *Chat) Run() {
	var client = <-chat.register
	var matchE = chat.findMatch(client)

	if matchE != nil {
		// Found a match.
		var match = matchE.Value.(*Client)
		chat.unpaired.Remove(matchE)
		chat.updateQueuePositions(match)

	} else {
		// Could not find a match, put the client in the queue.
		chat.unpaired.PushBack(client)

		// Notify the client of its position in the queue
		var queueNum = chat.getPositionInQueue(client)
		var json, _ = json.Marshal(QueuePositionMessage{
			msgType:  "queuePosition",
			position: queueNum,
		})
		queuedClient.send <- json
	}
}
