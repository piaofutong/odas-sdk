package order

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type LandListReq struct {
	odas.Req
}

func NewLandListReq(req *odas.Req) *LandListReq {
	return &LandListReq{Req: *req}
}

func (r *LandListReq) Api() string {
	params := r.Req.Params()
	return fmt.Sprintf("/v4/order/booking/orderList?%s", params.Encode())
}

type LandListResponse struct {
	Total  *BookingOrderTotal `json:"total"`
	Detail []*LandListDetail  `json:"detail"`
}

type BookingOrderTotal struct {
	OrderNum             int `json:"orderNum"`
	OrderTicket          int `json:"orderTicket"`
	OrderAmount          int `json:"orderAmount"`
	OrderCostMoney       int `json:"orderCostMoney"`
	VerifiedNum          int `json:"verifiedNum"`
	VerifiedTicket       int `json:"verifiedTicket"`
	VerifiedAmount       int `json:"verifiedAmount"`
	VerifiedCostMoney    int `json:"verifiedCostMoney"`
	FinishedNum          int `json:"finishedNum"`
	FinishedTicket       int `json:"finishedTicket"`
	FinishedAmount       int `json:"finishedAmount"`
	FinishedCostMoney    int `json:"finishedCostMoney"`
	RevokedNum           int `json:"revokedNum"`
	RevokedTicket        int `json:"revokedTicket"`
	RevokedAmount        int `json:"revokedAmount"`
	RevokedCostMoney     int `json:"revokedCostMoney"`
	CancelNum            int `json:"cancelNum"`
	CancelTicket         int `json:"cancelTicket"`
	CancelAmount         int `json:"cancelAmount"`
	CancelCostMoney      int `json:"cancelCostMoney"`
	AfterSaleTicketNum   int `json:"afterSaleTicketNum"`
	AfterSaleRefundMoney int `json:"afterSaleRefundMoney"`
	AfterSaleIncomeMoney int `json:"afterSaleIncomeMoney"`
	PrintNum             int `json:"printNum"`
}

type LandListDetail struct {
	Time int `json:"time"`
	BookingOrderTotal
}
