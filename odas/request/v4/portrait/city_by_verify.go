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
	Options *CityOptions
}

func NewCityByVerifyReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	opt ...CityOption,
) *CityByVerifyReq {
	options := &CityOptions{}
	for _, p := range opt {
		p(options)
	}
	return &CityByVerifyReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Options:             options,
	}
}

func (r CityByVerifyReq) Api() string {
	params := r.Req.Params()
	if r.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Options.Unknown))
	}
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	if r.CompareStart != "" {
		params.Add("compareStart", r.CompareStart)
	}
	if r.CompareEnd != "" {
		params.Add("compareEnd", r.CompareEnd)
	}

	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/cityByVerify?%s", params.Encode())
}
