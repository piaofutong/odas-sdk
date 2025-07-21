package tourist_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas"
	"github.com/piaofutong/odas-sdk/odas/request/common"
	"github.com/piaofutong/odas-sdk/odas/request/types"
	touristV5 "github.com/piaofutong/odas-sdk/odas/request/v5/tourist"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss"
	"gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss/tourist"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 测试配置变量
var (
	accessId   = "abcdefg"
	accessKey  = "abcdefg"
	token      = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	sid        = []int64{3385}
	lid        = []int64{116157, 116155}
	excludeLid = []int64{116157, 116156}
	timeType   = int64(0)
	orderType  = int64(0)
)

// gRPC 客户端连接变量
var (
	ctx                  context.Context
	idssConn             *grpc.ClientConn
	flowConn             tourist.FlowServiceClient
	portraitConn         tourist.PortraitServiceClient
	preBookingConn       tourist.PreBookingServiceClient
	temporaryBookingConn tourist.TemporaryBookingServiceClient
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
	flowConn = tourist.NewFlowServiceClient(idssConn)
	portraitConn = tourist.NewPortraitServiceClient(idssConn)
	preBookingConn = tourist.NewPreBookingServiceClient(idssConn)
	temporaryBookingConn = tourist.NewTemporaryBookingServiceClient(idssConn)

	os.Exit(m.Run())
}

// TestForecastTouristSummary 测试预测游客汇总接口
func TestForecastTouristSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.ForecastTouristSummaryReq{
		Request: common.PassedTimeSpanByOrderTypeV4Request{
			Start:      startTime,
			End:        endTime,
			Type:       0,
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  0,
		},
	}
	grpcReq := &tourist.ForecastTouristSummaryRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
			OrderType:  idsscommon.OrderType_MAIN,
		},
	}

	// converter := utils.NewStructConverter()
	// converter.ConvertStruct(req, grpcReq)
	// // 直接比较两个请求结构体
	// utils.CompareResponses(t, &req, grpcReq, "ForecastTouristSummaryReq")

	// API 调用
	var apiResp touristV5.ForecastTouristSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := flowConn.ForecastTouristSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "ForecastTouristSummary")

	t.Logf("响应 List: ListCount=%d", len(apiResp.List))
}

// TestInsideAndOutsideByProvince 测试省内外游客接口
func TestInsideAndOutsideByProvince(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.InsideAndOutsideByProvinceReq{
		Request: common.PassedTimeSpanByOrderTypeV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			OrderType: orderType,
		},
		DimensionType: 1,
		Province:      "福建省",
	}
	grpcReq := &tourist.InsideAndOutsideByProvinceRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:       sid,
			Start:     timestamppb.New(startTime),
			End:       timestamppb.New(endTime),
			DateType:  idsscommon.DateType_DAILY,
			OrderType: idsscommon.OrderType_MAIN,
		},
		DimensionType: 1,
		Province:      "福建省",
	}

	// API 调用
	var apiResp touristV5.InsideAndOutsideByProvinceResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.InsideAndOutsideByProvince(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "InsideAndOutsideByProvince")

	t.Logf("省内外游客统计结果: Total=%+v, InsideCount=%d, OutsideCount=%d", apiResp.Total, len(apiResp.Inside), len(apiResp.Outside))
}

// TestLastDayTemporaryBookingSummary 测试昨日临时预约汇总接口
func TestLastDayTemporaryBookingSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &touristV5.TemporaryBookingSummaryReq{
		Sid: sid[0],
	}
	grpcReq := &tourist.TemporaryBookingSummaryRequest{
		Sid: sid[0],
	}

	// API 调用
	var apiResp touristV5.TemporaryBookingSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := temporaryBookingConn.LastDayTemporaryBookingSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "LastDayTemporaryBookingSummary")

	t.Logf("昨日临时预约统计结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestPreBookingSummary 测试预约汇总接口
func TestPreBookingSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.PreBookingServiceSummaryReq{
		Request: &common.PassedTimeSpanByOrderTypeV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:       lid,
			OrderType: orderType,
		},
	}
	grpcReq := &tourist.PreBookingServiceSummaryRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:       sid,
			Lid:       lid,
			Start:     timestamppb.New(startTime),
			End:       timestamppb.New(endTime),
			DateType:  idsscommon.DateType_DAILY,
			OrderType: idsscommon.OrderType_MAIN,
		},
	}

	// API 调用
	var apiResp touristV5.PreBookingServiceSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := preBookingConn.Summary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "PreBookingSummary")

	t.Logf("预订汇总结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestSexSummaryByAge 测试年龄性别汇总接口
func TestSexSummaryByAge(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.SexSummaryByAgeReq{
		Request: common.PassedTimeSpanByOrderTypeV5Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
		DimensionType: 1,
	}
	grpcReq := &tourist.SexSummaryByAgeRequest{
		Request: &idsscommon.PassedTimeSpanV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
		},
		DimensionType: 1,
	}

	// API 调用
	var apiResp touristV5.SexSummaryByAgeResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.SexSummaryByAge(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SexSummaryByAge")

	t.Logf("按年龄性别统计结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestSummaryByType 测试按类型统计接口
func TestSummaryByType(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.SummaryByTypeReq{
		Request: &common.PassedTimeSpanByOrderTypeV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
	}
	grpcReq := &tourist.SummaryByTypeRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
			OrderType:  idsscommon.OrderType_MAIN,
		},
	}

	// API 调用
	var apiResp touristV5.SummaryByTypeResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := preBookingConn.SummaryByType(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByType")

	t.Logf("按类型统计结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestTicketSummaryByPayChannel 测试按支付渠道统计门票分布接口
func TestTicketSummaryByPayChannel(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TicketSummaryByPayChannelReq{
		Request: common.PassedTimeSpanByOrderTypeV5Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
	}
	grpcReq := &tourist.TicketSummaryByPayChannelRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
			OrderType:  idsscommon.OrderType_MAIN,
		},
	}

	// API 调用
	var apiResp touristV5.TicketSummaryByPayChannelResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TicketSummaryByPayChannel(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TicketSummaryByPayChannel")

	t.Logf("按支付渠道统计门票分布结果: ListCount=%d", len(apiResp.List))
}

// TestTouristLocalByTicket 测试按门票统计本地游客接口
func TestTouristLocalByTicket(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryByLocalReq{
		Request: common.PassedTimeSpanByOrderTypeV5Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
		RegionType: 1,
		Province:   "福建省",
		City:       "厦门市",
		Limit:      10,
		Unknown:    true,
	}
	grpcReq := &tourist.TouristSummaryByLocalRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
			OrderType:  idsscommon.OrderType_MAIN,
		},
		RegionType: 1,
		Province:   "福建省",
		City:       "厦门市",
		Limit:      10,
		Unknown:    true,
	}

	// API 调用
	var apiResp touristV5.TouristSummaryByLocalResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TouristLocalByTicket(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummaryByLocal")

	t.Logf("按门票统计本地游客结果: ListCount=%d", len(apiResp.List))
}

// TestTouristSummary 测试游客统计接口
func TestTouristSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryReq{
		Request: common.PassedTimeSpanByOrderTypeV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			OrderType: orderType,
		},
	}
	grpcReq := &tourist.TouristSummaryRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:       sid,
			Start:     timestamppb.New(startTime),
			End:       timestamppb.New(endTime),
			DateType:  idsscommon.DateType_DAILY,
			OrderType: idsscommon.OrderType_MAIN,
		},
	}

	// API 调用
	var apiResp touristV5.TouristSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := flowConn.TouristSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummary")

	t.Logf("游客统计结果: ListCount=%d", len(apiResp.List))
}

// TestTouristSummaryByCity 测试按城市统计游客分布接口
func TestTouristSummaryByCity(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryByCityReq{
		Request: common.PassedTimeSpanV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
		},
		OrderType:     types.OrderType(orderType),
		DimensionType: 1,
		Province:      "福建省",
		Limit:         10,
	}
	grpcReq := &tourist.TouristSummaryByCityRequest{
		Request: &idsscommon.PassedTimeSpanV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
		},
		OrderType:     idsscommon.OrderType_MAIN,
		DimensionType: 1,
		Province:      "福建省",
		Limit:         10,
	}

	// API 调用
	var apiResp touristV5.TouristSummaryByCityResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TouristSummaryByCity(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummaryByCity")

	t.Logf("按城市统计游客分布结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestTouristSummaryByDistrict 测试按区县统计游客分布接口
func TestTouristSummaryByDistrict(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryByDistrictReq{
		Request: common.PassedTimeSpanV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
		},
		OrderType:     types.OrderType(orderType),
		DimensionType: 1,
		Province:      "福建省",
		City:          "厦门市",
		Limit:         10,
	}
	grpcReq := &tourist.TouristSummaryByDistrictRequest{
		Request: &idsscommon.PassedTimeSpanV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
		},
		OrderType:     idsscommon.OrderType_MAIN,
		DimensionType: 1,
		Province:      "福建省",
		City:          "厦门市",
		Limit:         10,
	}

	// API 调用
	var apiResp touristV5.TouristSummaryByDistrictResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TouristSummaryByDistrict(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummaryByDistrict")

	t.Logf("按区县统计游客分布结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestTouristSummaryByPeer 测试按同行人数统计游客分布接口
func TestTouristSummaryByPeer(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryByPeerReq{
		Request: common.PassedTimeSpanByOrderTypeV5Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
		Province: "福建省",
	}
	grpcReq := &tourist.TouristSummaryByPeerRequest{
		Request: &idsscommon.PassedTimeSpanByOrderTypeV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
			OrderType:  idsscommon.OrderType_MAIN,
		},
		Province: "福建省",
	}

	// API 调用
	var apiResp touristV5.TouristSummaryByPeerResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TouristSummaryByPeer(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummaryByPeer")

	t.Logf("按同行人数统计游客分布结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}

// TestTouristSummaryByProvince 测试按省份统计游客分布接口
func TestTouristSummaryByProvince(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &touristV5.TouristSummaryByProvinceReq{
		Request: common.PassedTimeSpanByOrderTypeV4Request{
			Start: startTime,
			End:   endTime,
			Type:  timeType, Sid: sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			OrderType:  orderType,
		},
		DimensionType: 1,
		Limit:         10,
	}
	grpcReq := &tourist.TouristSummaryByProvinceRequest{
		Request: &idsscommon.PassedTimeSpanV4Request{
			Sid:        sid,
			Lid:        lid,
			ExcludeLid: excludeLid,
			Start:      timestamppb.New(startTime),
			End:        timestamppb.New(endTime),
			DateType:   idsscommon.DateType_DAILY,
		},
		OrderType:     idsscommon.OrderType_MAIN,
		DimensionType: 1,
		Limit:         10,
	}

	// API 调用
	var apiResp touristV5.TouristSummaryByProvinceResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := portraitConn.TouristSummaryByProvince(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "TouristSummaryByProvince")

	t.Logf("按省份统计游客分布结果: Total=%+v, ListCount=%d", apiResp.Total, len(apiResp.List))
}
