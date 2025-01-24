package account

import "github.com/piaofutong/odas-sdk/odas"

type MyTicketReq struct {
	odas.Req
}

func NewMyTicketReq() *MyTicketReq {
	return &MyTicketReq{}
}

func (r *MyTicketReq) Api() string {
	return "/account/my_tickets"
}

type MyTicketResponse struct {
	//Total int              `json:"total"`
	List []*MyTicketsItem `json:"list"`
}

type MyTicketsItem struct {
	Id    int    `json:"id"`    // 表id
	Title string `json:"title"` // 景区名称
}
