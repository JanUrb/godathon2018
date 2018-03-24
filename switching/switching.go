package switching

import (
	"fmt"

	"github.com/JanUrb/godathon2018"
)

var callIDCounter = 0

//Group sucks
type Group struct {
	talker  int
	clients map[int]godathon2018.Client
}

//New instance of a group
func NewGroup() Group {
	g := Group{
		clients: make(map[int]godathon2018.Client),
	}
	return g
}

func (g Group) AddClient(clientID int, client godathon2018.Client) {
	fmt.Printf("Group::AddClient clientID %d\n", clientID)
	g.clients[clientID] = client
}

func (g Group) RemoveClient(clientID int) {
	fmt.Printf("Group::RemoveClient clientID %d\n", clientID)
	delete(g.clients, clientID)
}

func (g Group) SetTalkingParty(clientID int) {
	fmt.Printf("Group::SetTalkingParty clientID %d\n", clientID)
	g.talker = clientID
}

func (g Group) GetTalkingParty() int {
	fmt.Printf("Group::SetTalkingParty talker %d\n", g.talker)
	return g.talker
}

func (g Group) GetCalledClients() map[int]godathon2018.Client {
	// create a new map we can copy clients to
	var calledClients = make(map[int]godathon2018.Client)
	// create a copy of the original map
	for clientId, client := range g.clients {
		calledClients[clientId] = client
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

//Switcher sucks
type Switcher struct {
	ongoingCalls map[int]Call
	groups       map[int]Group
	clients      map[int]godathon2018.Client
}

func NewSwitcher() Switcher {
	g := Switcher{
		ongoingCalls: make(map[int]Call),
		groups:       make(map[int]Group),
		clients:      make(map[int]godathon2018.Client),
	}
	return g
}

//Call distributes voice data to all called partys
func (s Switcher) Call(voiceData []byte, groupID int) {
	fmt.Printf("Switching::Call groupID %d\n", groupID)
	var group = s.groups[groupID]
	var calledClients = group.GetCalledClients()
	for _, client := range calledClients {
		client.Write(voiceData)
	}
}

//AttachGroup attaches a client to a group
func (s Switcher) AttachGroup(groupID int, clientID int, client godathon2018.Client) {
	fmt.Printf("Switching::AttachGroup groupID %d clientID %d\n", groupID, clientID)
	if _, ok := s.groups[groupID]; ok {
		// do nothing
	} else {
		s.groups[groupID] = NewGroup()
	}
	s.clients[clientID] = client
	s.groups[groupID].AddClient(clientID, client)
}

//DetachGroup detaches a client from a group
func (s Switcher) DetachGroup(groupID int, clientID int) {
	fmt.Printf("Switching::DetachGroup groupID %d clientID %d\n", groupID, clientID)
	delete(s.clients, clientID)
	s.groups[groupID].RemoveClient(clientID)
}

//DisconnectRequest lol
func (s Switcher) DisconnectRequest(groupID int, clientID int) {
	fmt.Printf("Switching::DisconnectRequest groupID %d clientID %d\n", groupID, clientID)
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
	fmt.Printf("Switching::RequestSetup groupID %d clientID %d\n", groupID, clientID)
	callIDCounter++
	client := s.clients[clientID]
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
