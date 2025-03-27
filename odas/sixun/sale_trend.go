package sixun

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// SaleTrendReq 营收趋势
type SaleTrendReq struct {
	odas.Req
}

func NewSaleTrendReq(req *odas.Req) *SaleTrendReq {
	return &SaleTrendReq{
		Req: *req,
	}
}

func (h SaleTrendReq) Api() string {
	params := h.Req.Params()
	return fmt.Sprintf("/v4/sixun/saleTrend?%s", params.Encode())
}

type SaleTrendListItem struct {
	Amount     int    `json:"amount"`
	OrderNum   int    `json:"orderNum"`
	Time       int    `json:"time"`
	FormatTime string `json:"formatTime"`
}

type SaleTrendResponse struct {
	List []*SaleTrendListItem `json:"list"`
}
