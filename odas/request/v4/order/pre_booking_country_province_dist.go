package order

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type PreBookingCountryProvinceDistReq struct {
	odas.Req
}

func (p PreBookingCountryProvinceDistReq) Api() string {
	params := p.Req.Api()

	return fmt.Sprintf("/v4/order/preBookingCountryProvinceDist?%s", params.Encode())
}

func NewPreBookingCountryProvinceDistReq(req odas.Req) *PreBookingCountryProvinceDistReq {
	return &PreBookingCountryProvinceDistReq{
		Req: req,
	}
}

type PreBookingCountryProvinceDistCountListItem struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	Count    int    `json:"count"`
}

type PreBookingCountryProvinceDistListItem struct {
	CountList []*PreBookingCountryProvinceDistCountListItem `json:"countList"`
	Time      int                                           `json:"time"`
}

type PreBookingCountryProvinceDistResponse struct {
	List []*PreBookingCountryProvinceDistListItem `json:"list"`
}
