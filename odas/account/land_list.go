package account

import "github.com/piaofutong/odas-sdk/odas"

type MyLandListReq struct {
	odas.Req
}

func NewMyLandListReq() *MyLandListReq {
	return &MyLandListReq{}
}

func (r *MyLandListReq) Api() string {
	return "/account/my_lands"
}

type MyLandListResponse struct {
	//Total int              `json:"total"`
	List []*MyLandItem `json:"list"`
}

type MyLandItem struct {
	Id    int    `json:"id"`    // 表id
	Title string `json:"title"` // 景区名称
}
