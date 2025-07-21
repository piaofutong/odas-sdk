package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

type PayMethodOptions struct {
	Province string `json:"province"`
	Limit    int    `json:"limit"`
}

type PaymentMethodOption func(options *PayMethodOptions)

func WithPaymentMethodLimit(limit int) PaymentMethodOption {
	return func(options *PayMethodOptions) {
		options.Limit = limit
	}
}

func WithPaymentMethodProvince(province string) PaymentMethodOption {
	return func(options *PayMethodOptions) {
		options.Province = province
	}
}

// PaymentMethodReq 支付渠道
type PaymentMethodReq struct {
	odas.Req
	Options *PayMethodOptions
}

func NewPaymentMethodReq(req *odas.Req, opt ...PaymentMethodOption) *PaymentMethodReq {
	options := &PayMethodOptions{}
	for _, p := range opt {
		p(options)
	}
	return &PaymentMethodReq{
		Req:     *req,
		Options: options,
	}
}

func (r PaymentMethodReq) Api() string {
	params := r.Req.Params()
	if r.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Options.Limit))
	}
	if r.Options.Province != "" {
		params.Add("province", r.Options.Province)
	}
	return fmt.Sprintf("/v4/portrait/paymentMethod?%s", params.Encode())
}

type PaymentMethodResponse struct {
	Channel string  `json:"name"`
	Total   int     `json:"count"`
	Rate    float64 `json:"rate"`
}
