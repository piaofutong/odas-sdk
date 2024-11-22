package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type OrderFullChannelReq struct {
	odas.Req
	Limit int `json:"limit"`
}

func NewOrderFullChannelReq(req *odas.Req, limit int) *OrderFullChannelReq {
	return &OrderFullChannelReq{
		Req:   *req,
		Limit: limit,
	}
}

func (o OrderFullChannelReq) Api() string {
	params := o.Req.Params()
	if o.Limit > 0 {
		params.Add("limit", strconv.Itoa(o.Limit))
	}
	return fmt.Sprintf("/v4/channel/orderFullChannel?%s", params.Encode())
}

type OrderFullChannelResponse struct {
	Total *OrderChannelTotal      `json:"total"`
	List  []*OrderFullChannelList `json:"list"`
}

type OrderChannelTotal struct {
	OrderCount  int `json:"orderCount"`
	TicketCount int `json:"ticketCount"`
	Amount      int `json:"amount"`
}

type OrderFullChannelList struct {
	ChannelClassId   int    `json:"channelClassId"`
	ChannelClassName string `json:"channelClassName"`
	OrderCount       int    `json:"orderCount"`
	TicketCount      int    `json:"ticketCount"`
	Amount           int    `json:"amount"`
}
