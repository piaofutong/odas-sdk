package hotel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type RoomSaleReportListReq struct {
	odas.DateRangeReq
}

func (r RoomSaleReportListReq) Api() string {
	params := r.DateRangeReq.Api()
	return fmt.Sprintf("/v2/hotel/rmSaleReportList?%s", params.Encode())
}

func NewRmSaleReportListReq(req *odas.DateRangeReq) *RoomSaleReportListReq {
	return &RoomSaleReportListReq{DateRangeReq: *req}
}

type RmSaleReportListResponse struct {
	Total *RmSaleReportTotal      `json:"total"`
	List  []*RmSaleReportListData `json:"list"`
}

type RmSaleReportListData struct {
	BindId int    `json:"bindId"`
	Name   string `json:"name"`
	RmSaleReportTotal
}
