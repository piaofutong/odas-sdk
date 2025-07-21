package test

import (
	"testing"

	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/request/v4/report"
)

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
	}, report.WithTerminalType("1,2,4,19,20,46"))
	var r report.TerminalPassSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_TerminalPassSummaryGroupLid(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewTerminalPassSummaryGroupLidReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, report.WithTerminalType("1,2,4,19,20,46"))
	var r report.TerminalPassSummaryGroupLidResponse
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

func TestService_VerifiedSummaryHour(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := report.NewVerifiedSummaryHourReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	})
	var r report.VerifiedSummaryHourResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
