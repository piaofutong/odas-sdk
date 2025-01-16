package portrait

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type TicketPaymentMethodOptions struct {
	Province string `json:"province"`
	Limit    int    `json:"limit"`
}

type TicketPaymentMethodOption func(options *TicketPaymentMethodOptions)

func WithTicketPaymentMethodLimit(limit int) TicketPaymentMethodOption {
	return func(options *TicketPaymentMethodOptions) {
		options.Limit = limit
	}
}

func WithTicketPaymentMethodProvince(province string) TicketPaymentMethodOption {
	return func(options *TicketPaymentMethodOptions) {
		options.Province = province
	}
}

// TicketPaymentMethodReq 支付渠道
type TicketPaymentMethodReq struct {
	odas.Req
	Options *TicketPaymentMethodOptions
}

func NewTicketPaymentMethodReq(req *odas.Req, opt ...TicketPaymentMethodOption) *TicketPaymentMethodReq {
	options := &TicketPaymentMethodOptions{}
	for _, p := range opt {
		p(options)
	}
	return &TicketPaymentMethodReq{
		Req:     *req,
		Options: options,
	}
}

func (r TicketPaymentMethodReq) Api() string {
	params := r.Req.Params()
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/ticket/paymentMethod?%s", params.Encode())
}

type TicketPaymentMethodResponse struct {
	Total int                            `json:"total"`
	List  []*TicketPaymentMethodListItem `json:"list"`
}

type TicketPaymentMethodListItem struct {
	Channel string  `json:"name"`
	Total   int     `json:"count"`
	Rate    float64 `json:"rate"`
}
