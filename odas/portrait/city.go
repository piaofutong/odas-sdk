package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type CityOptions struct {
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
	Limit    int    `json:"limit"`
}

type CityOption func(options *CityOptions)

func WithCityLimit(limit int) CityOption {
	return func(options *CityOptions) {
		options.Limit = limit
	}
}

func WithCityUnknown(unknown bool) CityOption {
	return func(options *CityOptions) {
		options.Unknown = unknown
	}
}

func WithCityProvince(province string) CityOption {
	return func(options *CityOptions) {
		options.Province = province
	}
}

// CityReq 市客源排行
type CityReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Options *CityOptions
}

func NewCityReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	opt ...CityOption,
) *CityReq {
	options := &CityOptions{}
	for _, p := range opt {
		p(options)
	}
	return &CityReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Options:             options,
	}
}

func (r CityReq) Api() string {
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
	return fmt.Sprintf("/v4/portrait/city?%s", params.Encode())
}

type CityRankResponse struct {
	City             string  `json:"city"`
	Total            int     `json:"total"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Rate             float64 `json:"rate"`
}
