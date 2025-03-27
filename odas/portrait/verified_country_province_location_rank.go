package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// VerifiedCountryProvinceLocationRankReq 客流来源TopN-国家、大陆分组
type VerifiedCountryProvinceLocationRankReq struct {
	odas.Req
	Province string `json:"province"`
	Limit    int    `json:"limit"`
	Unknown  bool   `json:"unknown"`
}

func NewVerifiedCountryProvinceLocationRankReq(req *odas.Req) *VerifiedCountryProvinceLocationRankReq {
	return &VerifiedCountryProvinceLocationRankReq{
		Req: *req,
	}
}

func (l VerifiedCountryProvinceLocationRankReq) Api() string {
	params := l.Req.Params()
	return fmt.Sprintf("/v4/portrait/verifiedCountryProvinceLocationRank?%s", params.Encode())
}
