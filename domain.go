package godathon2018

const example = 1

type client interface {
	write([]byte) error
}

type protocol interface {
	decode([]byte) (error, int, interface{})
}
