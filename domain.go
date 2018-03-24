package godathon2018

const example = 1

type Client interface {
	write([]byte) error
}

type Protocol interface {
	decode([]byte) (error, int, interface{})
}

type Web interface{}

type Auth interface{}
