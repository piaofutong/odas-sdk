package test

import (
	"strconv"
	"testing"

	"github.com/piaofutong/odas-sdk/odas"
	touristV2 "github.com/piaofutong/odas-sdk/odas/request/v2/tourist"
	"github.com/piaofutong/odas-sdk/odas/request/v4/tourist"
)

func TestService_FlowByDevice(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := touristV2.NewFlowByDeviceReq("5da80a7dc69bdcc503ec5da69888f8c1", 1)
	var r touristV2.FlowByDeviceResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_FlowByGids(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := touristV2.NewFlowByGIdsReq("42,43", "2024-11-22")
	var r touristV2.FlowByGIdsResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_FlowBySid(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := touristV2.NewFlowBySidReq("3385")
	var r touristV2.FlowBySidResponse
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
	req := touristV2.NewGroupByIdReq(42)
	var r touristV2.GroupByIdResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_GroupList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := touristV2.NewGroupListReq(sid)
	var r []*touristV2.GroupListResponse
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
