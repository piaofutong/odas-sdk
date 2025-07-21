package tourist

import (
	"fmt"
	"net/url"
	"strconv"
)

type SummaryByDateReq struct {
	SummaryByTimeReq
}

func (s SummaryByDateReq) Api() string {
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
	return fmt.Sprintf("/v4/tourist/inout/summaryByDate?%s", params.Encode())
}

func NewSummaryByDateReq(opt ...SummaryReqOptions) *SummaryByDateReq {
	req := &SummaryByTimeReq{}
	for _, options := range opt {
		options(req)
	}
	return &SummaryByDateReq{
		SummaryByTimeReq: *req,
	}
}
