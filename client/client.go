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
			err = c.switcher.DetachGroup(c.groupID, c.clientID)
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
			err = c.Write([]byte("error hallo"))
			if err != nil {
				log.Println("Error while writing error hello")
			}
			log.Printf("Error while reading generic message. (Client Id: %d, message: %s)", c.clientID, genericMsg)
			continue // continue reading messages. No need to kill the client
		}

		log.Println("Message content: ", b)
		c.resolveMessageID(genericMsg.Msg_type, genericMsg.Payload)
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

func (c *Client) resolveMessageID(messageType string, payload []byte) {
	switch messageType {
	case protocol.MessageType_groupAttach_req:
		{
			req, err := protocol.DecodeGroupAttachReq(payload)
			if err != nil {
				log.Println("Could not decode group attach req ", payload)
				return
			}
			err = c.switcher.AttachGroup(req.Id, 0, c)
			if err != nil {
				log.Println("Error while attaching to group ")
				return
			}
		}
	case protocol.MessageType_setup_req:
		{
			req, err := protocol.DecodeSetupReq(payload)
			if err != nil {
				log.Println("Error decide setup req")
				return
			}
			log.Println("Decide setup request calltype: ", req.Call_type, " calledId ", req.Called_id)
			c.switcher.RequestSetup(0, 0)
		}
	default:
		log.Println("Received unknown message")
	}
}

//OnSetupAck sends to the underlying connection
func (c *Client) OnSetupAck(result int, callID int) {
	setupAck := protocol.Setup_ack{
		Result:  result,
		Call_id: callID,
	}
	b, err := protocol.EncodeSetupAck(setupAck)
	if err != nil {
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while writing SetUpAckRequest ", err)
		return
	}
}

//OnSetupInd sends to the underlying connection
func (c *Client) OnSetupInd(groupID, callID, clientID int) {
	setupInd := protocol.Setup_ind{
		Calling_id: groupID,
		Call_id:    callID,
		Called_id:  clientID,
	}
	b, err := protocol.EncodeSetupInd(setupInd)
	if err != nil {
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while writing setupInd ", err)
		return
	}
}

//OnSetupFailed sends a SetupAck with error code
func (c *Client) OnSetupFailed() {
	//send with result 500
	setupAck := protocol.Setup_ack{Result: 417, Call_id: 417}
	b, err := protocol.EncodeSetupAck(setupAck)
	if err != nil {
		log.Println("Error while sending onsetupfailed ", err)
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while sending onsetupFailed ", err)
		return
	}
}

//OnGroupAttachAck sends an groupattachack
func (c *Client) OnGroupAttachAck() {
	setupAck := protocol.Setup_ack{}
	b, err := protocol.EncodeSetupAck(setupAck)
	if err != nil {
		log.Println("Error encode setupack ", err)
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while writing groupAttachAck ", err)
		return
	}
}

//OnDisconnectAck sends a disconnectAck
func (c *Client) OnDisconnectAck() {
	disconnectAck := protocol.Disconnect_ack{}
	b, err := protocol.EncodeDisconnectAck(disconnectAck)
	if err != nil {
		log.Println("Error while writing disconnectAck ", err)
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while writing disconnectAck ", err)
		return
	}
}

//OnDisconnectInd sends a disconnectInd
func (c *Client) OnDisconnectInd() {
	disconnectInd := protocol.Disconnect_ind{}
	b, err := protocol.EncodeDisconnectInd(disconnectInd)
	if err != nil {
		log.Println("Error encode disconnect ind ", err)
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Error while writing disconnectInd ", err)
		return
	}
}
