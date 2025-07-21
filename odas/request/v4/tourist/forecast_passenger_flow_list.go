package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ForecastPassengerFlowListReq 预测客流每日以及汇总数据数据
type ForecastPassengerFlowListReq struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	Lid        string `json:"lid"`
	ExcludeLid string `json:"excludeLid"`
	Sid        int    `json:"sid"`
	OrderType  int    `json:"orderType"`
}

func (f ForecastPassengerFlowListReq) Api() string {
	params := url.Values{}
	if f.Start != "" {
		params.Add("start", f.Start)
	}
	if f.End != "" {
		params.Add("end", f.End)
	}
	if f.Lid != "" {
		params.Add("lid", f.Lid)
	}
	if f.Sid > 0 {
		params.Add("sid", strconv.Itoa(f.Sid))
	}
	if f.OrderType > 0 {
		params.Add("orderType", strconv.Itoa(f.OrderType))
	}
	if f.ExcludeLid != "" {
		params.Add("excludeLid", f.ExcludeLid)
	}
	return fmt.Sprintf("/v4/tourist/forecastPassengerFlowList?%s", params.Encode())
}

func (f ForecastPassengerFlowListReq) Body() []byte {
	return nil
}

func (f ForecastPassengerFlowListReq) Method() string {
	return http.MethodGet
}

func (f ForecastPassengerFlowListReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (f ForecastPassengerFlowListReq) AuthRequired() bool {
	return true
}

func NewForecastPassengerFlowListReq(
	start, end, lid, excludeLid string,
	sid, orderType int,
) *ForecastPassengerFlowListReq {
	return &ForecastPassengerFlowListReq{
		Start:      start,
		End:        end,
		Lid:        lid,
		ExcludeLid: excludeLid,
		Sid:        sid,
		OrderType:  orderType,
	}
}

type ForecastPassengerFlowListResponse struct {
	Total  *FlowForecastTotal    `json:"total"`
	Detail []*FlowForecastDetail `json:"detail"`
}

type FlowForecastTotal struct {
	TodayFlow    int `json:"todayFlow"`
	TomorrowFlow int `json:"tomorrowFlow"`
	ThirdDayFlow int `json:"thirdDayFlow"`
}

type FlowForecastDetail struct {
	TimeRange string `json:"timeRange"`
	Count     int    `json:"count"`
}
