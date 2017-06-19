package chat

import (
	"../model"

	"container/list"
	"encoding/json"
)

// Hub is a middleman between incoming connections
// and chat rooms.
type Hub struct {
	// Set of clients that have not been paired with someone to talk.
	unpaired *list.List

	// Map clients to users
	users map[*Client]*model.User

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from the clients.
	unregister chan *Client

	// SearchMarch request to get matched to a user.
	searchMatch chan *Client
}

// New creates a new chat application instance.
func New() *Hub {
	var newHub = &Hub{
		unpaired:    list.New(),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		searchMatch: make(chan *Client),
	}

	go newHub.run()
	return newHub
}

func sameInterests(userA *model.User, userB *model.User) bool {
	return (userA.TalkingTo == userB.TalkingTo && userA.Sex == userB.Sex)
}

func isMatch(userA *model.User, userB *model.User) bool {
	return (userA.TalkingTo == userB.Sex && userA.Sex == userB.TalkingTo)
}

// findMatch finds another client to pair with or returns nil
func (hub *Hub) findMatch(client *Client) *list.Element {
	var user = hub.users[client]
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var possibleMatch = hub.users[e.Value.(*Client)]

		if isMatch(user, possibleMatch) {
			return e
		}
	}

	return nil
}

func (hub *Hub) getPositionInQueue(client *Client) int {
	var user = hub.users[client]
	var queueNum = 1
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var competitor = hub.users[e.Value.(*Client)]

		if sameInterests(user, competitor) {
			queueNum++
		}
	}

	return queueNum
}

func (hub *Hub) updateQueuePositions(paired *Client) {
	var pairedUser = hub.users[paired]
	var queueNum = 1
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var queuedClient = e.Value.(*Client)
		var queuedUser = hub.users[queuedClient]

		if sameInterests(pairedUser, queuedUser) {
			// Update queue number
			var json, _ = json.Marshal(QueuePositionMessage{
				Type:     "queuePosition",
				Position: queueNum,
			})

			queuedClient.send <- json
			queueNum++
		}
	}
}

func (hub *Hub) run() {
	var client = <-hub.searchMatch
	var matchE = hub.findMatch(client)

	if matchE != nil {
		// Found a match.
		var match = matchE.Value.(*Client)
		hub.unpaired.Remove(matchE)
		hub.updateQueuePositions(match)

	} else {
		// Could not find a match, put the client in the queue.
		hub.unpaired.PushBack(client)

		// Notify the client of its position in the queue
		var queueNum = hub.getPositionInQueue(client)
		var json, _ = json.Marshal(QueuePositionMessage{
			Type:     "queuePosition",
			Position: queueNum,
		})
		client.send <- json
	}
}
