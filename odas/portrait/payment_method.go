package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// PaymentMethodReq 支付渠道
type PaymentMethodReq struct {
	odas.Req
	Province string `json:"province"`
	Limit    int    `json:"limit"`
}

func NewPaymentMethodReq(req *odas.Req, province string, limit int) *PaymentMethodReq {
	return &PaymentMethodReq{
		Req:      *req,
		Province: province,
		Limit:    limit,
	}
}

func (r PaymentMethodReq) Api() string {
	params := r.Req.Params()
	if r.Limit > 0 {
		params.Add("limit", strconv.Itoa(r.Limit))
	}
	ps := fmt.Sprintf("/v4/portrait/paymentMethod?%s", params.Encode())
	if r.Province != "" {
		ps += fmt.Sprintf("&province=%s", r.Province)
	}
	return ps
}

type PaymentMethodResponse struct {
	Channel string  `json:"name"`
	Total   int     `json:"count"`
	Rate    float64 `json:"rate"`
}
