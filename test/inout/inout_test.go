package inout_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/common"
	inoutV5 "github.com/piaofutong/odas-sdk/odas/inout"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/pb"
	"gitlab.12301.test/gopkg/idss-go-sdk/pb/inout"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 测试配置变量
var (
	accessId  = "abcdefg"
	accessKey = "abcdefg"
	token     = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	sid       = []int64{3385}
	gid       = []int64{1, 48, 122, 56, 34, 60, 59, 128, 140}
	devices   = []string{"device1", "device2"}
	startTime = time.Now().AddDate(-1, 0, 0)
	endTime   = time.Now()

	granter int64 = sid[0]
	grantee int64 = 1003
)

var (
	ctx               context.Context
	idssConn          *grpc.ClientConn
	amendmentConn     inout.AmendmentServiceClient
	authorizationConn inout.AuthorizationServiceClient
	groupConn         inout.GroupServiceClient
	statConn          inout.StatServiceClient
)

// TestMain 设置测试环境
func TestMain(m *testing.M) {
	var err error
	// 设置本地模式，用于测试
	odas.SetLocalMode()
	idssConn, err = idsssdk.NewClient(
		context.Background(),
		"127.0.0.1:50052",
		"rklXMeLpnwK2fz1B",
		"Oyr4nbXc67TYFZ8VISLfHsK9JCjgRiDM",
	)
	idssConn, err = idsssdk.NewClient(
		context.Background(),
		"127.0.0.1:50052",
		"M7GoG0xSmPSUqhEA",
		"KIr20mMACKCqrDj5U0PBDf7XapVt31dI",
	)

	if err != nil {
		fmt.Printf("NewIDSSClient err: %v", err)
		os.Exit(-1)
	}

	ctx = context.Background()
	amendmentConn = inout.NewAmendmentServiceClient(idssConn)
	authorizationConn = inout.NewAuthorizationServiceClient(idssConn)
	groupConn = inout.NewGroupServiceClient(idssConn)
	statConn = inout.NewStatServiceClient(idssConn)

	os.Exit(m.Run())
}

// // TestAuthorizationCreate 测试创建授权接口
// func TestAuthorizationCreate(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AuthorizationCreateReq{
// 		Granter: granter,
// 		Grantee: grantee,
// 		GroupId: gid[0],
// 	}
// 	var r inoutV5.AuthorizationCreateResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if id := getAuthorizationId(t, gid[0]); id == 0 {
// 		t.Fatalf("测试失败，数据返回为空，id: %d", id)
// 	}
// }

// // TestAuthorizationList 测试获取授权列表接口
// func TestAuthorizationList(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AuthorizationListReq{
// 		Granter:  sid[0],
// 		Page:     1,
// 		PageSize: 10,
// 	}
// 	var r inoutV5.AuthorizationListResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	grpcResp, err := authorizationConn.List(ctx, &inout.AuthorizationListRequest{
// 		Granter:  sid[0],
// 		Page:     1,
// 		PageSize: 10,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "AuthorizationList")
// }

// func getAuthorizationId(t *testing.T, gid int64) int64 {
// 	grpcResp, err := authorizationConn.List(ctx, &inout.AuthorizationListRequest{
// 		Granter:  sid[0],
// 		Page:     1,
// 		PageSize: 100,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, v := range grpcResp.List {
// 		if v.GroupId != gid {
// 			continue
// 		}

// 		if v.Grantee == grantee && v.Granter == granter {
// 			return v.Id
// 		}
// 	}

// 	return 0
// }

// // TestAuthorizationUpdate 测试更新授权接口
// func TestAuthorizationUpdate(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AuthorizationUpdateReq{
// 		Id:      getAuthorizationId(t, gid[0]),
// 		GroupId: gid[1],
// 	}
// 	var r inoutV5.AuthorizationUpdateResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := authorizationConn.Update(ctx, &inout.AuthorizationUpdateRequest{
// 		Id:      getAuthorizationId(t, gid[1]),
// 		GroupId: gid[0],
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "AuthorizationUpdate")
// }

// // TestAuthorizationDelete 测试删除授权接口
// func TestAuthorizationDelete(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AuthorizationDeleteReq{
// 		Id: getAuthorizationId(t, gid[0]),
// 	}
// 	var r inoutV5.AuthorizationDeleteResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if id := getAuthorizationId(t, gid[0]); id != 0 {
// 		t.Fatalf("测试失败，数据返回不为空，id: %d", id)
// 	}

// }

// // TestSave 测试保存优化客流接口
// func TestSave(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	switch1 := int64(1)
// 	inPlus := int64(10)
// 	outPlus := int64(-5)
// 	req := &inoutV5.AmendmentSaveReq{
// 		Gid:            gid[0],
// 		Switch:         &switch1,
// 		Type:           0, // 加减方式
// 		InPlus:         &inPlus,
// 		OutPlus:        &outPlus,
// 		InCoefficient:  1.0,
// 		OutCoefficient: 1.0,
// 	}
// 	var r inoutV5.AmendmentSaveResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := amendmentConn.Save(ctx, &inout.AmendmentSaveRequest{
// 		Gid:            gid[0],
// 		Switch:         &switch1,
// 		Type:           0,
// 		InPlus:         &inPlus,
// 		OutPlus:        &outPlus,
// 		InCoefficient:  1.0,
// 		OutCoefficient: 1.0,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "AmendmentSave")
// }

// func getAmendmentId(t *testing.T, gid int64) int64 {
// 	grpcResp, err := amendmentConn.Get(ctx, &inout.AmendmentGetRequest{
// 		Gid: gid,
// 	})
// 	if err != nil {
// 		return 0
// 	}

// 	if grpcResp.CreatedAt == "" {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	return grpcResp.Id
// }

// // TestGet 测试获取优化客流详情接口
// func TestGet(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AmendmentGetReq{
// 		Gid: gid[0],
// 	}
// 	var r inoutV5.AmendmentGetResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := amendmentConn.Get(ctx, &inout.AmendmentGetRequest{
// 		Gid: gid[0],
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if grpcResp.CreatedAt == "" {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "AmendmentGet")
// }

// // TestDelete 测试删除优化客流接口
// func TestDelete(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.AmendmentDeleteReq{
// 		Id: getAmendmentId(t, gid[0]),
// 	}
// 	var r inoutV5.AmendmentDeleteResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if id := getAmendmentId(t, gid[0]); id != 0 {
// 		t.Fatalf("测试失败，数据返回不为空，id: %d", id)
// 	}
// }

// // TestCreate 测试创建出入园统计组接口
// func TestCreate(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	gates := []string{"gate1", "gate2", "gate3", "gate4"}
// 	gates = append(gates, time.Now().String())

// 	req := &inoutV5.GroupCreateReq{
// 		Sid:        (sid[0]),
// 		Name:       "测试统计组",
// 		Gates:      gates,
// 		Capacity:   1000,
// 		UpperLimit: 5000,
// 		Config: []inoutV5.GroupConfig{
// 			{Label: "低客流", Min: 0, Max: 30, Color: "#00FF00"},
// 			{Label: "中客流", Min: 31, Max: 70, Color: "#FFFF00"},
// 			{Label: "高客流", Min: 71, Max: 100, Color: "#FF0000"},
// 		},
// 	}
// 	var r inoutV5.GroupCreateResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if r.Id == 0 || r.SerialNo == "" {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}
// }

// // TestGroupGet 测试获取出入园统计组详情接口
// func TestGroupGet(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.GroupGetReq{
// 		Id: gid[0],
// 	}
// 	var r inoutV5.GroupGetResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := groupConn.Get(ctx, &inout.GroupGetRequest{
// 		Id: gid[0],
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if grpcResp.Id == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "GroupGet")
// }

// // TestGroupList 测试获取出入园统计组列表接口
// func TestGroupList(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.GroupListReq{
// 		Sid:   sid[0],
// 		Owner: true,
// 	}
// 	var r inoutV5.GroupListResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := groupConn.List(ctx, &inout.GroupListRequest{
// 		Sid:   sid[0],
// 		Owner: true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "GroupList")
// }

// // TestList 测试获取优化客流列表接口
// func TestList(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	switch1 := int64(1)
// 	req := &inoutV5.AmendmentListReq{
// 		Sid:      sid[0],
// 		Switch:   &switch1,
// 		Page:     1,
// 		PageSize: 10,
// 	}
// 	var r inoutV5.AmendmentListResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := amendmentConn.List(ctx, &inout.AmendmentListRequest{
// 		Sid:      sid[0],
// 		Switch:   &switch1,
// 		Page:     1,
// 		PageSize: 10,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "AmendmentList")
// }

// // TestUpdate 测试更新出入园统计组接口
// func TestUpdate(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.GroupUpdateReq{
// 		Id:         gid[0],
// 		Name:       "更新后的测试统计组",
// 		Gates:      []string{"gate1", "gate2", "gate3"},
// 		Capacity:   1200,
// 		UpperLimit: 6000,
// 		Config: []inoutV5.GroupConfig{
// 			{Label: "低客流", Min: 0, Max: 25, Color: "#00FF00"},
// 			{Label: "中客流", Min: 26, Max: 75, Color: "#FFFF00"},
// 			{Label: "高客流", Min: 76, Max: 100, Color: "#FF0000"},
// 		},
// 	}
// 	var r inoutV5.GroupUpdateResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := groupConn.Update(ctx, &inout.GroupUpdateRequest{
// 		Id:         gid[0],
// 		Name:       "更新后的测试统计组",
// 		Gates:      []string{"gate1", "gate2", "gate3"},
// 		Capacity:   1200,
// 		UpperLimit: 6000,
// 		Config: []*inout.GroupConfig{
// 			{Label: "低客流", Min: 0, Max: 25, Color: "#00FF00"},
// 			{Label: "中客流", Min: 26, Max: 75, Color: "#FFFF00"},
// 			{Label: "高客流", Min: 76, Max: 100, Color: "#FF0000"},
// 		},
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "GroupUpdate")
// }

// // TestHourSummaryByDevice 测试按设备小时统计出入园数据接口
// func TestHourSummaryByDevice(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	dateStr := time.Now().Format("2006-01-02")
// 	req := &inoutV5.HourSummaryByDeviceReq{
// 		Date:    dateStr,
// 		Sid:     sid[0],
// 		Devices: devices,
// 		Hour:    10,
// 	}
// 	var r inoutV5.HourSummaryByDeviceResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := statConn.HourSummaryByDevice(ctx, &inout.HourSummaryByDeviceRequest{
// 		Date:    dateStr,
// 		Sid:     sid[0],
// 		Devices: devices,
// 		Hour:    10,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// TODO: 测试数据为空的情况
// 	// if len(grpcResp.List) == 0 {
// 	// 	t.Fatalf("测试失败，数据返回为空")
// 	// }

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "HourSummaryByDevice")
// }

// TestSummaryByDate 测试按日期维度统计出入园数据接口
func TestSummaryByDate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.SummaryByDateReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Sid:   sid,
		},
		Gid:     gid,
		NoAmend: false,
	}
	var apiResp inoutV5.SummaryByDateResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.SummaryByDate(ctx, &inout.SummaryByDateRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Start: timestamppb.New(startTime),
			End:   timestamppb.New(endTime),
			Sid:   sid,
		},
		Gid:     gid,
		NoAmend: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTime")
}

// TestSummaryByTime 测试按时间维度统计出入园数据接口
func TestSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.SummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Sid:   sid,
		},
		Gid:     gid,
		NoAmend: false,
	}
	grpcReq := &inout.SummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Start: timestamppb.New(startTime),
			End:   timestamppb.New(endTime),
			Sid:   sid,
		},
		Gid:     gid,
		NoAmend: false,
	}

	var apiResp inoutV5.SummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.Summary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTime")
}

// TestSummaryByTime 测试按时间维度统计出入园数据接口 - 不传Sid
func TestSummaryAndMisSid(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.SummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
		},
		Gid:     gid,
		NoAmend: false,
	}
	grpcReq := &inout.SummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Start: timestamppb.New(startTime),
			End:   timestamppb.New(endTime),
			Sid:   sid,
		},
		Gid:     gid,
		NoAmend: false,
	}

	var apiResp inoutV5.SummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.Summary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TestSummaryAndMisSid")
}

// TestSummaryByTime 测试按时间维度统计出入园数据接口 - Sid和gid权限不匹配
func TestSummaryAndSidMismatch(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.SummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Sid:   []int64{33850},
		},
		Gid:     gid,
		NoAmend: false,
	}

	var apiResp inoutV5.SummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err == nil {
		t.Fatal("测试失败，请求成功但应该失败原因为：供应商权限校验未通过")
	}

	if !strings.Contains(err.Error(), "供应商权限校验未通过") {
		t.Fatalf("测试失败，错误不匹配(%v != 供应商权限校验未通过)", err.Error())
	}
}

// // TestTodaySummaryByGroupHour 测试今日分组小时汇总接口
// func TestTodaySummaryByGroupHour(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &inoutV5.SummaryByGroupReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Sid:   sid,
// 		},
// 		Gid:     gid,
// 		NoAmend: false,
// 	}
// 	var r inoutV5.SummaryByGroupResp
// 	err := iam.Do(req, &r, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	grpcResp, err := statConn.SummaryByGroup(ctx, &inout.SummaryByGroupRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Start: timestamppb.New(startTime),
// 			End:   timestamppb.New(endTime),
// 			Sid:   sid,
// 		},
// 		Gid:     gid,
// 		NoAmend: false,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &r, grpcResp, "TodaySummaryByGroupHour")

// }

// TestSummaryByTime 测试按时间维度统计出入园数据接口
func TestTodaySummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.TodaySummaryReq{
		Gid:     gid,
		NoAmend: false,
	}
	grpcReq := &inout.SummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Start: timestamppb.New(time.Now()),
			End:   timestamppb.New(time.Now()),
			Sid:   sid,
		},
		Gid:     []int64{gid[0]},
		NoAmend: false,
	}

	var apiResp inoutV5.SummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.Summary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTime")
}
