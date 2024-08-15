package odas

import "encoding/json"

type IRequest interface {
	Api() string
	Body() []byte
	Method() string
	ContentType() string
}

type ODASResponse struct {
	Code   int             `json:"code"`
	Msg    string          `json:"msg"`
	Result json.RawMessage `json:"result,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

const BaseURL = "https://odas.12301.cc"
const Ok = 0
