package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// VerifiedSummaryHourReq 验证订单小时数据
type VerifiedSummaryHourReq struct {
	odas.Req
}

func NewVerifiedSummaryHourReq(req *odas.Req) *VerifiedSummaryHourReq {
	return &VerifiedSummaryHourReq{
		Req: *req,
	}
}

func (req *VerifiedSummaryHourReq) Api() string {
	params := req.Req.Params()
	return fmt.Sprintf("/v4/report/verifiedSummaryByHour?%s", params.Encode())
}

type VerifiedSummaryHourResponse struct {
	Total *VerifiedSummaryResponse   `json:"total"`
	List  []*VerifiedSummaryHourList `json:"list"`
}

type VerifiedSummaryHourList struct {
	Hour string `json:"hour"`
	VerifiedSummaryResponse
}
