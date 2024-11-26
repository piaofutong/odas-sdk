package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// SexAgeSummaryByVerifyReq 性别年龄分布(验证维度)
type SexAgeSummaryByVerifyReq struct {
	odas.Req
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
}

func NewSexAgeSummaryByVerifyReq(req *odas.Req, province string, unknown bool) *SexAgeSummaryByVerifyReq {
	return &SexAgeSummaryByVerifyReq{
		Req:      *req,
		Province: province,
		Unknown:  unknown,
	}
}

func (r SexAgeSummaryByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	if r.Province != "" {
		params.Add("province", r.Province)
	}
	return fmt.Sprintf("/v4/portrait/ageSummaryByVerify?%s", params.Encode())
}
