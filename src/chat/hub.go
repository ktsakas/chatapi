package chat

import (
	"strings"

	"../model"

	"container/list"
	"encoding/json"
)

// Hub is a middleman between incoming connections
// and chat rooms.
type Hub struct {
	// Map channels to rooms.
	rooms map[*model.Channel]*Room

	// Set of clients that have not been paired with someone to talk.
	unpaired *list.List

	// Register requests from the clients (ie. when the user opens the app).
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
		rooms:       make(map[*model.Channel]*Room),
	}

	go newHub.run()
	return newHub
}

func sameInterests(clientA *Client, clientB *Client) bool {
	return (clientA.user.TalkingTo == clientB.user.TalkingTo && clientA.user.Sex == clientB.user.Sex)
}

func isMatch(userA *model.User, userB *model.User) bool {
	return (userA.TalkingTo == userB.Sex && userA.Sex == userB.TalkingTo)
}

// findMatch finds another client to pair with or returns nil
func (hub *Hub) findMatch(client *Client) *list.Element {
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var possibleMatch = e.Value.(*Client).user

		if isMatch(client.user, possibleMatch) {
			return e
		}
	}

	return nil
}

// Gets the current position of the given client in the queue
func (hub *Hub) getPositionInQueue(client *Client) int {
	var queueNum = 1
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var competitor = e.Value.(*Client)

		if sameInterests(client, competitor) {
			queueNum++
		}
	}

	return queueNum
}

// Updates the position of all clients when a new client is paired
func (hub *Hub) updateQueuePositions(paired *Client) {
	var queueNum = 1
	for e := hub.unpaired.Front(); e != nil; e = e.Next() {
		var queuedClient = e.Value.(*Client)

		if sameInterests(paired, queuedClient) {
			// Update queue position
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
	for {
		select {
		case client := <-hub.register:
			var domain = strings.Split(client.user.Email, "@")[1]
			var channel, err = model.ChannelByDomain(domain)

			if err != nil {
				// Critical this should never happen
				continue
			}

			// Create room if it does not exist
			var room, ok = hub.rooms[channel]
			if !ok {
				room = NewRoom([]*Client{client})
				hub.rooms[channel] = room
			}

			room.register <- client

		case client := <-hub.unregister:
			// Remove from unpaired users if the client is there
			for e := hub.unpaired.Front(); e != nil; e = e.Next() {
				var unpairedClient = e.Value.(*Client)
				if client == unpairedClient {
					hub.unpaired.Remove(e)
					break
				}
			}

		case client := <-hub.searchMatch:
			var matchE = hub.findMatch(client)

			if matchE != nil {
				// Found a match.
				var match = matchE.Value.(*Client)
				hub.unpaired.Remove(matchE)
				hub.updateQueuePositions(match)

				var channel = model.FindOrCreatePrivateChannel(client.user, match.user)
				hub.rooms[channel] = NewRoom([]*Client{client, match})

			} else {
				var queueNum = hub.getPositionInQueue(client)

				// Put the client in the queue of unpaired clients.
				hub.unpaired.PushBack(client)
				// Notify the client of its position in the queue
				var json, _ = json.Marshal(QueuePositionMessage{
					Type:     "queuePosition",
					Position: queueNum,
				})

				client.send <- json
			}
		}
	}
}
