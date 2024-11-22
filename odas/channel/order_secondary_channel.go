package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type OrderSecondaryChannelReq struct {
	*odas.Req
	ChannelClassId int `json:"channelClassId"`
	Limit          int `json:"limit"`
}

func NewOrderSecondaryChannel(req *odas.Req, channelClassId int, limit int) *OrderSecondaryChannelReq {
	return &OrderSecondaryChannelReq{
		Req:            req,
		ChannelClassId: channelClassId,
		Limit:          limit,
	}
}

func (o OrderSecondaryChannelReq) Api() string {
	params := o.Req.Params()
	if o.ChannelClassId > 0 {
		params.Add("channelClassId", strconv.Itoa(o.ChannelClassId))
	}
	if o.Limit > 0 {
		params.Add("limit", strconv.Itoa(o.Limit))
	}
	return fmt.Sprintf("/v4/channel/orderSecondaryChannel?%s", params.Encode())
}
