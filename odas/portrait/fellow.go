package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type FellowOptions struct {
	Province string
}

type FellowOption func(options *FellowOptions)

func WithFellowProvince(province string) FellowOption {
	return func(options *FellowOptions) {
		options.Province = province
	}
}

// FellowReq 同行人数
type FellowReq struct {
	odas.Req
	Options *FellowOptions
}

func NewFellowReq(req *odas.Req, opt ...FellowOption) *FellowReq {
	options := &FellowOptions{}
	for _, p := range opt {
		p(options)
	}
	return &FellowReq{
		Req:     *req,
		Options: options,
	}
}

func (r *FellowReq) Api() string {
	params := r.Req.Params()
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/fellow?%s", params.Encode())
}

type FellowResponse struct {
	Total int           `json:"total"`
	List  []*FellowList `json:"list"`
}

type FellowList struct {
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}
