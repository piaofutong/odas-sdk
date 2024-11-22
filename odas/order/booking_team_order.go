package order

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type BookingTeamOrderReq struct {
	odas.Req
}

func NewBookingTeamOrderReq(req *odas.Req) *BookingTeamOrderReq {
	return &BookingTeamOrderReq{Req: *req}
}

func (r *BookingTeamOrderReq) Api() string {
	params := r.Req.Params()
	return fmt.Sprintf("/v4/order/booking/teamOrder?%s", params.Encode())
}

type BookingTeamOrderResponse struct {
	Total  *TeamTotal  `json:"total"`
	Detail *TeamDetail `json:"detail"`
}

type TeamTotal struct {
	Team       *int `json:"team"`
	TeamTicket *int `json:"teamTicket"`
	TeamVerify *int `json:"teamVerify"`
}

type TeamDetail struct {
	TeamTrend       []*TeamTrend       `json:"teamTrend"`
	TeamTicketTrend []*TeamTicketTrend `json:"teamTicketTrend"`
	TeamAmountTrend []*TeamAmountTrend `json:"teamAmountTrend"`
}

type TeamTrend struct {
	Time        int     `json:"time"`
	Team        *int    `json:"team"`
	CompareTeam float64 `json:"compareTeam"`
}

type TeamTicketTrend struct {
	Time          int     `json:"time"`
	Ticket        *int    `json:"ticket"`
	CompareTicket float64 `json:"compareTicket"`
}

type TeamAmountTrend struct {
	Time          int     `json:"time"`
	Amount        *int    `json:"amount"`
	CompareAmount float64 `json:"compareAmount"`
}
