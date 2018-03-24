package switching

import (
	"github.com/JanUrb/godathon2018"
)

var callIDCounter = 0

//Group sucks
type Group struct {
	talker  int
	clients map[int]godathon2018.Client
}

//New instance of a group
func New() Group {
	g := Group{
		clients: make(map[int]godathon2018.Client),
	}
	return g
}

func (g Group) AddClient(clientId int, client godathon2018.Client) {
	g.clients[clientId] = client
}

func (g Group) RemoveClient(clientId int) {
	delete(g.clients, clientId)
}

func (g Group) SetTalkingParty(clientId int) {
	g.talker = clientId
}

func (g Group) GetTalkingParty() int {
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

//Call distributes voice data to all called partys
func (s Switcher) Call(voiceData []byte, groupID int) {
	var group = s.groups[groupID]
	var calledClients = group.GetCalledClients()
	for clientID, client := range calledClients {
		client.Write(voiceData)
	}
}

//AttachGroup attaches a client to a group
func (s Switcher) AttachGroup(groupID int, clientID int, client godathon2018.Client) {
	s.groups[groupID].AddClient(clientID, client)
}

//DetachGroup detaches a client from a group
func (s Switcher) DetachGroup(groupID int, clientID int) {
	s.groups[groupID].RemoveClient(clientID)
}

//RequestTxDemand sets the current talking party of a group. Throws error when there is already a talking party active
/*func (s Switcher) RequestTxDemand(groupID int, clientID int) {
	var group = s.groups[groupID]
	// there is already a caller, inform requesting client
	if group.GetTalkingParty() != 0 {
		// TODO throw error somehow
	}
	group.SetTalkingParty(clientID)
	// inform calling party about TxCeased
	client := s.clients[group.GetTalkingParty()]
	client.OnTxCeasedAck()
	// inform called partys about TxCeased
	var calledClients = group.GetCalledClients()
	for clientID, client := range calledClients {
		client.OnTxInfoInd()
	}
}*/

func (s Switcher) DisconnectRequest(groupID int, clientID int) {
	var group = s.groups[groupID]
	client := s.clients[group.GetTalkingParty()]
	client.OnDisconnectAck()
	var calledClients = group.GetCalledClients()
	for clientID, client := range calledClients {
		client.OnDisconnectInd()
	}
	// remove call from map
	delete(s.ongoingCalls, groupID)
}

//RequestSetup bla
func (s Switcher) RequestSetup(groupID int, clientID int) {
	callIDCounter++
	client := s.clients[clientID]
	if val, ok := s.ongoingCalls[groupID]; ok {
		client.OnSetupFailed()
	} else {
		call := NewCall(callIDCounter, groupID)
		s.ongoingCalls[call.groupID] = call
		var group = s.groups[groupID]
		// inform calling party about call setup
		client.OnSetupAck(200, call.callID)
		// inform called partys about call setup
		var calledClients = group.GetCalledClients()
		for clientID, client := range calledClients {
			client.OnSetupInd(groupID, clientID, call.callID)
		}
	}
}
