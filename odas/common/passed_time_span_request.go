package common

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// PassedTimeSpanRequest 时间跨度请求
type PassedTimeSpanRequest struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
	Type  int       `json:"type,omitempty"`
	Sid   []int     `json:"sid,omitempty"`
}

func (p PassedTimeSpanRequest) Params(prefix string) url.Values {
	params := url.Values{}
	if !p.Start.IsZero() {
		params.Add(prefix+"start", p.Start.Format("2006-01-02 15:04:05"))
	}
	if !p.End.IsZero() {
		params.Add(prefix+"end", p.End.Format("2006-01-02 15:04:05"))
	}
	if p.Type > 0 {
		params.Add(prefix+"type", strconv.Itoa(p.Type))
	}
	if len(p.Sid) > 0 {
		sidStrs := make([]string, len(p.Sid))
		for i, sid := range p.Sid {
			sidStrs[i] = strconv.Itoa(sid)
		}
		params.Add(prefix+"sid", strings.Join(sidStrs, ","))
	}
	return params
}

func (p PassedTimeSpanRequest) Body() []byte {
	return nil
}

func (p PassedTimeSpanRequest) Method() string {
	return http.MethodGet
}

func (p PassedTimeSpanRequest) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (p PassedTimeSpanRequest) AuthRequired() bool {
	return true
}
