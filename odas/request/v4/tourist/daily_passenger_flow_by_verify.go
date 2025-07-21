package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type DailyPassengerFlowByVerifyReq struct {
	odas.Req
	Unknown bool `form:"unknown" json:"unknown" binding:"omitempty"`
}

func NewDailyPassengerFlowByVerifyReq(req *odas.Req, unknown bool) *DailyPassengerFlowByVerifyReq {
	return &DailyPassengerFlowByVerifyReq{
		Req:     *req,
		Unknown: unknown,
	}
}

func (r DailyPassengerFlowByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	return fmt.Sprintf("/v4/tourist/dailyPassengerFlowByVerify?%s", params.Encode())
}
