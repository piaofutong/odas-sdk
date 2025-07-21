package tourist

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type LocalByTicketOptions struct {
	Limit      int    `json:"limit"`
	Unknown    bool   `json:"unknown"`
	RegionType string `json:"regionType"`
}

type LocalByTicketOption func(options *LocalByTicketOptions)

// LocalByTicketReq 客流来源TopN
type LocalByTicketReq struct {
	*odas.Req
	*odas.DateRangeCompareReq
	Options  *LocalByTicketOptions
	Province string
	City     string
}

func WithRegionType(regionType string) LocalByTicketOption {
	return func(options *LocalByTicketOptions) {
		options.RegionType = regionType
	}
}
func WithLocalByTicketLimit(limit int) LocalByTicketOption {
	return func(options *LocalByTicketOptions) {
		options.Limit = limit
	}
}
func NewLocalByTicketReq(req *odas.Req, compareDateReq *odas.DateRangeCompareReq, province string, city string, opt ...LocalByTicketOption) *LocalByTicketReq {
	options := &LocalByTicketOptions{}
	for _, p := range opt {
		p(options)
	}

	return &LocalByTicketReq{
		Req:                 req,
		DateRangeCompareReq: compareDateReq,
		Options:             options,
		Province:            province,
		City:                city,
	}
}

func (l LocalByTicketReq) Api() string {
	params := l.Req.Params()
	if l.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Options.Limit))
	}
	if l.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Options.Unknown))
	}
	if l.Province != "" {
		params.Add("province", l.Province)
	}
	if l.City != "" {
		params.Add("city", l.City)
	}
	if l.Options.RegionType != "" {
		params.Add("regionType", l.Options.RegionType)
	}
	if l.CompareStart != "" {
		params.Add("compareStart", l.CompareStart)
	}
	if l.CompareEnd != "" {
		params.Add("compareEnd", l.CompareEnd)
	}
	return fmt.Sprintf("/v4/tourist/touristLocalByTicket?%s", params.Encode())
}

type LocalByTicketResponse struct {
	Total   *LocalByTicketTotal                 `json:"total"`
	Inside  []*LocalByTicketInsideProvinceList  `json:"inside"`
	Outside []*LocalByTicketOutsideProvinceList `json:"outside"`
}

type LocalByTicketTotal struct {
	InsideProvince  int     `json:"insideProvince"`
	InsideRate      float64 `json:"insideRate"`
	OutsideProvince int     `json:"outsideProvince"`
	OutsideRate     float64 `json:"outsideRate"`
}

type LocalByTicketInsideProvinceList struct {
	City  string  `json:"city"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}

type LocalByTicketOutsideProvinceList struct {
	Province string  `json:"province"`
	Count    int     `json:"count"`
	Rate     float64 `json:"rate"`
}

type LocalByTicketResponseListItem struct {
	ProvinceName     string  `json:"provinceName"`
	CityName         string  `json:"cityName"`
	DistrictName     string  `json:"districtName"`
	Count            int     `json:"count"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Percent          float64 `json:"percent"`
}
