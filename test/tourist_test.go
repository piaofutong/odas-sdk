package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/tourist"
	"strconv"
	"testing"
)

func TestService_FlowByDevice(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewFlowByDeviceReq("5da80a7dc69bdcc503ec5da69888f8c1", 9)
	var r tourist.FlowByDeviceResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_FlowByGids(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewFlowByGIdsReq("42,43", "2024-11-22")
	var r tourist.FlowByGIdsResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_FlowBySid(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewFlowBySidReq("3385")
	var r tourist.FlowBySidResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_ForecastPassengerFlowList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewForecastPassengerFlowListReq(start, end, lid, excludeLid, sid, 0)
	var r tourist.ForecastPassengerFlowListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_ForecastPassengerFlowSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewForecastPassengerFlowSummaryReq(start, end, lid, excludeLid, sid, 0)
	var r tourist.ForecastPassengerFlowSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_GroupById(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewGroupByIdReq(42)
	var r tourist.GroupByIdResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_GroupList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewGroupListReq(sid)
	var r []*tourist.GroupListResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_InoutSummaryByTime(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewSummaryByTimeReq(
		tourist.WithStart(start),
		tourist.WithEnd(end),
		tourist.WithSid(sid),
		tourist.WithGid(strconv.Itoa(gid)),
		tourist.WithDateType(0),
		tourist.WithNoAmend(),
	)
	var r tourist.InoutSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_InoutSummaryByDate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewSummaryByDateReq(
		tourist.WithStart(start),
		tourist.WithEnd(end),
		tourist.WithSid(sid),
		tourist.WithGid(strconv.Itoa(gid)),
		tourist.WithDateType(0),
		tourist.WithNoAmend(),
	)
	var r tourist.InoutSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_DailyPassengerFlow(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewDailyPassengerFlowReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, true)
	var r tourist.PassengerFlowByDateResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_DailyPassengerFlowByVerify(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewDailyPassengerFlowByVerifyReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, true)
	var r tourist.PassengerFlowByDateResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_TouristLocal(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewLocalReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, tourist.WithLocalLimit(10), tourist.WithLocalProvince("福建省"), tourist.WithLocalUnknown(true))
	var r tourist.LocalResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_TouristLocalByVerify(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewLocalByVerifyReq(&odas.Req{
		DateRangeReq: odas.DateRangeReq{
			Sid:   sid,
			Start: start,
			End:   end,
		},
		Lid:        lid,
		ExcludeLid: excludeLid,
	}, tourist.WithLocalLimit(10), tourist.WithLocalProvince("福建省"), tourist.WithLocalUnknown(true))
	var r tourist.LocalResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
