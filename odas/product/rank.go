package product

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// RankReq 产品排行数据
type RankReq struct {
	odas.Req
	Limit int `json:"limit"`
}

func NewRankReq(req *odas.Req, limit int) *RankReq {
	return &RankReq{
		Req:   *req,
		Limit: limit,
	}
}

func (r RankReq) Api() string {
	params := r.Req.Params()
	if r.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
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
