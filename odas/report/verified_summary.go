package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// VerifiedSummaryReq 验证订单数据
type VerifiedSummaryReq struct {
	odas.Req
}

func NewVerifiedSummaryReq(req *odas.Req) *VerifiedSummaryReq {
	return &VerifiedSummaryReq{
		Req: *req,
	}
}

func (req *VerifiedSummaryReq) Api() string {
	params := req.Req.Params()
	return fmt.Sprintf("/v4/report/verifiedSummary?%s", params.Encode())
}

type VerifiedSummaryResponse struct {
	odas.BaseReportSummaryVO
	CalcTicketNum int     `json:"calcTicketNum"`
	CalcOrderNum  int     `json:"calcOrderNum"`
	CalcAmount    float64 `json:"calcAmount"`
}
