package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JanUrb/godathon2018/switching"
	"github.com/JanUrb/godathon2018/web"
)

func main() {

	interuptChan := make(chan os.Signal, 1)

	signal.Notify(interuptChan, syscall.SIGINT, syscall.SIGTERM)

	switcher := switching.New()

	webServer := web.Web{
		Switcher: switcher,
	}

	go webServer.Run()

	s := <-interuptChan
	fmt.Println("Got signal: ", s)
	switcher.DisconnectAll()
}
