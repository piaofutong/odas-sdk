package tourist

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ForecastPassengerFlowSummaryReq struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	Lid        string `json:"lid"`
	ExcludeLid string `json:"excludeLid"`
	Sid        int    `json:"sid"`
	OrderType  int    `json:"orderType"`
}

func (f ForecastPassengerFlowSummaryReq) Api() string {
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
	return fmt.Sprintf("/v4/tourist/forecastPassengerFlowSummary?%s", params.Encode())
}

func (f ForecastPassengerFlowSummaryReq) Body() []byte {
	return nil
}

func (f ForecastPassengerFlowSummaryReq) Method() string {
	return http.MethodGet
}

func (f ForecastPassengerFlowSummaryReq) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (f ForecastPassengerFlowSummaryReq) AuthRequired() bool {
	return true
}

func NewForecastPassengerFlowSummaryReq(
	start, end, lid, excludeLid string,
	sid, orderType int,
) *ForecastPassengerFlowSummaryReq {
	return &ForecastPassengerFlowSummaryReq{
		Start:      start,
		End:        end,
		Lid:        lid,
		ExcludeLid: excludeLid,
		Sid:        sid,
		OrderType:  orderType,
	}
}

type ForecastPassengerFlowSummaryResponse struct {
	Total int `json:"total"`
}
