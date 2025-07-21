package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// HotReq 查询热门景区订单数据
type HotReq struct {
	odas.Req
	Limit int `json:"limit"`
}

func NewHotReq(req *odas.Req, limit int) *HotReq {
	return &HotReq{
		Req:   *req,
		Limit: limit,
	}
}

func (h HotReq) Api() string {
	params := h.Req.Params()
	if h.Limit > 0 {
		params.Add("limit", strconv.Itoa(h.Limit))
	}
	return fmt.Sprintf("/v4/order/hot?%s", params.Encode())
}

type HotResponse struct {
	Lid         int     `json:"lid"`
	TicketCount int     `json:"ticketCount"`
	OrderCount  int     `json:"orderCount"`
	Amount      int     `json:"amount"`
	UnitPrice   float64 `json:"unitPrice"`
}
