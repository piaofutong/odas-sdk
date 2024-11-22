package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/report"
	"testing"
)

func TestService_TerminalPassHourSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewTerminalPassHourSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r report.TerminalPassHourSummaryReq
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_TerminalPassSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewTerminalPassSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r report.TerminalPassSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_ReportTicketList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewTicketListReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, 10, "2598429,347718")
	var r report.TicketListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_VerifiedSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewVerifiedSummaryReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r report.VerifiedSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
