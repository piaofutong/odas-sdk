package sixun

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type SaleProductTopNReqOptions struct {
	Limit int `json:"limit"`
}

type SaleProductTopNReqOption func(options *SaleProductTopNReqOptions)

func WithSaleProductTopNReqLimit(limit int) SaleProductTopNReqOption {
	return func(options *SaleProductTopNReqOptions) {
		options.Limit = limit
	}
}

// SaleProductTopNReq 商品TOPN
type SaleProductTopNReq struct {
	odas.Req
	options *SaleProductTopNReqOptions
}

func NewSaleProductTopNReq(req *odas.Req, opt ...SaleProductTopNReqOption) *SaleProductTopNReq {
	options := &SaleProductTopNReqOptions{}
	for _, p := range opt {
		p(options)
	}
	return &SaleProductTopNReq{
		Req:     *req,
		options: options,
	}
}

func (h SaleProductTopNReq) Api() string {
	params := h.Req.Params()

	if h.options.Limit > 0 {
		params.Add("limit", strconv.Itoa(h.options.Limit))
	}
	return fmt.Sprintf("/v4/sixun/saleProductTopN?%s", params.Encode())
}

type SaleProductTopNListItem struct {
	ProductName string `json:"productName"`
	SaleQuntity int    `json:"saleQuntity"`
	SaleMoney   int    `json:"saleMoney"`
}

type SaleProductTopNResponse struct {
	List []*SaleProductTopNListItem `json:"list"`
}
