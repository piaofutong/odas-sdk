package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// SexAgeSummaryByVerifyReq 性别年龄分布(验证维度)
type SexAgeSummaryByVerifyReq struct {
	odas.Req
	Options *SexAgeOptions
}

func NewSexAgeSummaryByVerifyReq(req *odas.Req, opt ...SexAgeOption) *SexAgeSummaryByVerifyReq {
	options := &SexAgeOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SexAgeSummaryByVerifyReq{
		Req:     *req,
		Options: options,
	}
}

func (r SexAgeSummaryByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Options.Unknown))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/ageSummaryByVerify?%s", params.Encode())
}
