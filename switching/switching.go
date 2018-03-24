package switching

import (
	"github.com/JanUrb/godathon2018"
)

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

//Switcher sucks
type Switcher struct {
	groups  map[int]Group
	clients map[int]godathon2018.Client
}

//Call distributes voice data to all called partys
func (s Switcher) Call(voiceData []byte, groupID int) {
	var group = s.groups[groupID]
	var calledClients = group.GetCalledClients()
	for clientID, client := range calledClients {
		client.Write(voiceData)
	}
}

//AttachClientToGroup attaches a client to a group
func (s Switcher) AttachClientToGroup(groupID int, clientID int, client godathon2018.Client) error {
	s.groups[groupID].AddClient(clientID, client)
}

//DetachClientFromGroup detaches a client from a group
func (s Switcher) DetachClientFromGroup(groupID int, clientID int) error {
	s.groups[groupID].RemoveClient(clientID)
}

//SetCaller sets the current talking party of a group. Throws error when there is already a talking party active
func (s Switcher) SetCaller(groupID int, clientID int) error {
	var group = s.groups[groupID]
	// there is already a caller, throw error
	if group.GetTalkingParty() != 0 {
		// TODO throw error somehow
	}
}

//RemoveCaller txCeased of current talking party
func (s Switcher) RemoveCaller(groupID int, clientID int) error {
	var group = s.groups[groupID]
	// there is no caller, throw error
	if group.GetTalkingParty() == 0 {
		// TODO throw error somehow
	}
	// inform all clients that the talking party ceased tx
}
