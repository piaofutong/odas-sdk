package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/product"
	"testing"
)

func TestService_Rank(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := product.NewRankReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, product.WithRankLimit(10))
	var r []*product.RankResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_SalesDetail(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := product.NewSalesDetailReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, product.WithSalesDetailTicketId([]int{2598429, 347718}))
	var r []*product.SalesDetailResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_TicketList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := product.NewTicketListReq(&odas.Req{
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
	}, 1, 10)
	var r product.TicketListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
