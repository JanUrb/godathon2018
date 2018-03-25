package switching

import (
	"testing"

	"github.com/JanUrb/godathon2018"
)

func TestAttachGroup(t *testing.T) {
	switcher := NewSwitcher()

	var testData = []struct {
		groupID  int
		clientID int
		client   godathon2018.Client
	}{
		{1, 2, nil},
		{3, 5, nil},
		{2, 1000, nil},
		{3, 5, nil},
		{123, 22200, nil},
		{1, 0, nil},
		{1, 0, nil},
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
