package order

import (
	"fmt"

	"github.com/piaofutong/odas-sdk/odas"
)

type PreBookingAgeGenderDistReq struct {
	odas.Req
}

func (p PreBookingAgeGenderDistReq) Api() string {
	params := p.Req.Api()

	return fmt.Sprintf("/v4/order/preBookingAgeGenderDist?%s", params.Encode())
}

func NewPreBookingAgeGenderDistReq(req odas.Req) *PreBookingAgeGenderDistReq {
	return &PreBookingAgeGenderDistReq{
		Req: req,
	}
}

type PreBookingAgeGenderDistResponse struct {
	List []*AgeGenderCountListItem `json:"list"`
}

type AgeGenderCountListItem struct {
	MaleAge0to7     int `json:"maleAge0to7"`
	MaleAge8to17    int `json:"maleAge8to17"`
	MaleAge18to27   int `json:"maleAge18to27"`
	MaleAge28to40   int `json:"maleAge28to40"`
	MaleAge41to60   int `json:"maleAge41to60"`
	MaleAge61plus   int `json:"maleAge61plus"`
	FemaleAge0to7   int `json:"femaleAge0to7"`
	FemaleAge8to17  int `json:"femaleAge8to17"`
	FemaleAge18to27 int `json:"femaleAge18to27"`
	FemaleAge28to40 int `json:"femaleAge28to40"`
	FemaleAge41to60 int `json:"femaleAge41to60"`
	FemaleAge61plus int `json:"femaleAge61plus"`
	Time            int `json:"time"`
}
