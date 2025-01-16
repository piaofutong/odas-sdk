package tourist

import (
	"fmt"
	"strconv"

	"github.com/piaofutong/odas-sdk/odas"
)

type LocalV2Options struct {
	Limit   int    `json:"limit"`
	Unknown bool   `json:"unknown"`
	GroupBy string `json:"groupBy"`
}

type LocalV2Option func(options *LocalV2Options)

// LocalV2Req 客流来源TopN
type LocalV2Req struct {
	odas.Req
	Options  *LocalV2Options
	Province string
	City     string
}

func WithGroupBy(groupBy string) LocalV2Option {
	return func(options *LocalV2Options) {
		options.GroupBy = groupBy
	}
}
func NewLocalV2Req(req *odas.Req, province string, city string, opt ...LocalV2Option) *LocalV2Req {
	options := &LocalV2Options{}
	for _, p := range opt {
		p(options)
	}

	return &LocalV2Req{
		Req:      *req,
		Options:  options,
		Province: province,
		City:     city,
	}
}

func (l LocalV2Req) Api() string {
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
	if l.Options.GroupBy != "" {
		params.Add("groupBy", l.Options.GroupBy)
	}
	return fmt.Sprintf("/v4/tourist/touristLocalV2?%s", params.Encode())
}

type LocalV2Response struct {
	Total   *LocalV2Total                 `json:"total"`
	Inside  []*LocalV2InsideProvinceList  `json:"inside"`
	Outside []*LocalV2OutsideProvinceList `json:"outside"`
}

type LocalV2Total struct {
	InsideProvince  int     `json:"insideProvince"`
	InsideRate      float64 `json:"insideRate"`
	OutsideProvince int     `json:"outsideProvince"`
	OutsideRate     float64 `json:"outsideRate"`
}

type LocalV2InsideProvinceList struct {
	City  string  `json:"city"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}

type LocalV2OutsideProvinceList struct {
	Province string  `json:"province"`
	Count    int     `json:"count"`
	Rate     float64 `json:"rate"`
}

type LocalV2ResponseListItem struct {
	ProvinceName     string  `json:"provinceName"`
	CityName         string  `json:"cityName"`
	DistrictName     string  `json:"districtName"`
	Count            int     `json:"count"`
	CompareTotalRate float64 `json:"compareTotalRate"`
	Percent          float64 `json:"percent"`
}
