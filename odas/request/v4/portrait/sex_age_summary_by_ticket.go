package portrait

import (
	"fmt"
	"strconv"

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
	Sid     int
	Options *SexAgeByTicketOptions
}

func NewSexAgeSummaryByTicketReq(req *odas.Req, opt ...SexAgeByTicketOption) *SexAgeSummaryByTicketReq {
	options := &SexAgeByTicketOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SexAgeSummaryByTicketReq{
		Req:     *req,
		Options: options,
	}
}

func (r SexAgeSummaryByTicketReq) Api() string {
	params := r.Req.Params()
	if r.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Options.Unknown))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/ageSummaryByTicket?%s", params.Encode())
}
