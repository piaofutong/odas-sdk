package report

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// TerminalPassSummaryGroupLidReq 获取时间段终端验证景区分组汇总数据
type TerminalPassSummaryGroupLidReq struct {
	odas.Req
	Options *TerminalPassSummaryOptions
}

func NewTerminalPassSummaryGroupLidReq(req *odas.Req, opt ...TerminalPassSummaryOption) *TerminalPassSummaryGroupLidReq {
	options := &TerminalPassSummaryOptions{}
	for _, p := range opt {
		p(options)
	}

	return &TerminalPassSummaryGroupLidReq{
		Req:     *req,
		Options: options,
	}
}

func (req *TerminalPassSummaryGroupLidReq) Api() string {
	params := req.Req.Params()

	if req.Options.TerminalType != "" {
		params.Add("terminalType", req.Options.TerminalType)
	}

	return fmt.Sprintf("/v4/report/terminalPassSummaryGroupLid?%s", params.Encode())
}

type TerminalPassSummaryGroupLidResponse struct {
	Total *TerminalPassTotal          `json:"total"`
	List  []*TerminalPassGroupLidList `json:"list"`
}

type TerminalPassGroupLidList struct {
	Lid int `json:"lid"`
	TerminalPassTotal
}
