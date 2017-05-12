package data

import (
	"fmt"
)

type InformResponse interface {
	Message
	IsSuccess() bool
	HttpCode() int
}

type httpResponse struct {
	code int
}

func ResponseFromHttpCode(code int) InformResponse {
	return httpResponse{code}
}

func (r httpResponse) IsSuccess() bool {
	return r.code == 200
}

func (r httpResponse) HttpCode() int {
	return r.code
}

func (r httpResponse) Marshal() []byte {
	return []byte(r.String())
}

func (r httpResponse) String() string {
	return fmt.Sprintf(`{"code":%d}`, r.code)
}
