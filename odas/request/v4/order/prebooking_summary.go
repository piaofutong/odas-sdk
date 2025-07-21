package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type PreBookingSummaryReq struct {
	odas.DateRangeReq
	Options *PreBookingOptions
}

func (p PreBookingSummaryReq) Api() string {
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

func NewPreBookingSummaryReq(req odas.DateRangeReq, opt ...PreBookingByTypeOption) *PreBookingSummaryReq {
	options := &PreBookingOptions{}
	for _, p := range opt {
		p(options)
	}
	return &PreBookingSummaryReq{
		DateRangeReq: req,
		Options:      options,
	}
}

type PreBookingSummaryResponse struct {
	PreBookingTotal
	List []*PreBookingSummaryDateList `json:"list"`
}

type PreBookingSummaryDateList struct {
	Time string `json:"time"`
	PreBookingTotal
}
