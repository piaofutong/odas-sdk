package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// ToiSummaryReq 查询团散单汇总数据
type ToiSummaryReq struct {
	odas.Req
}

func NewToiSummaryReq(req *odas.Req) *ToiSummaryReq {
	return &ToiSummaryReq{Req: *req}
}

func (r *ToiSummaryReq) Api() string {
	params := r.Req.Params()
	return fmt.Sprintf("/v4/order/toi/summary?%s", params.Encode())
}

type ToiSummaryResponse struct {
	Total      *ToiTotal `json:"total"`
	Team       *ToiData  `json:"team"`
	Individual *ToiData  `json:"individual"`
}

type ToiTotal struct {
	Order  int `json:"order"`
	Ticket int `json:"ticket"`
	Amount int `json:"amount"`
}

type ToiData struct {
	Order      int     `json:"order"`
	Ticket     int     `json:"ticket"`
	Amount     int     `json:"amount"`
	OrderRate  float64 `json:"orderRate"`
	TicketRate float64 `json:"ticketRate"`
	AmountRate float64 `json:"amountRate"`
}
