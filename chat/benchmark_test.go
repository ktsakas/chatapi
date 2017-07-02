package chat

import (
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func websocketClient(reciever chan []byte, b *testing.B) {
	// Connect to websocket
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/chat/ceee16ac-3111-4efe-b61b-cfbe4563eeaf", nil)
	if err != nil {
		b.Fatal("Could not connect to socket: ", err)
	}
	defer conn.Close()
	defer close(reciever)

	// Channels for function exit or interupt signal
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer conn.Close()
		defer close(done)
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				b.Fatal("Response read error: ", err)
				return
			}
			// log.Printf("Server response: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case message := <-reciever:
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				b.Fatal("Failed to write message:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt ")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// TODO: write benchmark between two clients
func BenchmarkClient(b *testing.B) {
	var reciever = make(chan []byte)
	go websocketClient(reciever, b)

	for n := 0; n < b.N; n++ {
		reciever <- []byte("bad message")
	}
}
