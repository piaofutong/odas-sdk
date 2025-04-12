package portrait

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

// VerifiedCountryProvinceLocationRankReq 客流来源TopN-国家、大陆分组
type VerifiedCountryProvinceLocationRankReq struct {
	odas.Req
	Options *VerifiedCountryProvinceLocationRankOptions
}

type VerifiedCountryProvinceLocationRankOptions struct {
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
	Limit    int    `json:"limit"`
}

type VerifiedCountryProvinceLocationRankOption func(options *VerifiedCountryProvinceLocationRankOptions)

func WithVerifiedCountryProvinceLocationRankLimit(limit int) VerifiedCountryProvinceLocationRankOption {
	return func(options *VerifiedCountryProvinceLocationRankOptions) {
		options.Limit = limit
	}
}

func WithVerifiedCountryProvinceLocationRankUnknown(unknown bool) VerifiedCountryProvinceLocationRankOption {
	return func(options *VerifiedCountryProvinceLocationRankOptions) {
		options.Unknown = unknown
	}
}

func WithVerifiedCountryProvinceLocationRankProvince(province string) VerifiedCountryProvinceLocationRankOption {
	return func(options *VerifiedCountryProvinceLocationRankOptions) {
		options.Province = province
	}
}

func NewVerifiedCountryProvinceLocationRankReq(req *odas.Req, opt ...VerifiedCountryProvinceLocationRankOption) *VerifiedCountryProvinceLocationRankReq {
	options := &VerifiedCountryProvinceLocationRankOptions{}
	for _, p := range opt {
		p(options)
	}

	return &VerifiedCountryProvinceLocationRankReq{
		Req:     *req,
		Options: options,
	}
}

func (l VerifiedCountryProvinceLocationRankReq) Api() string {
	params := l.Req.Params()
	if l.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Options.Unknown))
	}
	if l.Options.Province != "" {
		params.Add("province", l.Options.Province)
	}
	if l.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Options.Limit))
	}

	return fmt.Sprintf("/v4/portrait/verifiedCountryProvinceLocationRank?%s", params.Encode())
}
