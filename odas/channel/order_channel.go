package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type OrderChannelReq struct {
	odas.Req
}

func (o OrderChannelReq) Api() string {
	params := o.Req.Params()
	a := fmt.Sprintf("/v4/channel/orderChannel?%s", params.Encode())
	return a
}

func NewOrderChannelReq(req *odas.Req) *OrderChannelReq {
	return &OrderChannelReq{Req: *req}
}

type OrderChannelResponse struct {
	ChannelName string  `json:"channelName"`
	Tickets     int     `json:"tickets"`
	Amount      int     `json:"amount"`
	Rate        float64 `json:"rate"`
}
