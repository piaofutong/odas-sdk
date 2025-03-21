package portrait

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

// AbsenceBlackListReq 查询热门景区订单数据
type AbsenceBlackListReq struct {
	odas.Req
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewAbsenceBlackListReq(req *odas.Req, page, pageSize int) *AbsenceBlackListReq {
	return &AbsenceBlackListReq{
		Req:      *req,
		Page:     page,
		PageSize: page,
	}
}

func (h AbsenceBlackListReq) Api() string {
	params := h.Req.Params()
	return fmt.Sprintf("/v4/portrait/absenceBlackList?%s", params.Encode())
}

type AbsenceBlackListResponse struct {
	Total int                     `json:"total"`
	List  []*AbsenceBlackListItem `json:"list"`
}

type AbsenceBlackListItem struct {
	TouristName       string `json:"touristName"`
	LimitNum          int    `json:"limitNum"`
	LimitType         int    `json:"limitType"`
	LimitContent      string `json:"limitContent"`
	LimitIdentityType int    `json:"limitIdentityType"`
	LimitIdentityCode string `json:"limitIdentityCode"`
	CreatedTime       string `json:"createdTime"`
	UpdatedTime       string `json:"updatedTime"`
}
