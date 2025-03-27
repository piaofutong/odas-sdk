package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// BookingCountryProvinceLocationRankReq 客流来源TopN-国家、大陆分组
type BookingCountryProvinceLocationRankReq struct {
	odas.Req
	Limit   int  `json:"limit"`
	Unknown bool `json:"unknown"`
}

func NewBookingCountryProvinceLocationRankReq(req *odas.Req) *BookingCountryProvinceLocationRankReq {
	return &BookingCountryProvinceLocationRankReq{
		Req: *req,
	}
}

func (l BookingCountryProvinceLocationRankReq) Api() string {
	params := l.Req.Params()
	return fmt.Sprintf("/v4/portrait/bookingCountryProvinceLocationRank?%s", params.Encode())
}

type CountryProvinceLocationRankItem struct {
	Country      string `json:"country"`
	ProvinceName string `json:"provinceName"`
	Count        int    `json:"count"`
}

type CountryProvinceLocationRankResponse struct {
	List []*CountryProvinceLocationRankItem `json:"list"`
}
