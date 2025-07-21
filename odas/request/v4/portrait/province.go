package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type ProvinceOptions struct {
	Limit   int  `json:"limit"`
	Unknown bool `json:"unknown"`
}

type ProvinceOption func(options *ProvinceOptions)

func WithProvinceLimit(limit int) ProvinceOption {
	return func(options *ProvinceOptions) {
		options.Limit = limit
	}
}

func WithProvinceUnknown(unknown bool) ProvinceOption {
	return func(options *ProvinceOptions) {
		options.Unknown = unknown
	}
}

// ProvinceReq 省客源排行
type ProvinceReq struct {
	odas.Req
	odas.DateRangeCompareReq
	Options *ProvinceOptions
}

func NewProvinceReq(
	req *odas.Req,
	dateRangeCompareReq *odas.DateRangeCompareReq,
	opt ...ProvinceOption,
) *ProvinceReq {
	options := &ProvinceOptions{}
	for _, p := range opt {
		p(options)
	}
	return &ProvinceReq{
		Req:                 *req,
		DateRangeCompareReq: *dateRangeCompareReq,
		Options:             options,
	}
}

func (r ProvinceReq) Api() string {
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
	return fmt.Sprintf("/v4/portrait/province?%s", params.Encode())
}

type ProvinceRankResponse struct {
	Province         string  `json:"province"`
	Total            int     `json:"total"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Rate             float64 `json:"rate"`
}
