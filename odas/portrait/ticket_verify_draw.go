package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type SexAgeV2Options struct {
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
}

type SexAgeV2Option func(options *SexAgeV2Options)

// SexAgeV2SummaryReq 性别年龄分布
type SexAgeV2SummaryReq struct {
	odas.Req
	Sid int
}

func NewSexAgeV2SummaryReq(req *odas.Req, opt ...SexAgeV2Option) *SexAgeV2SummaryReq {
	options := &SexAgeV2Options{}
	for _, p := range opt {
		p(options)
	}
	return &SexAgeV2SummaryReq{
		Req: *req,
	}
}

func (r SexAgeV2SummaryReq) Api() string {
	params := r.Req.Params()

	return fmt.Sprintf("/v4/portrait/ticket/ageSummary?%s", params.Encode())
}
