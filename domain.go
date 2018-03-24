package godathon2018

const example = 1

type Client interface {
	Write([]byte) error
	Listen()
}

type Protocol interface {
	Decode([]byte) (error, int, interface{})
}

type Web interface{}

type Switching interface {
	Call([]byte) error
}
