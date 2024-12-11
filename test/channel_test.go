package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/channel"
	"testing"
)

func TestService_OrderChannel(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := channel.NewOrderChannelReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r []*channel.OrderChannelResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_OrderFullChannel(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := channel.NewOrderFullChannelReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, channel.WithLimit(10))
	var r channel.OrderFullChannelResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_StatDistributorSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := channel.NewStatDistributorSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r []*channel.StatDistributorSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_OrderSecondaryChannel(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := channel.NewOrderSecondaryChannel(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, channel.WithSecondaryChannelLimit(10), channel.WithSecondaryChannelClassId(1))
	var r channel.OrderFullChannelResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
