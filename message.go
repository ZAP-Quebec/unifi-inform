package inform

type Message interface {
	Marshal() []byte
	String() string
}

type InformRequest struct {
	//
}

type InformResponse interface {
	Message
}
