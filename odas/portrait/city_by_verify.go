package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// CityByVerifyReq 市客源排行(验证维度)
type CityByVerifyReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
	Limit    int    `json:"limit"`
}

func NewCityByVerifyReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	province string,
	unknown bool,
	limit int,
) *CityByVerifyReq {
	return &CityByVerifyReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Province:            province,
		Unknown:             unknown,
		Limit:               limit,
	}
}

func (r CityByVerifyReq) Api() string {
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
	return fmt.Sprintf("/v4/portrait/cityByVerify?%s", params.Encode())
}
