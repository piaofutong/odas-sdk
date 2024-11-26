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
	Limit   int  `json:"limit"`
	Unknown bool `json:"unknown"`
}

func NewProvinceByVerifyReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	limit int,
	unknown bool,
) *ProvinceByVerifyReq {
	return &ProvinceByVerifyReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Limit:               limit,
		Unknown:             unknown,
	}
}

func (r ProvinceByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	if r.CompareStart != "" {
		params.Add("compareStart", r.CompareStart)
	}
	if r.CompareEnd != "" {
		params.Add("compareEnd", r.CompareEnd)
	}
	return fmt.Sprintf("/v4/portrait/provinceByVerify?%s", params.Encode())
}
