package product

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type RankOptions struct {
	Limit int
}

type RankOption func(options *RankOptions)

func WithRankLimit(limit int) RankOption {
	return func(options *RankOptions) {
		options.Limit = limit
	}
}

// RankReq 产品排行数据
type RankReq struct {
	odas.Req
	Options *RankOptions
}

func NewRankReq(req *odas.Req, opt ...RankOption) *RankReq {
	options := &RankOptions{}
	for _, p := range opt {
		p(options)
	}
	return &RankReq{
		Req:     *req,
		Options: options,
	}
}

func (r RankReq) Api() string {
	params := r.Req.Params()
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	return fmt.Sprintf("/v4/product/rank?%s", params.Encode())
}

type RankResponse struct {
	TicketId   int     `json:"ticketId"`
	TicketName string  `json:"ticketName"`
	Count      int     `json:"count"`
	Amount     int     `json:"amount"`
	Rate       float64 `json:"rate"`
}
