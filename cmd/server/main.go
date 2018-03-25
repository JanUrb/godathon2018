package main

import "github.com/JanUrb/godathon2018/web"
import "github.com/JanUrb/godathon2018/switching"

func main() {

	switcher := switching.NewSwitcher()

	webServer := web.Web{
		Switcher: switcher,
	}

	webServer.Run()
}
