package tourist

import (
	"fmt"
	"net/http"
)

type InoutByGroupId struct {
	GroupId int
}

func (o *InoutByGroupId) SetGroupId(groupId int) {
	o.GroupId = groupId
}

func (o *InoutByGroupId) Api() string {
	return fmt.Sprintf("/tourist/tourist/inout/flow?gid=%d", o.GroupId)
}

func (o *InoutByGroupId) Body() []byte {
	return nil
}

func (o *InoutByGroupId) Method() string {
	return http.MethodGet
}

func (o *InoutByGroupId) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (o *InoutByGroupId) AuthRequired() bool {
	return true
}

func NewInoutByGroupId(groupId int) *InoutByGroupId {
	return &InoutByGroupId{GroupId: groupId}
}

type InoutByGroupIdResponse struct {
	Total struct {
		In  int `json:"in"`
		Out int `json:"out"`
	} `json:"total"`
	List []struct {
		Time string `json:"time"`
		In   int    `json:"in"`
		Out  int    `json:"out"`
	} `json:"list"`
}
