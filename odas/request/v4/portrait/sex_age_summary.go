package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type SexAgeOptions struct {
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
}

type SexAgeOption func(options *SexAgeOptions)

func WithSexAgeUnknown(unknown bool) SexAgeOption {
	return func(options *SexAgeOptions) {
		options.Unknown = unknown
	}
}

func WithSexAgeProvince(province string) SexAgeOption {
	return func(options *SexAgeOptions) {
		options.Province = province
	}
}

// SexAgeSummaryReq 性别年龄分布
type SexAgeSummaryReq struct {
	odas.Req
	Options *SexAgeOptions
}

func NewSexAgeSummaryReq(req *odas.Req, opt ...SexAgeOption) *SexAgeSummaryReq {
	options := &SexAgeOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SexAgeSummaryReq{
		Req:     *req,
		Options: options,
	}
}

func (r SexAgeSummaryReq) Api() string {
	params := r.Req.Params()
	if r.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Options.Unknown))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/ageSummary?%s", params.Encode())
}

type AgeSummaryResponse struct {
	Total *AgeSummaryTotal  `json:"total"`
	List  []*AgeSummaryList `json:"list"`
}

type AgeSummaryTotal struct {
	Total  int `json:"total"`
	Male   int `json:"male"`
	Female int `json:"female"`
}

type AgeSummaryList struct {
	AgeGroup string `json:"ageGroup"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}
