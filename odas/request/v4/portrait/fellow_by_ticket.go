package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type FellowByTicketOptions struct {
	Province string
}

type FellowByTicketOption func(options *FellowByTicketOptions)

func WithFellowByTicketProvince(province string) FellowByTicketOption {
	return func(options *FellowByTicketOptions) {
		options.Province = province
	}
}

// FellowByTicketReq 同行人数
type FellowByTicketReq struct {
	odas.Req
	Options *FellowByTicketOptions
}

func NewFellowByTicketReq(req *odas.Req, opt ...FellowByTicketOption) *FellowByTicketReq {
	options := &FellowByTicketOptions{}
	for _, p := range opt {
		p(options)
	}
	return &FellowByTicketReq{
		Req:     *req,
		Options: options,
	}
}

func (r *FellowByTicketReq) Api() string {
	params := r.Req.Params()
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/fellowByTicket?%s", params.Encode())
}

type FellowByTicketResponse struct {
	Total int                   `json:"total"`
	List  []*FellowByTicketList `json:"list"`
}

type FellowByTicketList struct {
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}
