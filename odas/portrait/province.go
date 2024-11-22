package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// ProvinceReq 省客源排行
type ProvinceReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Limit   int  `json:"limit"`
	Unknown bool `json:"unknown"`
}

func NewProvinceReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	limit int,
	unknown bool,
) *ProvinceReq {
	return &ProvinceReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Limit:               limit,
		Unknown:             unknown,
	}
}

func (r ProvinceReq) Api() string {
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
	return fmt.Sprintf("/v4/portrait/province?%s", params.Encode())
}

type ProvinceRankResponse struct {
	Province         string  `json:"province"`
	Total            int     `json:"total"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Rate             float64 `json:"rate"`
}
