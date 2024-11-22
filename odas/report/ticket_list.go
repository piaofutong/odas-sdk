package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// TicketListReq 票务类型统计
type TicketListReq struct {
	odas.Req
	Limit    int    `json:"limit"`
	TicketId string `json:"ticketId"`
}

func NewTicketListReq(req *odas.Req, limit int, ticketId string) *TicketListReq {
	return &TicketListReq{
		Req:      *req,
		Limit:    limit,
		TicketId: ticketId,
	}
}

func (t TicketListReq) Api() string {
	params := t.Req.Params()
	if t.Limit > 0 {
		params.Add("limit", strconv.Itoa(t.Limit))
	}
	if t.TicketId != "" {
		params.Add("ticketId", t.TicketId)
	}
	return fmt.Sprintf("/v4/report/ticketList?%s", params.Encode())
}

type TicketListResponse struct {
	Total *odas.BaseReportSummaryVO `json:"total"`
	List  []*TicketListData         `json:"list"`
}

type TicketListData struct {
	TicketId   int    `json:"ticketId"`
	TicketName string `json:"ticketName"`
	odas.BaseReportSummaryVO
}
