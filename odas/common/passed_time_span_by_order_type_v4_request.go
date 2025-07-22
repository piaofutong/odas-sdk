package common

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// PassedTimeSpanByOrderTypeV4Request 按订单类型的时间跨度请求V4
type PassedTimeSpanByOrderTypeV4Request struct {
	// PassedTimeSpanRequest 字段
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
	Type  int64     `json:"type,omitempty"`
	Sid   []int64   `json:"sid,omitempty"`
	// PassedTimeSpanV4Request 字段
	Lid        []int64 `json:"lid,omitempty"` // 产品id
	ExcludeLid []int64 `json:"excludeLid,omitempty"`
	// 本结构字段
	OrderType int64 `json:"orderType,omitempty"`
}

func (p PassedTimeSpanByOrderTypeV4Request) Params(prefix string) url.Values {
	params := url.Values{}
	// PassedTimeSpanRequest 参数
	if !p.Start.IsZero() {
		params.Add("start", p.Start.Format("2006-01-02 15:04:05"))
	}
	if !p.End.IsZero() {
		params.Add("end", p.End.Format("2006-01-02 15:04:05"))
	}
	if p.Type > 0 {
		params.Add("type", strconv.FormatInt(p.Type, 10))
	}
	if len(p.Sid) > 0 {
		sidStrs := make([]string, len(p.Sid))
		for i, sid := range p.Sid {
			sidStrs[i] = strconv.FormatInt(sid, 10)
		}
		params.Add("sid", strings.Join(sidStrs, ","))
	}
	// PassedTimeSpanV4Request 参数
	if len(p.Lid) > 0 {
		lidStrs := make([]string, len(p.Lid))
		for i, lid := range p.Lid {
			lidStrs[i] = strconv.FormatInt(lid, 10)
		}
		params.Add("lid", strings.Join(lidStrs, ","))
	}
	if len(p.ExcludeLid) > 0 {
		excludeLidStrs := make([]string, len(p.ExcludeLid))
		for i, excludeLid := range p.ExcludeLid {
			excludeLidStrs[i] = strconv.FormatInt(excludeLid, 10)
		}
		params.Add("excludeLid", strings.Join(excludeLidStrs, ","))
	}
	// 本结构参数
	if p.OrderType > 0 {
		params.Add("orderType", strconv.FormatInt(p.OrderType, 10))
	}
	return params
}

func (p PassedTimeSpanByOrderTypeV4Request) Body() []byte {
	return nil
}

func (p PassedTimeSpanByOrderTypeV4Request) Method() string {
	return http.MethodGet
}

func (p PassedTimeSpanByOrderTypeV4Request) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (p PassedTimeSpanByOrderTypeV4Request) AuthRequired() bool {
	return true
}
