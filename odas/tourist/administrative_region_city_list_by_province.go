package tourist

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// AdministrativeRegionCityListByProvinceReq 客流来源TopN
type AdministrativeRegionCityListByProvinceReq struct {
	odas.Req
	Province string `form:"province" json:"province" binding:"omitempty"`
}

func NewAdministrativeRegionCityListByProvinceReq(req *odas.Req, province string) *AdministrativeRegionCityListByProvinceReq {
	return &AdministrativeRegionCityListByProvinceReq{
		Req:      *req,
		Province: province,
	}
}

func (l AdministrativeRegionCityListByProvinceReq) Api() string {
	params := l.Req.Params()
	if l.Province != "" {
		params.Add("province", l.Province)
	}
	return fmt.Sprintf("/v4/tourist/cityListByProvince?%s", params.Encode())
}

type CityListByProvinceResponse struct {
	List []*CityListByProvinceItem `json:"list"`
}

type CityListByProvinceItem struct {
	Name string `json:"name"`
}
