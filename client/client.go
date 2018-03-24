package client

import (
	"encoding/json"
	"errors"
	"log"

	god "github.com/JanUrb/godathon2018"
	"github.com/JanUrb/godathon2018/protocol"
	"github.com/gorilla/websocket"
)

var (
	//ErrWrite indicates that an error while writing happened.
	ErrWrite = errors.New("ErrWrite")
)

//Client represents a websocket client.
type Client struct {
	clientID int
	groupID  int
	switcher god.Switching
	conn     *websocket.Conn //gorilla/websocket uses pointer types for connection
}

var _ god.Client = (*Client)(nil) //compile time interface check

//New returns a new instance of the Client struct.
func New(switcher god.Switching, conn *websocket.Conn) *Client {
	return &Client{
		switcher: switcher,
		conn:     conn,
	}
}

//Listen starts listening for incomming data on the connection. Start as a goroutine!
func (c *Client) Listen() {
	for {
		websocketMessageType, b, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message from connection. ", err)
			err := c.switcher.DetachGroup(c.groupID, c.clientID)
			if err != nil {
				log.Println("Error while detaching from group: ", c.groupID, " with client: ", c.clientID)
				return
			}
			return
		}

		log.Println("Message type: ", websocketMessageType)
		//TODO protocol decode message

		var genericMsg protocol.Generic_message
		err = json.Unmarshal(b, &genericMsg)
		if err != nil {
			log.Printf("Error while reading generic message. (Client Id: %d, message: %s)", c.clientID, genericMsg)
			continue // continue reading messages. No need to kill the client
		}



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



func (c *Client) resolveMessage(messageType int){
	switch(messageType){
	case: 
	}
}



//OnTxCeasedAck sends to the underlying connection
func (c *Client) OnTxCeasedAck() {

	panic("not implemented")
}

//OnTxInfoInd sends to the underlying connection
func (c *Client) OnTxInfoInd() {
	panic("not implemented")
}

//OnTxDemandAck sends to the underlying connection
func (c *Client) OnTxDemandAck() {
	panic("not implemented")
}

//OnSetupAck sends to the underlying connection
func (c *Client) OnSetupAck() {
	panic("not implemented")
}

//OnSetupInd sends to the underlying connection
func (c *Client) OnSetupInd() {
	panic("not implemented")
}

//OnConnectAck sends to the underlying connection
func (c *Client) OnConnectAck() {
	panic("not implemented")
}
