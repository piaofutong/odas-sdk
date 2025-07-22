package common

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// PassedTimeSpanRequest 时间跨度请求
type PassedTimeSpanRequest struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
	Type  int64     `json:"dateType,omitempty"`
	Sid   []int64   `json:"sid,omitempty"`
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
		params.Add(prefix+"dateType", strconv.FormatInt(p.Type, 10))
	}
	if len(p.Sid) > 0 {
		sidStrs := make([]string, len(p.Sid))
		for i, sid := range p.Sid {
			sidStrs[i] = strconv.FormatInt(sid, 10)
		}
		params.Add(prefix+"sid", strings.Join(sidStrs, ","))
	}
	return params
}
