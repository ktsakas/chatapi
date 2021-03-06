package chat

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"../model"

	"encoding/json"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// User that the client corresponds to.
	user *model.User

	// Channel id to room mapping.
	// Client keeps it's own mapping to avoid connecting to room that is not authorized.
	rooms map[string]*Room

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readFromWS pumps messages from the websocket connection to the hub.
//
// The application runs readFromWS in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (client *Client) readFromWS() {
	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()
	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.conn.ReadMessage()
		// println("got new message!")
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		msgType, err := GetMessageType(message)
		if err != nil {
			client.send <- []byte(`{ "type": "error", "code": 400, "message": "bad message" }`)
			continue
		}

		if msgType == UserMsgType {
			var userMessage = UserMessage{}
			json.Unmarshal(message, &userMessage) // TODO: not needed

			if room, ok := client.rooms[userMessage.Channel]; ok {
				room.broadcast <- message
			} else {
				client.send <- []byte(`{ "type": "error", "code": 400, "message": "channel does not exist" }`)
			}
		}
	}
}

// writeToWS pumps messages from the hub to the websocket connection.
//
// A goroutine running writeToWS is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (client *Client) writeToWS() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// JoinRoom adds the user to the room with roomId
func (client *Client) JoinRoom(roomID string, room *Room) {
	client.rooms[roomID] = room

	client.send <- []byte(`{ "type": "joinedChannel", "channel": ` + roomID + ` }`)
}

// NewClient creates a new instance of a client struct
// for a given hub and user.
func (hub *Hub) NewClient(user *model.User) *Client {
	return &Client{
		hub:   hub,
		user:  user,
		rooms: make(map[string]*Room),
		conn:  nil, // The connection is set later.
		send:  make(chan []byte, 256),
	}
}

// Serve handles websocket requests from the peer.
func (client *Client) Serve(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client.conn = conn
	// TODO: consider case of multiple clients per user
	client.hub.register <- client
	client.hub.searchMatch <- client

	go client.writeToWS()
	client.readFromWS()
}
