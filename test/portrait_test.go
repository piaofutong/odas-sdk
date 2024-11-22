package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/portrait"
	"testing"
)

func TestService_City(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewCityReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, &odas.DateRangeCompareReq{
		CompareStart: startCompare,
		CompareEnd:   endCompare,
	}, "福建省", true, 10)
	var r []*portrait.CityRankResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Fellow(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewFellowReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, "福建省")
	var r portrait.FellowResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_PaymentMethod(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewPaymentMethodReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, "福建省", 10)
	var r []*portrait.PaymentMethodResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Province(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewProvinceReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, &odas.DateRangeCompareReq{
		CompareStart: startCompare,
		CompareEnd:   endCompare,
	}, 10, true)
	var r []*portrait.ProvinceRankResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_SexAgeSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewSexAgeSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, "福建省", true)
	var r portrait.AgeSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
