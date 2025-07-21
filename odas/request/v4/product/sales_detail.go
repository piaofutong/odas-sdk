package product

import (
	"encoding/json"
	"github.com/piaofutong/odas-sdk/odas"
	"net/http"
)

type SalesDetailOptions struct {
	TicketId []int `json:"ticketId"`
}

type SalesDetailOption func(options *SalesDetailOptions)

func WithSalesDetailTicketId(ticketId []int) SalesDetailOption {
	return func(options *SalesDetailOptions) {
		options.TicketId = ticketId
	}
}

// SalesDetailReq 根据票数据获取渠道以及销售额每日数据
type SalesDetailReq struct {
	odas.Req
	*SalesDetailOptions
}

func NewSalesDetailReq(req *odas.Req, opt ...SalesDetailOption) *SalesDetailReq {
	options := &SalesDetailOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SalesDetailReq{
		Req:                *req,
		SalesDetailOptions: options,
	}
}

func (s SalesDetailReq) Api() string {
	return "/v4/product/salesDetail"
}

func (s SalesDetailReq) Body() []byte {
	body, _ := json.Marshal(s)
	return body
}

func (s SalesDetailReq) Method() string {
	return http.MethodPost
}

func (s SalesDetailReq) ContentType() string {
	return "application/json"
}

type SalesDetailResponse struct {
	TicketId    int            `json:"ticketId"`
	AmountTrend []*AmountTrend `json:"amountTrend"`
	ChannelList []*ChannelList `json:"channelList"`
}

type AmountTrend struct {
	TimeRange   string `json:"timeRange"`
	TicketCount int    `json:"ticketCount"`
}

type ChannelList struct {
	ChannelName string  `json:"channelName"`
	TicketCount int     `json:"ticketCount"`
	Rate        float64 `json:"rate"`
}
