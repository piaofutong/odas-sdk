package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GroupByIdReq 根据id查询出入园统计组数据
type GroupByIdReq struct {
	Id int `json:"id"`
}

func (g GroupByIdReq) Api() string {
	params := url.Values{}
	if g.Id > 0 {
		params.Add("id", strconv.Itoa(g.Id))
	}
	return fmt.Sprintf("/v2/tourist/inout/groupById?%s", params.Encode())
}

func (g GroupByIdReq) Body() []byte {
	return nil
}

func (g GroupByIdReq) Method() string {
	return http.MethodGet
}

func (g GroupByIdReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (g GroupByIdReq) AuthRequired() bool {
	return true
}

func NewGroupByIdReq(id int) *GroupByIdReq {
	return &GroupByIdReq{
		Id: id,
	}
}

type GroupByIdResponse struct {
	GroupListResponse
}
