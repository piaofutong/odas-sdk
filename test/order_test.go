package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/order"
	"testing"
)

func TestService_BookingOrderList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewBookingOrderListReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r order.BookingOrderListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_BookingTeamOrder(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewBookingTeamOrderReq(&odas.Req{
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
	})
	var r order.BookingTeamOrderResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
func TestService_Hot(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewHotReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, 10)
	var r []*order.HotResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
func TestService_Summary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, order.WithOrderCompare())
	var r order.SummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
func TestService_ToiSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewToiSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r order.ToiSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_PrebookingByType(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewPreBookingByTypeReq(odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}, order.WithOrderType(0), order.WithLid(lid), order.WithExcludeLid(excludeLid))
	var r order.PreBookingByTypeResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_PrebookingSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := order.NewPreBookingByTypeReq(odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}, order.WithOrderType(0), order.WithLid(lid), order.WithExcludeLid(excludeLid))
	var r order.PreBookingSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
