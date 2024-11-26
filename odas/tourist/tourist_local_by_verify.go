package tourist

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
	"strconv"
)

// LocalByVerifyReq 客流来源TopN(验证维度)
type LocalByVerifyReq struct {
	odas.Req
	Province string `json:"province"`
	Limit    int    `json:"limit"`
	Unknown  bool   `json:"unknown"`
}

func NewLocalByVerifyReq(req *odas.Req, province string, limit int, unknown bool) *LocalByVerifyReq {
	return &LocalByVerifyReq{
		Req:      *req,
		Province: province,
		Limit:    limit,
		Unknown:  unknown,
	}
}

func (l LocalByVerifyReq) Api() string {
	params := l.Req.Params()
	if l.Limit > 0 {
		params.Add("limit", strconv.Itoa(l.Limit))
	}
	if l.Unknown {
		params.Add("unknown", strconv.FormatBool(l.Unknown))
	}
	if l.Province != "" {
		params.Add("province", l.Province)
	}
	return fmt.Sprintf("/v4/tourist/touristLocalByVerify?%s", params.Encode())
}
