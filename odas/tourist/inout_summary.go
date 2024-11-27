package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SummaryReq struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Sid      int    `json:"sid"`
	Gid      int    `json:"gid"`
	NoAmend  bool   `json:"noAmend"`
	DateType int    `json:"dateType"`
}

type SummaryReqOptions func(options *SummaryReq)

func WithDateType(dateType int) SummaryReqOptions {
	return func(options *SummaryReq) {
		options.DateType = dateType
	}
}

func WithNoAmend() SummaryReqOptions {
	return func(options *SummaryReq) {
		options.NoAmend = true
	}
}

func WithStart(start string) SummaryReqOptions {
	return func(options *SummaryReq) {
		options.Start = start
	}
}

func WithEnd(end string) SummaryReqOptions {
	return func(options *SummaryReq) {
		options.End = end
	}
}

func WithSid(sid int) SummaryReqOptions {
	return func(options *SummaryReq) {
		options.Sid = sid
	}
}

func WithGid(gid int) SummaryReqOptions {
	return func(options *SummaryReq) {
		options.Gid = gid
	}
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
	if s.DateType > 0 {
		params.Add("dateType", strconv.Itoa(s.DateType))
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

func NewSummaryReq(opt ...SummaryReqOptions) *SummaryReq {
	req := &SummaryReq{}
	for _, options := range opt {
		options(req)
	}
	return req
}

type InoutSummaryResponse struct {
	Total InoutTotal   `json:"total"`
	List  []*InoutList `json:"list"`
}
