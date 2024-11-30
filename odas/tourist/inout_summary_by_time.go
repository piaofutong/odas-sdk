package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SummaryByTimeReq struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Sid      int    `json:"sid"`
	Gid      string `json:"gid"`
	NoAmend  bool   `json:"noAmend"`
	DateType int    `json:"dateType"`
}

type SummaryReqOptions func(options *SummaryByTimeReq)

func WithDateType(dateType int) SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.DateType = dateType
	}
}

func WithNoAmend() SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.NoAmend = true
	}
}

func WithStart(start string) SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.Start = start
	}
}

func WithEnd(end string) SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.End = end
	}
}

func WithSid(sid int) SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.Sid = sid
	}
}

func WithGid(gid string) SummaryReqOptions {
	return func(options *SummaryByTimeReq) {
		options.Gid = gid
	}
}

func (s SummaryByTimeReq) Api() string {
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
	if s.Gid != "" {
		params.Add("gid", s.Gid)
	}
	if s.NoAmend {
		params.Add("noAmend", strconv.FormatBool(s.NoAmend))
	}
	if s.DateType > 0 {
		params.Add("dateType", strconv.Itoa(s.DateType))
	}
	return fmt.Sprintf("/v4/tourist/inout/summaryByTime?%s", params.Encode())
}

func (s SummaryByTimeReq) Body() []byte {
	return nil
}

func (s SummaryByTimeReq) Method() string {
	return http.MethodGet
}

func (s SummaryByTimeReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (s SummaryByTimeReq) AuthRequired() bool {
	return true
}

func NewSummaryByTimeReq(opt ...SummaryReqOptions) *SummaryByTimeReq {
	req := &SummaryByTimeReq{}
	for _, options := range opt {
		options(req)
	}
	return req
}

type InoutSummaryResponse struct {
	Total InoutTotal   `json:"total"`
	List  []*InoutList `json:"list"`
}
