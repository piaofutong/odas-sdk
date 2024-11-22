package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// CityReq 市客源排行
type CityReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
	Limit    int    `json:"limit"`
}

func NewCityReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	province string,
	unknown bool,
	limit int,
) *CityReq {
	return &CityReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Province:            province,
		Unknown:             unknown,
		Limit:               limit,
	}
}

func (r CityReq) Api() string {
	params := r.Req.Params()
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	if r.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.CompareStart != "" {
		params.Add("compareStart", r.CompareStart)
	}
	if r.CompareEnd != "" {
		params.Add("compareEnd", r.CompareEnd)
	}

	if r.Province != "" {
		params.Add("province", r.Province)
	}
	return fmt.Sprintf("/v4/portrait/city?%s", params.Encode())
}

type CityRankResponse struct {
	City             string  `json:"city"`
	Total            int     `json:"total"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Rate             float64 `json:"rate"`
}
