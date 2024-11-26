package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type TerminalPassSummaryOptions struct {
	TerminalType string `json:"terminalType"`
}

type TerminalPassSummaryOption func(options *TerminalPassSummaryOptions)

func WithTerminalType(terminalType string) func(options *TerminalPassSummaryOptions) {
	return func(options *TerminalPassSummaryOptions) {
		options.TerminalType = terminalType
	}
}

// TerminalPassSummaryReq 获取时间段终端验证汇总数据
type TerminalPassSummaryReq struct {
	odas.Req
	Options *TerminalPassSummaryOptions
}

func NewTerminalPassSummaryReq(req *odas.Req, opt ...TerminalPassSummaryOption) *TerminalPassSummaryReq {
	options := &TerminalPassSummaryOptions{}
	for _, p := range opt {
		p(options)
	}

	return &TerminalPassSummaryReq{
		Req:     *req,
		Options: options,
	}
}

func (req *TerminalPassSummaryReq) Api() string {
	params := req.Req.Params()

	if req.Options.TerminalType != "" {
		params.Add("terminalType", req.Options.TerminalType)
	}

	return fmt.Sprintf("/v4/report/terminalPassSummary?%s", params.Encode())
}

type TerminalPassSummaryResponse struct {
	Total *TerminalPassTotal  `json:"total"`
	List  []*TerminalPassList `json:"list"`
}

type TerminalPassTotal struct {
	VerifyTicket    int `json:"verifyTicket"`
	VerifySaleMoney int `json:"verifySaleMoney"`
}

type TerminalPassList struct {
	Date string `json:"date"`
	TerminalPassTotal
}
