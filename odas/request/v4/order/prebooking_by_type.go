package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type PreBookingOptions struct {
	Lid        string `json:"lid"`
	ExcludeLid string `json:"excludeLid"`
	OrderType  int    `json:"orderType"`
}

type PreBookingByTypeReq struct {
	odas.DateRangeReq
	Options *PreBookingOptions
}

func (p PreBookingByTypeReq) Api() string {
	params := p.DateRangeReq.Api()
	if p.Options.Lid != "" {
		params.Add("lid", p.Options.Lid)
	}
	if p.Options.ExcludeLid != "" {
		params.Add("excludeLid", p.Options.ExcludeLid)
	}
	if p.Options.OrderType > 0 {
		params.Add("orderType", strconv.Itoa(p.Options.OrderType))
	}

	return fmt.Sprintf("/v4/order/preBookingSummary?%s", params.Encode())
}

type PreBookingByTypeOption func(options *PreBookingOptions)

func WithLid(lid string) PreBookingByTypeOption {
	return func(options *PreBookingOptions) {
		options.Lid = lid
	}
}

func WithExcludeLid(excludeLid string) PreBookingByTypeOption {
	return func(options *PreBookingOptions) {
		options.ExcludeLid = excludeLid
	}
}

func WithOrderType(orderType int) PreBookingByTypeOption {
	return func(options *PreBookingOptions) {
		options.OrderType = orderType
	}
}

func NewPreBookingByTypeReq(req odas.DateRangeReq, opt ...PreBookingByTypeOption) *PreBookingByTypeReq {
	options := &PreBookingOptions{}
	for _, p := range opt {
		p(options)
	}
	return &PreBookingByTypeReq{
		DateRangeReq: req,
		Options:      options,
	}
}

type PreBookingByTypeResponse struct {
	Total *PreBookingTotal  `json:"total"`
	List  []*PreBookingList `json:"list"`
}

type PreBookingTotal struct {
	OrderNum    int `json:"orderNum"`
	OrderTicket int `json:"orderTicket"`
	OrderAmount int `json:"orderAmount"`
}

type PreBookingList struct {
	TypeName string  `json:"typeName"`
	Percent  float64 `json:"percent"`
	PreBookingTotal
}
