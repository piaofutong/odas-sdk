package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// LocalReq 客流来源TopN
type LocalReq struct {
	odas.Req
	Province string `json:"province"`
	Limit    int    `json:"limit"`
	Unknown  bool   `json:"unknown"`
}

func NewLocalReq(req *odas.Req, province string, limit int, unknown bool) *LocalReq {
	return &LocalReq{
		Req:      *req,
		Province: province,
		Limit:    limit,
		Unknown:  unknown,
	}
}

func (l LocalReq) Api() string {
	params := l.Req.Params()
	if l.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Limit))
	}
	if l.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Unknown))
	}
	ps := fmt.Sprintf("/v4/tourist/touristLocal?%s", params.Encode())
	if l.Province != "" {
		ps += fmt.Sprintf("&province=%s", l.Province)
	}
	return ps
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
