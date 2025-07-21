package inout_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/request/common"
	inoutV5 "github.com/piaofutong/odas-sdk/odas/request/v5/inout"
	"github.com/piaofutong/odas-sdk/test/compare"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss"
	"gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss/inout"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 测试配置变量
var (
	accessId  = "abcdefg"
	accessKey = "abcdefg"
	token     = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	sid       = int64(3385)
	gid       = []int64{12, 13}
	devices   = []string{"device1", "device2"}
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

// TestAuthorizationList 测试获取授权列表接口
func TestAuthorizationList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AuthorizationListReq{
		Granter:  sid,
		Page:     1,
		PageSize: 10,
	}
	var r inoutV5.AuthorizationListResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}
	grpcResp, err := authorizationConn.List(ctx, &inout.AuthorizationListRequest{
		Granter:  sid,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AuthorizationList")
}

// TestAuthorizationCreate 测试创建授权接口
func TestAuthorizationCreate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AuthorizationCreateReq{
		Granter: sid,
		Grantee: 1003,
		GroupId: gid[0],
	}
	var r inoutV5.AuthorizationCreateResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := authorizationConn.Create(ctx, &inout.AuthorizationCreateRequest{
		Granter: sid,
		Grantee: 1003,
		GroupId: gid[0],
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AuthorizationCreate")
}

// TestAuthorizationUpdate 测试更新授权接口
func TestAuthorizationUpdate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AuthorizationUpdateReq{
		Id:      2,
		GroupId: gid[0],
	}
	var r inoutV5.AuthorizationUpdateResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := authorizationConn.Update(ctx, &inout.AuthorizationUpdateRequest{
		Id:      2,
		GroupId: gid[0],
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AuthorizationUpdate")
}

// TestAuthorizationDelete 测试删除授权接口
func TestAuthorizationDelete(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AuthorizationDeleteReq{
		Id: 1,
	}
	var r inoutV5.AuthorizationDeleteResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := authorizationConn.Delete(ctx, &inout.AuthorizationDeleteRequest{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AuthorizationDelete")
}

// TestCreate 测试创建出入园统计组接口
func TestCreate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.GroupCreateReq{
		Sid:        int(sid),
		Name:       "测试统计组",
		Gates:      []string{"gate1", "gate2"},
		Capacity:   1000,
		UpperLimit: 5000,
		Config: []*inoutV5.GroupConfig{
			{Label: "低客流", Min: 0, Max: 30, Color: "#00FF00"},
			{Label: "中客流", Min: 31, Max: 70, Color: "#FFFF00"},
			{Label: "高客流", Min: 71, Max: 100, Color: "#FF0000"},
		},
	}
	var r inoutV5.GroupCreateResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := groupConn.Create(ctx, &inout.GroupCreateRequest{
		Sid:        sid,
		Name:       "测试统计组",
		Gates:      []string{"gate1", "gate2"},
		Capacity:   1000,
		UpperLimit: 5000,
		Config: []*inout.GroupConfig{
			{Label: "低客流", Min: 0, Max: 30, Color: "#00FF00"},
			{Label: "中客流", Min: 31, Max: 70, Color: "#FFFF00"},
			{Label: "高客流", Min: 71, Max: 100, Color: "#FF0000"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "GroupCreate")
}

// TestDelete 测试删除优化客流接口
func TestDelete(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AmendmentDeleteReq{
		Id: gid[0],
	}
	var r inoutV5.AmendmentDeleteResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := amendmentConn.Delete(ctx, &inout.AmendmentDeleteRequest{
		Id: gid[0],
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AmendmentDelete")
}

// TestGet 测试获取优化客流详情接口
func TestGet(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AmendmentGetReq{
		Gid: gid[0],
	}
	var r inoutV5.AmendmentGetResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := amendmentConn.Get(ctx, &inout.AmendmentGetRequest{
		Gid: gid[0],
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AmendmentGet")
}

// TestGroupGet 测试获取出入园统计组详情接口
func TestGroupGet(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.GroupGetReq{
		Id: gid[0],
	}
	var r inoutV5.GroupGetResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := groupConn.Get(ctx, &inout.GroupGetRequest{
		Id: gid[0],
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "GroupGet")
}

// TestGroupList 测试获取出入园统计组列表接口
func TestGroupList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.GroupListReq{
		Sid:   sid,
		Owner: true,
	}
	var r inoutV5.GroupListResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := groupConn.List(ctx, &inout.GroupListRequest{
		Sid:   sid,
		Owner: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "GroupList")
}

// TestList 测试获取优化客流列表接口
func TestList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.AmendmentListReq{
		Sid:      sid,
		Switch:   1,
		Page:     1,
		PageSize: 10,
	}
	var r inoutV5.AmendmentListResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	switch1 := int64(1)
	grpcResp, err := amendmentConn.List(ctx, &inout.AmendmentListRequest{
		Sid:      sid,
		Switch:   &switch1,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AmendmentList")
}

// TestSave 测试保存优化客流接口
func TestSave(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	switch1 := int64(1)
	inPlus := int64(10)
	outPlus := int64(-5)
	req := &inoutV5.AmendmentSaveReq{
		Gid:            gid[0],
		Switch:         &switch1,
		Type:           0, // 加减方式
		InPlus:         &inPlus,
		OutPlus:        &outPlus,
		InCoefficient:  1.0,
		OutCoefficient: 1.0,
	}
	var r inoutV5.AmendmentSaveResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := amendmentConn.Save(ctx, &inout.AmendmentSaveRequest{
		Gid:            gid[0],
		Switch:         &switch1,
		Type:           0,
		InPlus:         &inPlus,
		OutPlus:        &outPlus,
		InCoefficient:  1.0,
		OutCoefficient: 1.0,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "AmendmentSave")
}

// TestUpdate 测试更新出入园统计组接口
func TestUpdate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.GroupUpdateReq{
		Id:         gid[0],
		Name:       "更新后的测试统计组",
		Gates:      []string{"gate1", "gate2", "gate3"},
		Capacity:   1200,
		UpperLimit: 6000,
		Config: []*inoutV5.GroupConfig{
			{Label: "低客流", Min: 0, Max: 25, Color: "#00FF00"},
			{Label: "中客流", Min: 26, Max: 75, Color: "#FFFF00"},
			{Label: "高客流", Min: 76, Max: 100, Color: "#FF0000"},
		},
	}
	var r inoutV5.GroupUpdateResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := groupConn.Update(ctx, &inout.GroupUpdateRequest{
		Id:         gid[0],
		Name:       "更新后的测试统计组",
		Gates:      []string{"gate1", "gate2", "gate3"},
		Capacity:   1200,
		UpperLimit: 6000,
		Config: []*inout.GroupConfig{
			{Label: "低客流", Min: 0, Max: 25, Color: "#00FF00"},
			{Label: "中客流", Min: 26, Max: 75, Color: "#FFFF00"},
			{Label: "高客流", Min: 76, Max: 100, Color: "#FF0000"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "GroupUpdate")
}

// TestHourSummaryByDevice 测试按设备小时统计出入园数据接口
func TestHourSummaryByDevice(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	dateStr := time.Now().Format("2006-01-02")
	req := &inoutV5.HourSummaryByDeviceReq{
		Date:    dateStr,
		Sid:     sid,
		Devices: devices,
		Hour:    10,
	}
	var r inoutV5.HourSummaryByDeviceResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.HourSummaryByDevice(ctx, &inout.HourSummaryByDeviceRequest{
		Date:    dateStr,
		Sid:     sid,
		Devices: devices,
		Hour:    10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "HourSummaryByDevice")
}

// TestSummaryByDate 测试按日期维度统计出入园数据接口
func TestSummaryByDate(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &inoutV5.SummaryByDateReq{
		Request: common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Sid:   []int{int(sid)},
		},
		Gid: []int{int(gid[0])},
		Amend: &inoutV5.AmendConfig{
			NoAmend: false,
		},
	}
	var r inoutV5.SummaryByDateResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.SummaryByDate(ctx, &inout.SummaryByDateRequest{
		Request: &idsscommon.PassedTimeSpan{
			Start: timestamppb.New(startTime),
			End:   timestamppb.New(endTime),
		},
		Sid:     sid,
		Gid:     gid,
		NoAmend: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 添加调试信息，打印实际的响应数据
	fmt.Printf("=== SDK 响应数据 ===\n")
	fmt.Printf("SDK Response: %+v\n", r)
	fmt.Printf("SDK Total: In=%d, Out=%d, Hold=%d\n", r.Total.In, r.Total.Out, r.Total.Hold)
	if len(r.List) > 0 {
		fmt.Printf("SDK List[0]: Date=%s, Summary=%+v\n", r.List[0].Date, r.List[0].Summary)
		fmt.Printf("SDK List[0].Summary.Total: In=%d, Out=%d, Hold=%d\n", 
			r.List[0].Summary.Total.In, r.List[0].Summary.Total.Out, r.List[0].Summary.Total.Hold)
		if len(r.List[0].Summary.List) > 0 {
			fmt.Printf("SDK List[0].Summary.List 长度: %d\n", len(r.List[0].Summary.List))
		} else {
			fmt.Printf("SDK List[0].Summary.List 为空\n")
		}
	} else {
		fmt.Printf("SDK List 为空\n")
	}

	fmt.Printf("\n=== gRPC 响应数据 ===\n")
	fmt.Printf("gRPC Response: %+v\n", grpcResp)
	if grpcResp.Total != nil {
		fmt.Printf("gRPC Total: In=%d, Out=%d\n", grpcResp.Total.In, grpcResp.Total.Out)
	}
	if len(grpcResp.List) > 0 {
		fmt.Printf("gRPC List[0]: Time=%s, In=%d, Out=%d\n", 
			grpcResp.List[0].Time, grpcResp.List[0].In, grpcResp.List[0].Out)
	}

	// 启用调试模式查看字段解析过程
	options := &compare.ComparisonOptions{}
	resp := compare.CompareResponses(r, grpcResp, "SummaryByDate", options)
	
	// 手动创建解析器查看字段映射
	parser := compare.NewStructParser()
	parser.SetDebugMode(true)
	fmt.Printf("\n=== SDK 字段解析 ===\n")
	sdkFields := parser.ParseToFieldMap(r)
	fmt.Printf("SDK 解析出的字段数量: %d\n", len(sdkFields))
	for path, field := range sdkFields {
		fmt.Printf("SDK 字段: %s = %v (类型: %s)\n", path, field.Value, field.Type)
	}
	
	fmt.Printf("\n=== gRPC 字段解析 ===\n")
	grpcFields := parser.ParseToFieldMap(grpcResp)
	fmt.Printf("gRPC 解析出的字段数量: %d\n", len(grpcFields))
	for path, field := range grpcFields {
		fmt.Printf("gRPC 字段: %s = %v (类型: %s)\n", path, field.Value, field.Type)
	}
	fmt.Printf("\n=== 比较结果 ===\n")
	fmt.Printf("比较结果： %+v\n", resp)
	if resp.HasDiff {
		fmt.Printf("发现差异数量: %d\n", len(resp.Differences))
		for i, diff := range resp.Differences {
			fmt.Printf("差异 %d: %s - %s\n", i+1, diff.DiffType.String(), diff.Message)
		}
	} else {
		fmt.Printf("未发现差异，但测试失败，可能是空数据问题\n")
	}
	t.Fatal(resp)
}

// TestSummaryByTime 测试按时间维度统计出入园数据接口
func TestSummaryByTime(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &inoutV5.SummaryReq{
		Request: common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Sid:   []int{int(sid)},
		},
		Gid: []int{int(gid[0])},
		Amend: &inoutV5.AmendConfig{
			NoAmend: false,
		},
	}
	grpcReq := &inout.SummaryByTimeRequest{
		Request: &idsscommon.PassedTimeSpan{
			Start: timestamppb.New(startTime),
			End:   timestamppb.New(endTime),
		},
		Sid:     sid,
		Gid:     gid,
		NoAmend: false,
	}

	var apiResp inoutV5.SummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.SummaryByTime(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTime")
}

// TestTodaySummaryByGroupHour 测试今日分组小时汇总接口
func TestTodaySummaryByGroupHour(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &inoutV5.SummaryByGroupReq{
		Gid: []int{int(gid[0])},
		Amend: &inoutV5.AmendConfig{
			NoAmend: false,
		},
	}
	var r inoutV5.SummaryByGroupResp
	err := iam.Do(req, &r, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := statConn.TodaySummaryByGroupHour(ctx, &inout.TodaySummaryByGroupHourRequest{
		Gid:     gid,
		NoAmend: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &r, grpcResp, "TodaySummaryByGroupHour")

}
