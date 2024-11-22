package odas

import (
	"net/http"
	"net/url"
	"strconv"
)

type Req struct {
	DateRangeReq
	Lid        string `json:"lid"`
	ExcludeLid string `json:"excludeLid"`
	DateType   int    `json:"dateType"`
	OrderType  int    `json:"orderType"`
}

type DateRangeCompareReq struct {
	CompareStart string `json:"compareStart"`
	CompareEnd   string `json:"compareEnd"`
}

func (r Req) Params() url.Values {
	params := url.Values{}
	if r.Lid != "" {
		params.Add("lid", r.Lid)
	}
	if r.ExcludeLid != "" {
		params.Add("excludeLid", r.ExcludeLid)
	}
	if r.Sid > 0 {
		params.Add("sid", strconv.Itoa(r.Sid))
	}
	if r.DateType > 0 {
		params.Add("dateType", strconv.Itoa(r.DateType))
	}
	if r.Start != "" {
		params.Add("start", r.Start)
	}
	if r.End != "" {
		params.Add("end", r.End)
	}
	if r.OrderType > 0 {
		params.Add("orderType", strconv.Itoa(r.OrderType))
	}
	return params
}

func (r Req) Body() []byte {
	return nil
}

func (r Req) Method() string {
	return http.MethodGet
}

func (r Req) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (r Req) AuthRequired() bool {
	return true
}

type DateRangeReq struct {
	Sid   int    `json:"sid"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (r DateRangeReq) Api() url.Values {
	params := url.Values{}
	if r.Sid > 0 {
		params.Add("sid", strconv.Itoa(r.Sid))
	}
	if r.Start != "" {
		params.Add("start", r.Start)
	}
	if r.End != "" {
		params.Add("end", r.End)
	}
	return params
}

func (r DateRangeReq) Body() []byte {
	return nil
}

func (r DateRangeReq) Method() string {
	return http.MethodGet
}

func (r DateRangeReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (r DateRangeReq) AuthRequired() bool {
	return true
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Pages    int `json:"pages"`
}

type BaseVerifiedSummaryVO struct {
	OrderNum             int `json:"orderNum"`
	OrderTicket          int `json:"orderTicket"`
	OrderAmount          int `json:"orderAmount"`
	OrderCostMoney       int `json:"orderCostMoney"`
	VerifiedNum          int `json:"verifiedNum"`
	VerifiedTicket       int `json:"verifiedTicket"`
	VerifiedAmount       int `json:"verifiedAmount"`
	VerifiedCostMoney    int `json:"verifiedCostMoney"`
	FinishedNum          int `json:"finishedNum"`
	FinishedTicket       int `json:"finishedTicket"`
	FinishedAmount       int `json:"finishedAmount"`
	FinishedCostMoney    int `json:"finishedCostMoney"`
	RevokedNum           int `json:"revokedNum"`
	RevokedTicket        int `json:"revokedTicket"`
	RevokedAmount        int `json:"revokedAmount"`
	RevokedCostMoney     int `json:"revokedCostMoney"`
	CancelNum            int `json:"cancelNum"`
	CancelTicket         int `json:"cancelTicket"`
	CancelAmount         int `json:"cancelAmount"`
	CancelCostMoney      int `json:"cancelCostMoney"`
	PrintNum             int `json:"printNum"`
	AfterSaleTicketNum   int `json:"afterSaleTicketNum"`
	AfterSaleRefundMoney int `json:"afterSaleRefundMoney"`
	AfterSaleIncomeMoney int `json:"afterSaleIncomeMoney"`
}

type StatDO struct {
	In          int `json:"in"`           // 入园数
	Out         int `json:"out"`          // 出园数
	Employee    int `json:"employee"`     // 员工卡入园数
	DeviceCount int `json:"device_count"` // 设备数
}
