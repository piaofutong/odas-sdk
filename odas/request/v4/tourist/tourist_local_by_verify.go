package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type LocalOptions struct {
	Province string `json:"province"`
	Limit    int    `json:"limit"`
	Unknown  bool   `json:"unknown"`
}

type LocalOption func(options *LocalOptions)

func WithLocalUnknown(unknown bool) LocalOption {
	return func(options *LocalOptions) {
		options.Unknown = unknown
	}
}

func WithLocalLimit(limit int) LocalOption {
	return func(options *LocalOptions) {
		options.Limit = limit
	}
}

func WithLocalProvince(province string) LocalOption {
	return func(options *LocalOptions) {
		options.Province = province
	}
}

// LocalByVerifyReq 客流来源TopN(验证维度)
type LocalByVerifyReq struct {
	odas.Req
	Options *LocalOptions
}

func NewLocalByVerifyReq(req *odas.Req, opt ...LocalOption) *LocalByVerifyReq {
	options := &LocalOptions{}
	for _, p := range opt {
		p(options)
	}
	return &LocalByVerifyReq{
		Req:     *req,
		Options: options,
	}
}

func (l LocalByVerifyReq) Api() string {
	params := l.Req.Params()
	if l.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Options.Limit))
	}
	if l.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Options.Unknown))
	}
	if l.Options.Province != "" {
		params.Add("province", l.Options.Province)
	}
	return fmt.Sprintf("/v4/tourist/touristLocalByVerify?%s", params.Encode())
}
