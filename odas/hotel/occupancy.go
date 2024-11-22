package hotel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type OccupancyReq struct {
	odas.DateRangeReq
}

func NewOccupancyReq(req *odas.DateRangeReq) *OccupancyReq {
	return &OccupancyReq{DateRangeReq: *req}
}

func (r OccupancyReq) Api() string {
	params := r.DateRangeReq.Api()
	return fmt.Sprintf("/v2/hotel/occupancy?%s", params.Encode())
}

type OccupancyResponse int
