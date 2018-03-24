package switching

import (
	"log"

	"github.com/JanUrb/godathon2018"
)

type Switcher struct{}

func (s Switcher) Call(voiceData []byte, group int) error {
	log.Println("Calling")
}

func (s Switcher) AddClient(client godathon2018.Client)
