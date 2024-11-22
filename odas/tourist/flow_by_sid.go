package tourist

import (
	"fmt"
	"net/http"
	"net/url"
)

// FlowBySidReq 根据sid查询出入园数据
type FlowBySidReq struct {
	Sid string `json:"sid"`
}

func (f FlowBySidReq) Api() string {
	params := url.Values{}
	if f.Sid != "" {
		params.Add("sid", f.Sid)
	}
	return fmt.Sprintf("/v2/tourist/inout/flowBySid?%s", params.Encode())
}

func (f FlowBySidReq) Body() []byte {
	return nil
}

func (f FlowBySidReq) Method() string {
	return http.MethodGet
}

func (f FlowBySidReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (f FlowBySidReq) AuthRequired() bool {
	return true
}

func NewFlowBySidReq(sid string) *FlowBySidReq {
	return &FlowBySidReq{
		Sid: sid,
	}
}

type FlowBySidResponse struct {
	Total InoutTotal   `json:"total"`
	List  []*InoutList `json:"list"`
}
