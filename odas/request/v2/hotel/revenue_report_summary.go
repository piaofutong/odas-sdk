package hotel

import (
	"fmt"
	"github.com/piaofutong/odas-sdk/odas"
)

type RevenueReportSummaryReq struct {
	odas.DateRangeReq
	CodeCategory string `json:"codeCategory"`
}

func NewRevenueReportSummary(req *odas.DateRangeReq, codeCategory string) *RevenueReportSummaryReq {
	return &RevenueReportSummaryReq{
		DateRangeReq: *req,
		CodeCategory: codeCategory,
	}
}

func (r RevenueReportSummaryReq) Api() string {
	params := r.DateRangeReq.Api()
	if r.CodeCategory != "" {
		params.Add("codeCategory", r.CodeCategory)
	}
	return fmt.Sprintf("/v2/hotel/revenueReportSummary?%s", params.Encode())
}

type RevenueReportSummaryResponse struct {
	Total *RevenueReportTotal  `json:"total"`
	List  []*RevenueReportData `json:"list"`
}

type RevenueReportTotal struct {
	RevTotal    float64 `json:"revTotal"`
	RevRm       float64 `json:"revRm"`
	RevFb       float64 `json:"revFb"`
	RevMt       float64 `json:"revMt"`
	RevEn       float64 `json:"revEn"`
	RevSp       float64 `json:"revSp"`
	RevOt       float64 `json:"revOt"`
	RoomsTotal  float64 `json:"roomsTotal"`
	RoomsArr    float64 `json:"roomsArr"`
	RoomsDep    float64 `json:"roomsDep"`
	RoomsNoShow int     `json:"roomsNoShow"`
	RoomsCxl    int     `json:"roomsCxl"`
	People      int     `json:"people"`
	PeopleArr   int     `json:"peopleArr"`
	PeopleDep   int     `json:"peopleDep"`
}

type RevenueReportData struct {
	CodeCategory    string `json:"codeCategory"`
	CodeCategoryDes string `json:"codeCategoryDes"`
	RevenueReportTotal
}
