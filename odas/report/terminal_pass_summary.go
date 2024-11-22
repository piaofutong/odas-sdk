package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// TerminalPassSummaryReq 获取时间段终端验证汇总数据
type TerminalPassSummaryReq struct {
	odas.Req
}

func NewTerminalPassSummaryReq(req *odas.Req) *TerminalPassSummaryReq {
	return &TerminalPassSummaryReq{
		Req: *req,
	}
}

func (req *TerminalPassSummaryReq) Api() string {
	params := req.Req.Params()
	return fmt.Sprintf("/v4/report/terminalPassSummary?%s", params.Encode())
}

type TerminalPassSummaryResponse struct {
	Total *TerminalPassHourTotal `json:"total"`
	List  []*TerminalPassList    `json:"list"`
}

type TerminalPassList struct {
	Date string `json:"date"`
	TerminalPassHourTotal
}
