package common

import (
	"net/http"
	"net/url"
	"time"
)

// PassedTimeSpan 时间跨度
type PassedTimeSpan struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}

func (p PassedTimeSpan) Params(prefix string) url.Values {
	params := url.Values{}
	if !p.Start.IsZero() {
		params.Add(prefix+"start", p.Start.Format("2006-01-02 15:04:05"))
	}
	if !p.End.IsZero() {
		params.Add(prefix+"end", p.End.Format("2006-01-02 15:04:05"))
	}
	return params
}

func (p PassedTimeSpan) Body() []byte {
	return nil
}

func (p PassedTimeSpan) Method() string {
	return http.MethodGet
}

func (p PassedTimeSpan) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (p PassedTimeSpan) AuthRequired() bool {
	return true
}

// PassedTimeBetween represents a time range for passed time
type PassedTimeBetween struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// Params 返回查询参数
func (p *PassedTimeBetween) Params(prefix string) url.Values {
	params := url.Values{}
	if !p.Start.IsZero() {
		params.Add(prefix+"start", p.Start.Format("2006-01-02 15:04:05"))
	}
	if !p.End.IsZero() {
		params.Add(prefix+"end", p.End.Format("2006-01-02 15:04:05"))
	}
	return params
}
