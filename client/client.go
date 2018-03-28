package client

import (
	"encoding/json"
	"errors"

	god "github.com/JanUrb/godathon2018"
	"github.com/JanUrb/godathon2018/protocol"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	//ErrWrite indicates that an error while writing happened.
	ErrWrite = errors.New("ErrWrite")
)

//Client represents a websocket client.
type Client struct {
	clientID      int
	groupID       int
	userName      string
	switcher      god.Switching
	conn          *websocket.Conn
	log           *logrus.Entry
	sendLogger    *logrus.Entry
	receiveLogger *logrus.Entry
}

var _ god.Client = (*Client)(nil) //compile time interface check

var simpleIDGen = 1

//New returns a new instance of the Client struct.
func New(switcher god.Switching, conn *websocket.Conn) *Client {
	log := logrus.New().WithFields(logrus.Fields{
		"component":  "client",
		"clientIP: ": conn.RemoteAddr().String(),
	})
	sendLogger := log.WithField("event", "send")
	receiveLogger := log.WithField("event", "receive")

	simpleIDGen = simpleIDGen + 1
	clientID := simpleIDGen
	log.Infoln("UserID: ", clientID)

	return &Client{
		switcher:      switcher,
		conn:          conn,
		clientID:      clientID,
		log:           log,
		sendLogger:    sendLogger,
		receiveLogger: receiveLogger,
	}
}

//Listen starts listening for incomming data on the connection. Start as a goroutine!
func (c *Client) Listen() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			c.log.Warnln("Error while closing websocket")
		}
	}()
	for {
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			c.log.Warnln("Error while reading message from connection. ", err)
			// err = c.switcher.DetachGroup(c.groupID, c.clientID)
			// if err != nil {
			// 	log.Println("Error while detaching from group: ", c.groupID, " with client: ", c.clientID)
			// 	return
			// }
			return
		}

		var genericMsg protocol.Generic_message
		err = json.Unmarshal(b, &genericMsg)
		if err != nil {
			c.log.Warnln("Error reading generic message: ", err)
			c.log.Errorf("Error while reading generic message. (Client Id: %d, message: %s )", c.clientID, genericMsg)
			continue // continue reading messages. No need to kill the client
		}

		c.receiveLogger.Println("Handling Msg Type:", genericMsg.Msg_type)
		c.handleMessage(genericMsg.Msg_type, genericMsg.Payload)
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

func (c *Client) handleMessage(messageType string, payload []byte) {
	switch messageType {
	case protocol.MessageType_register_req:
		{

			req, err := protocol.DecodeRegisterReq(payload)
			if err != nil {
				c.receiveLogger.Warnln("Could not decode registerRequest: ", err)
				return
			}
			c.userName = req.User
			c.SendRegisterAck()

		}
	case protocol.MessageType_groupAttach_req:
		{
			req, err := protocol.DecodeGroupAttachReq(payload)
			if err != nil {
				c.receiveLogger.Warnln("Could not decode group attach req", err)
				return
			}
			c.receiveLogger.Info("Attaching to group: ", req.GroupID)
			err = c.switcher.AttachGroup(req.GroupID, c.clientID, c)
			if err != nil {
				c.receiveLogger.Warnln("Error while attaching to group ", err)
				return
			}
		}
	case protocol.MessageType_setup_req:
		{
			req, err := protocol.DecodeSetupReq(payload)
			if err != nil {
				c.log.Warnln("Error decode setup req", err)
				return
			}
			c.receiveLogger.Println("Decode setup request groupID: ", req.GroupID)
			c.switcher.RequestSetup(req.GroupID, c.clientID)
		}
	case protocol.MessageType_disconnect_req:
		{
			req, err := protocol.DecodeDisconnectReq(payload)
			if err != nil {
				c.log.Warnln("Error decoding disconnect request", err)
				return
			}
			c.receiveLogger.Println("Disconnecting call: ", req.GroupID)
			c.switcher.DisconnectRequest(c.clientID, c.groupID)
		}
	default:
		c.receiveLogger.Println("Received unknown message:", messageType)
	}
}

//SendRegisterAck sends a register ack.
func (c *Client) SendRegisterAck() {
	var registerAck protocol.Register_ack
	registerAck.Result = 200
	b, err := protocol.EncodeRegisterAck(registerAck)
	if err != nil {
		c.sendLogger.Warnln("Failed to encode register ack")
		return
	}
	c.sendLogger.Infoln("Sending registerAck")
	err = c.Write(b)
	if err != nil {
		c.sendLogger.Warnln("Failed to write to user")
		return
	}
}

//OnSetupAck sends to the underlying connection
func (c *Client) OnSetupAck(result, groupID int) {
	setupAck := protocol.Setup_ack{
		Result:  result,
		GroupID: groupID,
	}
	b, err := protocol.EncodeSetupAck(setupAck)
	if err != nil {
		c.sendLogger.Warnln("Failed to encode setupAck")
		return
	}
	c.sendLogger.Infoln("Sending setupAck")
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while writing SetUpAckRequest ", err)
		return
	}
}

//OnSetupInd sends to the underlying connection
func (c *Client) OnSetupInd(groupID, clientID int) {
	setupInd := protocol.Setup_ind{
		CalleeID: clientID,
		GroupID:  groupID,
	}
	b, err := protocol.EncodeSetupInd(setupInd)
	if err != nil {
		c.sendLogger.Warn("Failed to encode setupInd")
		return
	}
	c.sendLogger.Infoln("Sending onSetupInd")
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while writing setupInd ", err)
		return
	}
}

//OnSetupFailed sends a SetupAck with error code
func (c *Client) OnSetupFailed() {
	//send with result 500
	setupAck := protocol.Setup_ack{Result: 417, GroupID: 417}
	b, err := protocol.EncodeSetupAck(setupAck)
	if err != nil {
		c.sendLogger.Warnln("Error while sending onsetupfailed ", err)
		return
	}
	c.sendLogger.Infoln("Sending onSetupAck Failed")
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while sending onsetupFailed ", err)
		return
	}
}

//OnGroupAttachAck sends an groupattachack
func (c *Client) OnGroupAttachAck(groupID int) {
	groupAck := protocol.Group_attach_ack{GroupID: groupID, Result: 200}
	b, err := protocol.EncodeGroupAttachAck(groupAck)
	if err != nil {
		c.sendLogger.Warnln("Error encode setupack ", err)
		return
	}
	c.sendLogger.Infoln("Sending groupAttachAck with groupID", groupID)
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while writing groupAttachAck ", err)
		return
	}
}

//OnDisconnectAck sends a disconnectAck
func (c *Client) OnDisconnectAck() {
	disconnectAck := protocol.Disconnect_ack{}
	b, err := protocol.EncodeDisconnectAck(disconnectAck)
	if err != nil {
		c.sendLogger.Warnln("Error while writing disconnectAck ", err)
		return
	}
	c.sendLogger.Infoln("Sending disconnectAck")
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while writing disconnectAck ", err)
		return
	}
}

//OnDisconnectInd sends a disconnectInd
func (c *Client) OnDisconnectInd() {
	disconnectInd := protocol.Disconnect_ind{}
	b, err := protocol.EncodeDisconnectInd(disconnectInd)
	if err != nil {
		c.sendLogger.Warnln("Error encode disconnect ind ", err)
		return
	}
	c.sendLogger.Infoln("Sending disconnectInd")
	err = c.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		c.sendLogger.Warnln("Error while writing disconnectInd ", err)
		return
	}
}
