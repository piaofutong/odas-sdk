package product

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// TicketListReq 查询票列表数据
type TicketListReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewTicketListReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	page, pageSize int,
) *TicketListReq {
	return &TicketListReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Page:                page,
		PageSize:            pageSize,
	}
}

func (req *TicketListReq) Api() string {
	params := req.Req.Params()
	if req.Page > 0 {
		params.Add("page", strconv.Itoa(req.Page))
	}
	if req.PageSize > 0 {
		params.Add("pageSize", strconv.Itoa(req.PageSize))
	}
	if req.CompareStart != "" {
		params.Add("compareStart", req.CompareStart)
	}
	if req.CompareEnd != "" {
		params.Add("compareEnd", req.CompareEnd)
	}
	return fmt.Sprintf("/v4/product/ticketList?%s", params.Encode())
}

type TicketListResponse struct {
	List       []*TicketList   `json:"list"`
	Pagination odas.Pagination `json:"pagination"`
}

type TicketList struct {
	TicketId          int     `json:"ticketId"`
	TicketName        string  `json:"ticketName"`
	Count             int     `json:"count"`
	CompareCountRate  float64 `json:"compareCountRate"`
	Amount            int     `json:"amount"`
	CompareAmountRate float64 `json:"compareAmountRate"`
}
