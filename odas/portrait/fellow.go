package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

// FellowReq 同行人数
type FellowReq struct {
	odas.Req
	Province string `json:"province"`
}

func NewFellowReq(req *odas.Req, province string) *FellowReq {
	return &FellowReq{
		Req:      *req,
		Province: province,
	}
}

func (r *FellowReq) Api() string {
	ps := "/v4/portrait/fellow?" + r.Req.Params().Encode()
	if r.Province != "" {
		ps += fmt.Sprintf("&province=%s", r.Province)
	}
	return ps
}

type FellowResponse struct {
	Total int           `json:"total"`
	List  []*FellowList `json:"list"`
}

type FellowList struct {
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Rate  float64 `json:"rate"`
}
