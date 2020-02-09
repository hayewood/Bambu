package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

// Configure the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Payload
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	for {
		var aMsg Payload
		// Read in a new message as JSON and map it to a Message object

		err := c.ws.ReadJSON(&aMsg)
		if err != nil {
			log.Printf("error reading json: %v", err)
			delete(h.rooms, s.room)
			return
		}

		log.Printf("Message: " + aMsg.Message)

		// Send the newly received message to the broadcast channel
		msg := message{payload: aMsg, room: s.room}
		prettyPrintJSON(aMsg)
		h.broadcast <- msg
	}
}

// write writes a message with the given message type and payload.
func (c *connection) WriteJSON(payload Payload) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteJSON(payload)

}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case Message, ok := <-c.send:
			if !ok {
				c.WriteJSON(Payload{})
				return
			}
			if err := c.ws.WriteJSON(Message); err != nil {
				log.Printf("error: %v", err)
				c.ws.Close()
				delete(h.rooms, s.room)
				return
			}
		case <-ticker.C:
			if err := c.WriteJSON(Payload{}); err != nil {
				return
			}
		}
	}
}

func prettyPrintJSON(amessage Payload) {
	prettyJSON, err := json.MarshalIndent(amessage, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	vars := mux.Vars(r)
	c := &connection{send: make(chan Payload), ws: ws}

	if vars == nil {
		s := subscription{c, ""}
		h.register <- s
		go s.writePump()
		s.readPump()

	} else {
		log.Println("Room:" + vars["room"])
		s := subscription{c, vars["room"]}
		h.register <- s
		go s.writePump()
		s.readPump()
	}

}
