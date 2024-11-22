package hotel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type RmSaleReportDateListReq struct {
	odas.DateRangeReq
}

func (r RmSaleReportDateListReq) Api() string {
	params := r.DateRangeReq.Api()
	return fmt.Sprintf("/v2/hotel/rmSaleReportDateList?%s", params.Encode())
}

func NewRmSaleReportDateListReq(req *odas.DateRangeReq) *RmSaleReportDateListReq {
	return &RmSaleReportDateListReq{DateRangeReq: *req}
}

type RmSaleReportDateListResponse struct {
	Total *RmSaleReportTotal          `json:"total"`
	List  []*RmSaleReportDateListData `json:"list"`
}

type RmSaleReportTotal struct {
	RoomsTotal float64 `json:"roomsTotal"` // 房间总数
	RoomsOoo   float64 `json:"roomsOoo"`   // 空房数
	RoomsOs    float64 `json:"roomsOs"`    // 锁房数
	RoomsHse   float64 `json:"roomsHse"`   // 自用房
	RoomsAvl   float64 `json:"roomsAvl"`   // 可用房
	RoomsVac   float64 `json:"roomsVac"`   // 空房
	SoldFit    float64 `json:"soldFit"`    // 散客
	SoldGrp    float64 `json:"soldGrp"`    // 团队
	SoldLong   float64 `json:"soldLong"`   // 长包
	SoldEnt    float64 `json:"soldEnt"`    // 免费
	RevFit     float64 `json:"revFit"`     // 散客房费
	RevGrp     float64 `json:"revGrp"`     // 团队房费
	RevLong    float64 `json:"revLong"`    // 长包房费
	PeopleFit  int     `json:"peopleFit"`  // 散客人数
	PeopleGrp  int     `json:"peopleGrp"`  // 团队人数
	PeopleLong int     `json:"peopleLong"` // 长包人数
}

type RmSaleReportDateListData struct {
	Date string `json:"date"`
	RmSaleReportTotal
}
