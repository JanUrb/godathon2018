package switching

import (
	"testing"

	"github.com/JanUrb/godathon2018"
)

type fakeClient struct {
	//test struct used to verify that onGroupAttach was called and the right group was attached
	onGroupAttachAck struct {
		called bool
		group  int
	}

	onSetupInd struct {
		indicationSend bool
		group          int
	}
	onSetupFailedCalled bool
	onSetupAck          struct {
		ackSend    bool
		group      int
		resultCode int
	}
}

func (fc *fakeClient) Write([]byte) error {
	return nil
}

func (fc *fakeClient) Listen() {}

func (fc *fakeClient) OnGroupAttachAck(group int) {
	fc.onGroupAttachAck.called = true
	fc.onGroupAttachAck.group = group
}

func (fc *fakeClient) OnSetupAck(resultCode, groupID int) {
	fc.onSetupAck.ackSend = true
	fc.onSetupAck.group = groupID
	fc.onSetupAck.resultCode = resultCode
}

func (fc *fakeClient) OnSetupInd(groupID, clientID int) {
	fc.onSetupInd.indicationSend = true
	fc.onSetupInd.group = groupID
}

func (fc *fakeClient) OnSetupFailed() {
	fc.onSetupFailedCalled = true
}

func (fc *fakeClient) OnDisconnectAck() {}

func (fc *fakeClient) OnDisconnectInd() {}

func generateFakeClient() *fakeClient {
	fc := &fakeClient{
		onSetupFailedCalled: false,
	}
	//init embedded struct
	// fc.onSetupAck.ackSend = false
	return fc
}

func TestAttachGroup(t *testing.T) {
	switcher := NewSwitcher()

	var testData = []struct {
		groupID  int
		clientID int
		client   godathon2018.Client
	}{
		{1, 2, generateFakeClient()},
		{3, 5, generateFakeClient()},
		{2, 1000, generateFakeClient()},
		{3, 5, generateFakeClient()},
		{123, 22200, generateFakeClient()},
		{1, 0, generateFakeClient()},
		{1, 0, generateFakeClient()},
	}
	for _, tt := range testData {
		_ = switcher.AttachGroup(tt.groupID, tt.clientID, tt.client)

		group, ok := switcher.groups[tt.groupID]
		if !ok {
			t.Errorf("Group with id %d not found", tt.groupID)
		}

		client, ok := group.clients[tt.clientID]
		if !ok {
			t.Errorf("Client with id %d not found", tt.clientID)
		}

		if client != tt.client {
			t.Error("The clients do not match")
		}
	}
}

func TestDetachGroup(t *testing.T) {
	switcher := NewSwitcher()

	var testData = []struct {
		groupID  int
		clientID int
		client   godathon2018.Client
	}{
		{1, 0, generateFakeClient()},
		{1, 1, generateFakeClient()},
		{1, 2, generateFakeClient()},
		{1, 3, generateFakeClient()},
		{1, 4, generateFakeClient()},
	}

	for _, tt := range testData {
		_ = switcher.AttachGroup(tt.groupID, tt.clientID, tt.client)
	}

	for _, tt := range testData {
		err := switcher.DetachGroup(tt.groupID, tt.clientID)

		if err != nil {
			t.Errorf("Detaching %d from %d failed with error %s", tt.clientID, tt.groupID, err)
		}

		g, ok := switcher.groups[tt.groupID]
		if !ok {
			t.Errorf("Group with id %d not found", tt.groupID)
		}

		_, ok = g.clients[tt.clientID]
		if ok {
			t.Errorf("Client with id %d found", tt.clientID)
		}
	}
	if len(switcher.groups[1].clients) != 0 {
		t.Error("The gorup should not hold any clients anymore")
	}
}

func TestRequestSetupSingleGroupMember(t *testing.T) {
	switcher := NewSwitcher()

	groupID := 1
	clientID := 1
	c := generateFakeClient()

	err := switcher.AttachGroup(groupID, clientID, c)
	if err != nil {
		t.Error("Failed to attach group")
	}

	switcher.RequestSetup(groupID, clientID)

	g, ok := switcher.groups[groupID]
	if !ok {
		t.Error("Group should have been created")
	}

	talker := g.getTalkingParty()
	if talker != clientID {
		t.Error("No other member is in the group. Therefore, the only member should be the talker")
	}
}

func TestRequestSetupInCall(t *testing.T) {
	switcher := NewSwitcher()

	var testData = []struct {
		groupID  int
		clientID int
		client   godathon2018.Client
	}{
		{1, 0, generateFakeClient()},
		{1, 1, generateFakeClient()},
	}

	for _, tt := range testData {
		err := switcher.AttachGroup(tt.groupID, tt.clientID, tt.client)

		if err != nil {
			t.Errorf("Failed to attach client: %d group", tt.clientID)
		}
	}

	t1 := testData[0]
	t2 := testData[1]

	switcher.RequestSetup(t1.groupID, t1.clientID)

	switcher.RequestSetup(t2.groupID, t2.clientID)

	c1 := t1.client.(*fakeClient)
	c2 := t2.client.(*fakeClient)

	if !c1.onSetupAck.ackSend {
		t.Error("SetupAck should have been send")
	}

	if c1.onSetupAck.group != t1.groupID {
		t.Error("SetupAck was given for wrong group")
	}

	if !c2.onSetupFailedCalled {
		t.Error("The second client should get setupFailed")
	}

	if !c2.onSetupInd.indicationSend {
		t.Error("The second client should have been given a setupInd")
	}

	if c2.onSetupInd.group != t2.groupID {
		t.Error("The setupInd was about the wrong group")
	}

	g, ok := switcher.groups[t1.groupID]
	if !ok {
		t.Error("Group should have been created")
	}

	talker := g.getTalkingParty()
	if talker != t1.clientID {
		t.Error("Client1 should have gotten talking first")
	}

}
