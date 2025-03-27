package sixun

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// SaleTotalByTimeRangeReq 总营收
type SaleTotalByTimeRangeReq struct {
	odas.Req
}

func NewSaleTotalByTimeRangeReq(req *odas.Req) *SaleTotalByTimeRangeReq {
	return &SaleTotalByTimeRangeReq{
		Req: *req,
	}
}

func (h SaleTotalByTimeRangeReq) Api() string {
	params := h.Req.Params()
	return fmt.Sprintf("/v4/sixun/saleTotalByTimeRange?%s", params.Encode())
}

type SaleTotalByTimeRangeResponse struct {
	TotalAmount int `json:"totalAmount"`
}
