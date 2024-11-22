package odas

import "encoding/json"

type IRequest interface {
	Api() string
	Body() []byte
	Method() string
	ContentType() string
	AuthRequired() bool
}

type Response struct {
	Code   int             `json:"code"`
	Msg    string          `json:"msg"`
	Result json.RawMessage `json:"result,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

func (o *Response) IsOk() bool {
	return o.GetCode() == Ok
}

func (o *Response) GetCode() int {
	return o.Code
}

func (o *Response) GetMsg() string {
	return o.Msg
}

func (o *Response) GetResult() json.RawMessage {
	if o.Result != nil {
		return o.Result
	} else if o.Data != nil {
		return o.Data
	}
	return nil
}

const ProdBaseURL = "https://odas.12301.cc"
const TestBaseURL = "http://10.53.0.14:23080"
const LocalBaseURL = "http://127.0.0.1:80"
const Ok = 0

var baseURL = ProdBaseURL

func SetTestMode() {
	baseURL = TestBaseURL
}

func SetLocalMode() {
	baseURL = LocalBaseURL
}
