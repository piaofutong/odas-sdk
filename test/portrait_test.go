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
	}, portrait.WithCityLimit(10), portrait.WithCityUnknown(true), portrait.WithCityProvince("福建省"))
	var r []*portrait.CityRankResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_CityByVerify(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewCityByVerifyReq(&odas.Req{
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
	}, portrait.WithCityLimit(10), portrait.WithCityUnknown(true), portrait.WithCityProvince("福建省"))
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
	}, portrait.WithFellowProvince("福建省"))
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
	}, portrait.WithPaymentMethodLimit(10), portrait.WithPaymentMethodProvince("福建省"))
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
	}, portrait.WithProvinceLimit(10), portrait.WithProvinceUnknown(true))
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
	}, portrait.WithSexAgeUnknown(true), portrait.WithSexAgeProvince("福建省"))
	var r portrait.AgeSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_ProvinceByVerify(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewProvinceByVerifyReq(&odas.Req{
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
	}, portrait.WithProvinceLimit(10), portrait.WithProvinceUnknown(true))
	var r []*portrait.ProvinceRankResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_SexAgeSummaryByVerify(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := portrait.NewSexAgeSummaryByVerifyReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, portrait.WithSexAgeUnknown(true), portrait.WithSexAgeProvince("福建省"))
	var r portrait.AgeSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
