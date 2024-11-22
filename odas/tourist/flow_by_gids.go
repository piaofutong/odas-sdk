package tourist

import (
	"fmt"
	"net/http"
	"net/url"
)

// FlowByGIdsReq 根据gids查询出入园数据
type FlowByGIdsReq struct {
	GIds string `json:"gIds"`
	Date string `json:"date"`
}

func (f FlowByGIdsReq) Api() string {
	params := url.Values{}
	if f.Date != "" {
		params.Add("date", f.Date)
	}
	ps := fmt.Sprintf("/v2/tourist/inout/flowByGIds?%s", params.Encode())
	if f.GIds != "" {
		ps += fmt.Sprintf("&gIds=%s", f.GIds)
	}
	return ps
}

func (f FlowByGIdsReq) Body() []byte {
	return nil
}

func (f FlowByGIdsReq) Method() string {
	return http.MethodGet
}

func (f FlowByGIdsReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (f FlowByGIdsReq) AuthRequired() bool {
	return true
}

func NewFlowByGIdsReq(gIds, date string) *FlowByGIdsReq {
	return &FlowByGIdsReq{
		GIds: gIds,
		Date: date,
	}
}

type FlowByGIdsResponse struct {
	Total InoutTotal   `json:"total"`
	List  []*InoutList `json:"list"`
}

type InoutTotal struct {
	In   int `json:"in"`
	Out  int `json:"out"`
	Hold int `json:"hold"`
}

type InoutList struct {
	Time string `json:"time"`
	In   int    `json:"in"`
	Out  int    `json:"out"`
}
