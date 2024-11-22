package hotel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type RoomOrderDateListReq struct {
	odas.DateRangeReq
}

func NewRmOrderDateListReq(req *odas.DateRangeReq) *RoomOrderDateListReq {
	return &RoomOrderDateListReq{DateRangeReq: *req}
}

func (r RoomOrderDateListReq) Api() string {
	params := r.DateRangeReq.Api()
	return fmt.Sprintf("/v2/hotel/rmOrderDateList?%s", params.Encode())
}

type RmOrderDateListResponse struct {
	Total *RmOrderDateTotal      `json:"total"`
	List  []*RmOrderDateListData `json:"list"`
}

type RmOrderDateTotal struct {
	BookingCount    int `json:"bookingCount"`
	BookingRoomNum  int `json:"bookingRoomNum"`
	BookingAdult    int `json:"bookingAdult"`
	BookingChildren int `json:"bookingChildren"`
	BookingPeople   int `json:"bookingPeople"`
	BookingCharge   int `json:"bookingCharge"`
	BookingPay      int `json:"bookingPay"`
	CheckInCount    int `json:"checkInCount"`
	CheckInRoomNum  int `json:"checkInRoomNum"`
	CheckInAdult    int `json:"checkInAdult"`
	CheckInChildren int `json:"checkInChildren"`
	CheckInPeople   int `json:"checkInPeople"`
	CheckInCharge   int `json:"checkInCharge"`
	CheckInPay      int `json:"checkInPay"`
}

type RmOrderDateListData struct {
	Date string `json:"date"`
	RmOrderDateTotal
}
