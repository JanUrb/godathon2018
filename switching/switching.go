package switching

import (
	"github.com/sirupsen/logrus"

	"github.com/JanUrb/godathon2018"
)

var callIDCounter = 0
var log = logrus.WithField("component", "switching")

//Group sucks
type Group struct {
	talker      int
	clients     map[int]godathon2018.Client
	groupLogger *logrus.Entry
}

//NewGroup returns an instance of a group
func NewGroup() *Group {
	g := &Group{
		clients:     make(map[int]godathon2018.Client),
		groupLogger: log.WithField("subcomponent", "group"),
	}
	return g
}

//AddClient adds a client.
func (g *Group) AddClient(clientID int, client godathon2018.Client) {
	g.groupLogger.Infoln("AddClient clientID", clientID)
	g.clients[clientID] = client
}

//RemoveClient removes a client
func (g *Group) RemoveClient(clientID int) {
	g.groupLogger.Infoln("RemoveClient clientID", clientID)
	delete(g.clients, clientID)
}

//SetTalkingParty sets the talking party of a group
func (g *Group) SetTalkingParty(clientID int) {
	g.groupLogger.Infoln("SetTalkingParty clientID", clientID)
	g.talker = clientID
}

//GetTalkingParty returns the current talking party
func (g *Group) GetTalkingParty() int {
	g.groupLogger.Infoln("SetTalkingParty talker", g.talker)
	return g.talker
}

//GetCalledClients returns the clients that are currently called
func (g *Group) GetCalledClients() map[int]godathon2018.Client {
	// create a new map we can copy clients to
	var calledClients = make(map[int]godathon2018.Client)
	// create a copy of the original map
	for clientID, client := range g.clients {
		calledClients[clientID] = client
	}
	// remove the callee from list
	delete(calledClients, g.GetTalkingParty())
	// return the new map
	return calledClients
}

type Call struct {
	callID  int
	groupID int
}

//NewCall instance of a call
func NewCall(callID int, groupID int) Call {
	c := Call{
		callID:  callID,
		groupID: groupID,
	}
	return c
}

var _ godathon2018.Switching = Switcher{} //compile time interface check

//Switcher sucks
type Switcher struct {
	ongoingCalls map[int]Call
	groups       map[int]*Group
	clients      map[int]godathon2018.Client
}

//NewSwitcher returns an instance of a switcher
func NewSwitcher() Switcher {
	g := Switcher{
		ongoingCalls: make(map[int]Call),
		groups:       make(map[int]*Group),
		clients:      make(map[int]godathon2018.Client),
	}
	return g
}

//Call distributes voice data to all called partys
func (s Switcher) Call(voiceData []byte, groupID int) {
	log.Infof("Call groupID %d", groupID)
	var group = s.groups[groupID]
	var calledClients = group.GetCalledClients()
	for _, client := range calledClients {
		client.Write(voiceData)
	}
}

//AttachGroup attaches a client to a group
func (s Switcher) AttachGroup(groupID int, clientID int, client godathon2018.Client) error {
	log.Infof("AttachGroup groupID %d clientID %d", groupID, clientID)
	if _, ok := s.groups[groupID]; ok {
		// do nothing
	} else {
		s.groups[groupID] = NewGroup()
	}
	s.clients[clientID] = client
	s.groups[groupID].AddClient(clientID, client)
	client.OnGroupAttachAck(groupID)
	return nil
}

//DetachGroup detaches a client from a group
func (s Switcher) DetachGroup(groupID int, clientID int) error {
	log.Infof("DetachGroup groupID %d clientID %d", groupID, clientID)
	delete(s.clients, clientID)
	s.groups[groupID].RemoveClient(clientID)
	return nil
}

//DisconnectRequest lol
func (s Switcher) DisconnectRequest(groupID int, clientID int) {
	log.Infof("DisconnectRequest groupID %d clientID %d", groupID, clientID)
	// TODO check if ongoing call exists
	var group = s.groups[groupID]
	client := s.clients[group.GetTalkingParty()]
	group.SetTalkingParty(-1)
	client.OnDisconnectAck()
	var calledClients = group.GetCalledClients()
	for _, client := range calledClients {
		client.OnDisconnectInd()
	}
	// remove call from map
	delete(s.ongoingCalls, groupID)
}

//RequestSetup bla
func (s Switcher) RequestSetup(groupID int, clientID int) {
	log.Infof("RequestSetup groupID %d clientID %d", groupID, clientID)
	callIDCounter++
	client, ok := s.clients[clientID]
	if !ok {
		log.Warnln("No client with id ", clientID, " found in currently saved clients")
		return
	}
	if _, ok := s.ongoingCalls[groupID]; ok {
		client.OnSetupFailed()
	} else {
		call := NewCall(callIDCounter, groupID)
		s.ongoingCalls[call.groupID] = call
		var group = s.groups[groupID]
		group.SetTalkingParty(clientID)
		// inform calling party about call setup
		client.OnSetupAck(200, call.callID)
		// inform called partys about call setup
		var calledClients = group.GetCalledClients()
		for clientID, client := range calledClients {
			client.OnSetupInd(groupID, clientID, call.callID)
		}
	}
}
