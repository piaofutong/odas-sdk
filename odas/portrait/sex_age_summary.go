package portrait

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// SexAgeSummaryReq 性别年龄分布
type SexAgeSummaryReq struct {
	odas.Req
	Province string `json:"province"`
	Unknown  bool   `json:"unknown"`
}

func NewSexAgeSummaryReq(req *odas.Req, province string, unknown bool) *SexAgeSummaryReq {
	return &SexAgeSummaryReq{
		Req:      *req,
		Province: province,
		Unknown:  unknown,
	}
}

func (r SexAgeSummaryReq) Api() string {
	params := r.Req.Params()
	if r.Unknown {
		params.Add("unknown", strconv.FormatBool(r.Unknown))
	}
	if r.Province != "" {
		params.Add("province", r.Province)
	}
	return fmt.Sprintf("/v4/portrait/ageSummary?%s", params.Encode())
}

type AgeSummaryResponse struct {
	Total *AgeSummaryTotal  `json:"total"`
	List  []*AgeSummaryList `json:"list"`
}

type AgeSummaryTotal struct {
	Total  int `json:"total"`
	Male   int `json:"male"`
	Female int `json:"female"`
}

type AgeSummaryList struct {
	AgeGroup string `json:"ageGroup"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}
