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
	ps := fmt.Sprintf("/v4/report/ticketList?%s", params.Encode())
	if t.TicketId != "" {
		ps += fmt.Sprintf("&ticletId=%s", t.TicketId)
	}
	return ps
}

type TicketListResponse struct {
	Total *odas.BaseVerifiedSummaryVO `json:"total"`
	List  []*TicketListData           `json:"list"`
}

type TicketListData struct {
	TicketId   int    `json:"ticketId"`
	TicketName string `json:"ticketName"`
	odas.BaseVerifiedSummaryVO
}
