package test

import (
	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/tourist"
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

func TestService_InoutSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewSummaryReq(start, end, sid, 42, true)
	var r tourist.InoutSummaryResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_PassengerFlowByDate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := tourist.NewPassengerFlowByDateReq(&odas.Req{
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
	}, "福建省", 10, true)
	var r tourist.LocalResponse
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
}
