package channel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type StatDistributorSummaryReq struct {
	odas.Req
}

func NewStatDistributorSummaryReq(req *odas.Req) *StatDistributorSummaryReq {
	return &StatDistributorSummaryReq{Req: *req}
}

func (s StatDistributorSummaryReq) Api() string {
	params := s.Req.Params()
	return fmt.Sprintf("/v4/channel/statDistributorSummary?%s", params.Encode())
}

type StatDistributorSummaryResponse struct {
	DistributorID   int    `json:"distributor_id"`
	DistributorName string `json:"distributor_name"`
	Date            string `json:"date"`
	OrderCount      int    `json:"order_count"`
	TicketCount     int    `json:"ticket_count"`
	Amount          int    `json:"amount"`
}
