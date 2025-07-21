package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type SecondaryChannelOptions struct {
	ChannelClassId int `json:"channelClassId"`
	Limit          int `json:"limit"`
}

type SecondaryChannelOption func(options *SecondaryChannelOptions)

func WithSecondaryChannelClassId(channelClassId int) SecondaryChannelOption {
	return func(options *SecondaryChannelOptions) {
		options.ChannelClassId = channelClassId
	}
}

func WithSecondaryChannelLimit(limit int) SecondaryChannelOption {
	return func(options *SecondaryChannelOptions) {
		options.Limit = limit
	}
}

type OrderSecondaryChannelReq struct {
	*odas.Req
	Options *SecondaryChannelOptions
}

func NewOrderSecondaryChannel(req *odas.Req, opt ...SecondaryChannelOption) *OrderSecondaryChannelReq {
	options := &SecondaryChannelOptions{}
	for _, p := range opt {
		p(options)
	}
	return &OrderSecondaryChannelReq{
		Req:     req,
		Options: options,
	}
}

func (o OrderSecondaryChannelReq) Api() string {
	params := o.Req.Params()
	if o.Options.ChannelClassId > 0 {
		params.Add("channelClassId", strconv.Itoa(o.Options.ChannelClassId))
	}
	if o.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(o.Options.Limit))
	}
	return fmt.Sprintf("/v4/channel/orderSecondaryChannel?%s", params.Encode())
}
