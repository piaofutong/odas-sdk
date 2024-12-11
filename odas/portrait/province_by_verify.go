package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// ProvinceByVerifyReq 省客源排行
type ProvinceByVerifyReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Options *ProvinceOptions
}

func NewProvinceByVerifyReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	opt ...ProvinceOption,
) *ProvinceByVerifyReq {
	options := &ProvinceOptions{}
	for _, p := range opt {
		p(options)
	}
	return &ProvinceByVerifyReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Options:             options,
	}
}

func (r ProvinceByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	if r.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Options.Unknown))
	}
	if r.CompareStart != "" {
		params.Add("compareStart", r.CompareStart)
	}
	if r.CompareEnd != "" {
		params.Add("compareEnd", r.CompareEnd)
	}
	return fmt.Sprintf("/v4/portrait/provinceByVerify?%s", params.Encode())
}
