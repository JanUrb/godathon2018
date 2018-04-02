package switching

import (
	"testing"

	"github.com/JanUrb/godathon2018"
)

type fakeClient struct {
	onGroupAttachAck struct {
		called bool
		group  int
	}
}

func (fc fakeClient) Write([]byte) error {
	return nil
}

func (fc fakeClient) Listen() {}

func (fc fakeClient) OnGroupAttachAck(group int) {
	fc.onGroupAttachAck.called = true
	fc.onGroupAttachAck.group = group
}

func (fc fakeClient) OnSetupAck(int, int) {}

func (fc fakeClient) OnSetupInd(int, int) {}

func (fc fakeClient) OnSetupFailed() {}

func (fc fakeClient) OnDisconnectAck() {}

func (fc fakeClient) OnDisconnectInd() {}

func generateFakeClient() *fakeClient {
	return &fakeClient{}
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
		_ = switcher.DetachGroup(tt.groupID, tt.clientID)

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

func TestRequestSetup(t *testing.T) {}
