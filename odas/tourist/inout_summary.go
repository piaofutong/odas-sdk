package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SummaryReq struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	Sid     int    `json:"sid"`
	Gid     int    `json:"gid"`
	NoAmend bool   `json:"noAmend"`
}

func (s SummaryReq) Api() string {
	params := url.Values{}
	if s.Start != "" {
		params.Add("start", s.Start)
	}
	if s.End != "" {
		params.Add("end", s.End)
	}
	if s.Sid > 0 {
		params.Add("sid", strconv.Itoa(s.Sid))
	}
	if s.Gid > 0 {
		params.Add("gid", strconv.Itoa(s.Gid))
	}
	if s.NoAmend {
		params.Add("noAmend", strconv.FormatBool(s.NoAmend))
	}
	return fmt.Sprintf("/v4/tourist/inout/summary?%s", params.Encode())
}

func (s SummaryReq) Body() []byte {
	return nil
}

func (s SummaryReq) Method() string {
	return http.MethodGet
}

func (s SummaryReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (s SummaryReq) AuthRequired() bool {
	return true
}

func NewSummaryReq(start, end string, sid, gid int, noAmend bool) *SummaryReq {
	return &SummaryReq{
		Start:   start,
		End:     end,
		Sid:     sid,
		Gid:     gid,
		NoAmend: noAmend,
	}
}

type InoutSummaryResponse struct {
	Total InoutTotal   `json:"total"`
	List  []*InoutList `json:"list"`
}
