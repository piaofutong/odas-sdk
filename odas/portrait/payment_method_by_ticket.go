package portrait

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type PaymentMethodByTicketOptions struct {
	Province string `json:"province"`
	Limit    int    `json:"limit"`
}

type PaymentMethodByTicketOption func(options *PaymentMethodByTicketOptions)

func WithPaymentMethodByTicketLimit(limit int) PaymentMethodByTicketOption {
	return func(options *PaymentMethodByTicketOptions) {
		options.Limit = limit
	}
}

func WithPaymentMethodByTicketProvince(province string) PaymentMethodByTicketOption {
	return func(options *PaymentMethodByTicketOptions) {
		options.Province = province
	}
}

// PaymentMethodByTicketReq 支付渠道
type PaymentMethodByTicketReq struct {
	odas.Req
	Options *PaymentMethodByTicketOptions
}

func NewPaymentMethodByTicketReq(req *odas.Req, opt ...PaymentMethodByTicketOption) *PaymentMethodByTicketReq {
	options := &PaymentMethodByTicketOptions{}
	for _, p := range opt {
		p(options)
	}
	return &PaymentMethodByTicketReq{
		Req:     *req,
		Options: options,
	}
}

func (r PaymentMethodByTicketReq) Api() string {
	params := r.Req.Params()
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/paymentMethodByTicket?%s", params.Encode())
}

type PaymentMethodByTicketResponse struct {
	Total int                              `json:"total"`
	List  []*PaymentMethodByTicketListItem `json:"list"`
}

type PaymentMethodByTicketListItem struct {
	Channel string  `json:"name"`
	Total   int     `json:"count"`
	Rate    float64 `json:"rate"`
}
