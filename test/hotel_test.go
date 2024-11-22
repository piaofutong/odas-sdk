package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/hotel"
	"testing"
)

func TestService_Occupancy(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := hotel.NewOccupancyReq(&odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	})
	var r hotel.OccupancyResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_RevenueReportSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := hotel.NewRevenueReportSummary(&odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}, "A")
	var r *hotel.RevenueReportSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_RmOrderDateList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := hotel.NewRmOrderDateListReq(&odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	})
	var r *hotel.RmOrderDateListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_RmSaleReportDateList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := hotel.NewRmSaleReportDateListReq(&odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	})
	var r *hotel.RmSaleReportDateListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_RmSaleReportList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := hotel.NewRmSaleReportListReq(&odas.DateRangeReq{
		Sid:   sid,
		Start: start,
		End:   end,
	})
	var r *hotel.RmSaleReportListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
