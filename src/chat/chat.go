package chat

import (
	"github.com/googollee/go-socket.io"
)

var server *socketio.Server

type chat struct {
	queuedChannels []int
	orgChannels    []int
}

// New starts accepting websocket connections for chatting
func New() *socketio.Server {
	var err error
	server, err = socketio.NewServer([]string{"websocket"})

	if err != nil {
		panic(err)
	}

	server.On("connection", func(so socketio.Socket) {
		println("test")
		so.Join("global")

		so.On("joinQueue", func(so socketio.Socket) {

		})

		so.On("global message", func(so socketio.Socket) {
			println("gotmessage")
			// so.BroadcastTo("global", "message", "What is this?")
		})
	})

	return server
}
