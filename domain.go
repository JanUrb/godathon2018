package godathon2018

//Client is the layer that directly talks to the client of the server
type Client interface {
	Write([]byte) error
	Listen()
	OnGroupAttachAck(int)
	OnSetupAck(int, int)
	OnSetupInd(int, int)
	OnSetupFailed()
	OnDisconnectAck()
	OnDisconnectInd()
}

//Web accepts socket connections and hosts the client app
type Web interface {
	Run()
}

//Switching contains the logic for connecting clients in groups and manages calls.
type Switching interface {
	Call([]byte, int)
	AttachGroup(int, int, Client) error
	DetachGroup(int, int) error
	RequestSetup(int, int)
	DisconnectRequest(int, int)
}
