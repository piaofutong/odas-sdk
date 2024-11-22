package product

import (
	"encoding/json"
	"github.com/piaofutong/odas-sdk/odas"
	"net/http"
)

// SalesDetailReq 根据票数据获取渠道以及销售额每日数据
type SalesDetailReq struct {
	odas.Req
	TicketId []int `json:"ticketId"`
}

func NewSalesDetailReq(req *odas.Req, ticketId []int) *SalesDetailReq {
	return &SalesDetailReq{
		Req:      *req,
		TicketId: ticketId,
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
