package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// LocalReq 客流来源TopN
type LocalReq struct {
	odas.Req
	Options *LocalOptions
}

func NewLocalReq(req *odas.Req, opt ...LocalOption) *LocalReq {
	options := &LocalOptions{}
	for _, p := range opt {
		p(options)
	}

	return &LocalReq{
		Req:     *req,
		Options: options,
	}
}

func (l LocalReq) Api() string {
	params := l.Req.Params()
	if l.Options.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Options.Limit))
	}
	if l.Options.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Options.Unknown))
	}
	if l.Options.Province != "" {
		params.Add("province", l.Options.Province)
	}
	return fmt.Sprintf("/v4/tourist/touristLocal?%s", params.Encode())
}

type LocalResponse struct {
	Total   *LocalTotal                 `json:"total"`
	Inside  []*LocalInsideProvinceList  `json:"inside"`
	Outside []*LocalOutsideProvinceList `json:"outside"`
}

type LocalTotal struct {
	InsideProvince  int     `json:"insideProvince"`
	InsideRate      float64 `json:"insideRate"`
	OutsideProvince int     `json:"outsideProvince"`
	OutsideRate     float64 `json:"outsideRate"`
}

type LocalInsideProvinceList struct {
	City  string  `json:"city"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}

type LocalOutsideProvinceList struct {
	Province string  `json:"province"`
	Count    int     `json:"count"`
	Rate     float64 `json:"rate"`
}
