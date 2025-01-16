package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type TicketFellowOptions struct {
	Province string
}

type TicketFellowOption func(options *TicketFellowOptions)

func WithTicketFellowProvince(province string) TicketFellowOption {
	return func(options *TicketFellowOptions) {
		options.Province = province
	}
}

// TicketFellowReq 同行人数
type TicketFellowReq struct {
	odas.Req
	Options *TicketFellowOptions
}

func NewTicketFellowReq(req *odas.Req, opt ...TicketFellowOption) *TicketFellowReq {
	options := &TicketFellowOptions{}
	for _, p := range opt {
		p(options)
	}
	return &TicketFellowReq{
		Req:     *req,
		Options: options,
	}
}

func (r *TicketFellowReq) Api() string {
	params := r.Req.Params()
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/ticket/fellow?%s", params.Encode())
}

type TicketFellowResponse struct {
	Total int                 `json:"total"`
	List  []*TicketFellowList `json:"list"`
}

type TicketFellowList struct {
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}
