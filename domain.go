package godathon2018

const example = 1

type Client interface {
	Write([]byte) error
	Listen()
	OnTxCeasedAck()
	OnTxInfoInd()
	OnTxDemandAck()
	OnSetupAck()
	OnSetupInd()
	OnConnectAck()
}

type Protocol interface {
	Decode([]byte) (error, int, interface{})
}

type Web interface{}

type Switching interface {
	Call([]byte, int)
	AttachGroup(int, int, Client) error
	DetachGroup(int, int) error
	RequestTxCeased()
	RequestTxDemand()
	RequestSetup()
	RequestConnect()
}
