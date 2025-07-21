package sixun

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type SaleShopTOpNReqOptions struct {
	Limit int `json:"limit"`
}

type SaleShopTOpNReqOption func(options *SaleShopTOpNReqOptions)

func WithSaleShopTOpNReqLimit(limit int) SaleShopTOpNReqOption {
	return func(options *SaleShopTOpNReqOptions) {
		options.Limit = limit
	}
}

// SaleShopTopNReq 商户TOPN
type SaleShopTopNReq struct {
	odas.Req
	options *SaleShopTOpNReqOptions
}

func NewSaleShopTopNReq(req *odas.Req, opt ...SaleShopTOpNReqOption) *SaleShopTopNReq {
	options := &SaleShopTOpNReqOptions{}
	for _, p := range opt {
		p(options)
	}

	return &SaleShopTopNReq{
		Req:     *req,
		options: options,
	}
}

func (h SaleShopTopNReq) Api() string {
	params := h.Req.Params()

	if h.options.Limit > 0 {
		params.Add("limit", strconv.Itoa(h.options.Limit))
	}

	return fmt.Sprintf("/v4/sixun/saleShopTopN?%s", params.Encode())
}

type SaleShopTopNListItem struct {
	ShopName        string `json:"shopName"`
	ShopCategory    string `json:"shopCategory"`
	SaleQuntity     int    `json:"saleQuntity"`
	SettlementMoney int    `json:"settlementMoney"`
}

type SaleShopTopNResponse struct {
	List []*SaleShopTopNListItem `json:"list"`
}
