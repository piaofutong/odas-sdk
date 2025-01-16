package tourist

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type AdministrativeRegionProvinceListOptions struct {
	Province string `json:"province"`
	Limit    int    `json:"limit"`
	Unknown  bool   `json:"unknown"`
	GroupBy  string `json:"groupBy"`
}

type AdministrativeRegionProvinceListOption func(options *AdministrativeRegionProvinceListOptions)

// AdministrativeRegionProvinceListReq 客流来源TopN
type AdministrativeRegionProvinceListReq struct {
	odas.Req
	Options *AdministrativeRegionProvinceListOptions
}

func NewAdministrativeRegionProvinceListReq(req *odas.Req, opt ...AdministrativeRegionProvinceListOption) *AdministrativeRegionProvinceListReq {
	options := &AdministrativeRegionProvinceListOptions{}
	for _, p := range opt {
		p(options)
	}

	return &AdministrativeRegionProvinceListReq{
		Req:     *req,
		Options: options,
	}
}

func (l AdministrativeRegionProvinceListReq) Api() string {
	params := l.Req.Params()
	return fmt.Sprintf("/v4/tourist/provinceList?%s", params.Encode())
}

type ProvinceListResponse struct {
	List []*ProvinceListItem `json:"list"`
}

type ProvinceListItem struct {
	Name string `json:"name"`
}
