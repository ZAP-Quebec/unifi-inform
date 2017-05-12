package data

type Message interface {
	Marshal() []byte
	String() string
}
