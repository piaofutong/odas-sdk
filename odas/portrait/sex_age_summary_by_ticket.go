package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type SexAgeByTicketOptions struct {
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
}

type SexAgeByTicketOption func(options *SexAgeByTicketOptions)

// SexAgeSummaryByTicketReq 性别年龄分布
type SexAgeSummaryByTicketReq struct {
	odas.Req
	Sid int
}

func NewSexAgeSummaryByTicketReq(req *odas.Req, opt ...SexAgeByTicketOption) *SexAgeSummaryByTicketReq {
	options := &SexAgeByTicketOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SexAgeSummaryByTicketReq{
		Req: *req,
	}
}

func (r SexAgeSummaryByTicketReq) Api() string {
	params := r.Req.Params()

	return fmt.Sprintf("/v4/portrait/ageSummaryByTicket?%s", params.Encode())
}
