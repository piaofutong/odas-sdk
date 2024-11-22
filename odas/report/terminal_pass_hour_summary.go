package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// TerminalPassHourSummaryReq 获取终端验证数小时数据
type TerminalPassHourSummaryReq struct {
	odas.Req
}

func NewTerminalPassHourSummaryReq(req *odas.Req) *TerminalPassHourSummaryReq {
	return &TerminalPassHourSummaryReq{
		Req: *req,
	}
}

func (req *TerminalPassHourSummaryReq) Api() string {
	params := req.Req.Params()
	return fmt.Sprintf("/v4/report/terminalPassHourSummary?%s", params.Encode())
}

type TerminalPassHourSummaryResponse struct {
	Total *TerminalPassHourTotal  `json:"total"`
	List  []*TerminalPassHourList `json:"list"`
}

type TerminalPassHourTotal struct {
	VerifyTicket    int `json:"verifyTicket"`
	VerifySaleMoney int `json:"verifySaleMoney"`
}

type TerminalPassHourList struct {
	Time string `json:"time"`
	TerminalPassHourTotal
}
