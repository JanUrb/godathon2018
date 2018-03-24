package godathon2018

const example = 1

type Client interface {
	Write([]byte) error
	Listen()
	OnGroupAttachAck()
	OnSetupAck(int, int)
	OnSetupInd(int, int, int)
	OnSetupFailed()
	OnDisconnectAck()
	OnDisconnectInd()
}

type Protocol interface {
	Decode([]byte) (error, int, interface{})
}

type Web interface {
	Run()
}

type Switching interface {
	Call([]byte, int)
	AttachGroup(int, int, Client) error
	DetachGroup(int, int) error
	RequestSetup(int, int)
	DisconnectRequest(int, int)
}
