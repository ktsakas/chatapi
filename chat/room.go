package chat

// Room maintains the set of active clients and broadcasts messages to the clients.
type Room struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewRoom creates a new room with the given clients.
func NewRoom([]*Client) *Room {
	var room = &Room{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	go room.run()
	return room
}

// run starts accepting connections and messages to the room.
func (room *Room) run() {
	for {
		select {
		case client := <-room.register:
			room.clients[client] = true

		case client := <-room.unregister:
			if _, ok := room.clients[client]; ok {
				delete(room.clients, client)
				close(client.send)
			}

		case message := <-room.broadcast:
			for client := range room.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(room.clients, client)
				}
			}
		}
	}
}
