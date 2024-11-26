package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type DailyPassengerFlowReq struct {
	odas.Req
	Unknown bool `form:"unknown" json:"unknown" binding:"omitempty"`
}

func NewDailyPassengerFlowReq(req *odas.Req, unknown bool) *DailyPassengerFlowReq {
	return &DailyPassengerFlowReq{
		Req:     *req,
		Unknown: unknown,
	}
}

func (r DailyPassengerFlowReq) Api() string {
	params := r.Req.Params()
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	return fmt.Sprintf("/v4/tourist/dailyPassengerFlow?%s", params.Encode())
}

type PassengerFlowByDateResponse struct {
	Total int `json:"total"`
	List  []*PassengerFlowByDateList
}

type PassengerFlowByDateList struct {
	Count int `json:"count"`
	Date  int `json:"date"`
}
