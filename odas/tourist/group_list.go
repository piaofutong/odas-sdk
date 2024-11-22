package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GroupListReq 获取账号的统计组
type GroupListReq struct {
	Sid int `json:"sid"`
}

func (g GroupListReq) Api() string {
	params := url.Values{}
	if g.Sid > 0 {
		params.Add("sid", strconv.Itoa(g.Sid))
	}
	return fmt.Sprintf("/v2/tourist/inout/groupList?%s", params.Encode())
}

func (g GroupListReq) Body() []byte {
	return nil
}

func (g GroupListReq) Method() string {
	return http.MethodGet
}

func (g GroupListReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (g GroupListReq) AuthRequired() bool {
	return true
}

func NewGroupListReq(sid int) *GroupListReq {
	return &GroupListReq{
		Sid: sid,
	}
}

type GroupListResponse struct {
	Id         int                 `json:"id"`
	Sid        int                 `json:"sid"`
	SerialNo   string              `json:"serialNo"`
	Name       string              `json:"name"`
	Gates      []string            `json:"gates"`
	Capacity   int                 `json:"capacity"`
	UpperLimit int                 `json:"upperLimit"`
	Config     []*InoutGroupConfig `json:"config"`
	CreatedAt  string              `json:"createdAt"`
	UpdatedAt  string              `json:"updatedAt"`
}

type InoutGroupConfig struct {
	Label string `json:"label" binding:"required,gt=0"`
	Min   int    `json:"min" binding:"required,gt=0"`
	Max   int    `json:"max" binding:"required,gt=0"`
	Color string `json:"color" binding:"required,gt=0"`
}
