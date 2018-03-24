package client

import (
	"errors"
	"log"

	god "github.com/JanUrb/godathon2018"
	"github.com/gorilla/websocket"
)

var (
	ErrWrite error = errors.New("ErrWrite")
)

type Client struct {
	protocol god.Protocol
	switcher god.Switching
	conn     *websocket.Conn //gorilla/websocket uses pointer types for connection
}

var _ god.Client = &Client{} //compile time interface check

//New returns a new instance of the Client struct.
func New(protocol god.Protocol, switcher god.Switching, conn *websocket.Conn) *Client {
	return &Client{
		protocol,
		switcher,
		conn,
	}
}

//Listen starts listening for incomming data on the connection. Start as a goroutine!
func (c *Client) Listen() {
	for {
		websocketMessageType, b, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message from connection. ", err)
			//TODO switching remove client
			return
		}

		log.Println("Message type: ", websocketMessageType)
		// TODO protocol decode message
		log.Println("Message content: ", b)

	}
}

//Write writes data to the underlying websocket connection.
func (c *Client) Write(data []byte) error {

	err := c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return ErrWrite
	}
	return nil
}
