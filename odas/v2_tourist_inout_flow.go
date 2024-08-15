package odas

import (
	"fmt"
)

type V2InoutByGroupId struct {
	GroupId int
}

func (o *V2InoutByGroupId) SetGroupId(groupId int) {
	o.GroupId = groupId
}

func (o *V2InoutByGroupId) Api() string {
	return fmt.Sprintf("/v2/tourist/inout/flow?gid=%d", o.GroupId)
}

func (o *V2InoutByGroupId) Body() []byte {
	return nil
}

func (o *V2InoutByGroupId) Method() string {
	return "GET"
}

func (o *V2InoutByGroupId) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func NewV2InoutByGroupId(groupId int) *V2InoutByGroupId {
	return &V2InoutByGroupId{GroupId: groupId}
}

type V2InoutByGroupIdResponse struct {
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
