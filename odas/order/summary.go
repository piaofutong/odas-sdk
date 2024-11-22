package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// Summary 订单单量、票数、金额及同环比数据
type Summary struct {
	*odas.Req
	Compare bool `json:"compare"`
}

func NewSummaryReq(req *odas.Req, compare bool) *Summary {
	return &Summary{Req: req, Compare: compare}
}

func (r *Summary) Api() string {
	params := r.Req.Params()
	if r.Compare {
		params.Add("compare", strconv.FormatBool(r.Compare))
	}
	return fmt.Sprintf("/v4/order/summary?%s", params.Encode())
}

type SummaryResponse struct {
	OrderTicket        int      `json:"orderTicket"`
	MomOrderTicket     *float64 `json:"momOrderTicket"`
	YoyOrderTicket     *float64 `json:"yoyOrderTicket"`
	VerifyTicket       int      `json:"verifyTicket"`
	MomVerifyTicket    *float64 `json:"momVerifyTicket"`
	YoyVerifyTicket    *float64 `json:"yoyVerifyT"`
	RefundTicket       int      `json:"refundTicket"`
	MomRefundTicket    *float64 `json:"momRefundTicket"`
	YoyRefundTicket    *float64 `json:"yoyRefundTicket"`
	FinishTicket       int      `json:"finishTicket"`
	MomFinishTicket    *float64 `json:"momFinishTicket"`
	YoyFinishTicket    *float64 `json:"yoyFinishTicket"`
	CancelTicket       int      `json:"cancelTicket"`
	MomCancelTicket    *float64 `json:"momCancelTicket"`
	YoyCancelTicket    *float64 `json:"yoyCancelTicket"`
	AfterSaleTicket    int      `json:"afterSaleTicket"`
	MomAfterSaleTicket *float64 `json:"momAfterSaleTicket"`
	YoyAfterSaleTicket *float64 `json:"yoyAfterSaleTicket"`
	OrderAmount        int      `json:"orderAmount"`
	MomOrderAmount     *float64 `json:"momOrderAmount"`
	YoyOrderAmount     *float64 `json:"yoyOrderAmount"`
	VerifyAmount       int      `json:"verifyAmount"`
	MomVerifyAmount    *float64 `json:"momVerifyAmount"`
	YoyVerifyAmount    *float64 `json:"yoyVerifyAmount"`
	RefundAmount       int      `json:"refundAmount"`
	MomRefundAmount    *float64 `json:"momRefundAmount"`
	YoyRefundAmount    *float64 `json:"yoyRefundAmount"`
	FinishAmount       int      `json:"finishAmount"`
	MomFinishAmount    *float64 `json:"momFinishAmount"`
	YoyFinishAmount    *float64 `json:"yoyFinishAmount"`
	CancelAmount       int      `json:"cancelAmount"`
	MomCancelAmount    *float64 `json:"momCancelAmount"`
	YoyCancelAmount    *float64 `json:"yoyCancelAmount"`
	AfterSaleAmount    int      `json:"afterSaleAmount"`
	MomAfterSaleAmount *float64 `json:"momAfterSaleAmount"`
	YoyAfterSaleAmount *float64 `json:"yoyAfterSaleAmount"`
}
