package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type LimitOptions struct {
	Limit int
}

type LimitOption func(options *LimitOptions)

func WithLimit(limit int) LimitOption {
	return func(options *LimitOptions) {
		options.Limit = limit
	}
}

type OrderFullChannelReq struct {
	odas.Req
	Options *LimitOptions
}

func NewOrderFullChannelReq(req *odas.Req, opt ...LimitOption) *OrderFullChannelReq {
	options := &LimitOptions{}
	for _, p := range opt {
		p(options)
	}

	return &OrderFullChannelReq{
		Req:     *req,
		Options: options,
	}
}

func (o OrderFullChannelReq) Api() string {
	params := o.Req.Params()
	if o.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(o.Options.Limit))
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
