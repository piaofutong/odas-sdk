package basic_analyze_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas"
	basicAnalyzeV5 "github.com/piaofutong/odas-sdk/odas/basic_analyze"
	"github.com/piaofutong/odas-sdk/odas/common"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/pb"
	"gitlab.12301.test/gopkg/idss-go-sdk/pb/basic_analyze"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 测试配置变量
var (
	accessId  = "abcdefg"
	accessKey = "abcdefg"
	startTime = time.Now().AddDate(-1, 0, 0)
	endTime   = time.Now()
	token     = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	orderType = common.OrderType(1)
	sid       = []int64{3385}
	lid       = []int64{}
	ticketId  = []int64{}
	// lid       = []int64{12, 13}
	// ticketId = []int64{100}
)

// gRPC 客户端连接变量
var (
	ctx                   context.Context
	idssConn              *grpc.ClientConn
	channelConn           basic_analyze.ChannelServiceClient
	orderConn             basic_analyze.OrderServiceClient
	productConn           basic_analyze.ProductServiceClient
	reportConn            basic_analyze.ReportServiceClient
	timeSpanTicketingConn basic_analyze.TimeSpanTicketingServiceClient
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
	channelConn = basic_analyze.NewChannelServiceClient(idssConn)
	orderConn = basic_analyze.NewOrderServiceClient(idssConn)
	productConn = basic_analyze.NewProductServiceClient(idssConn)
	reportConn = basic_analyze.NewReportServiceClient(idssConn)
	timeSpanTicketingConn = basic_analyze.NewTimeSpanTicketingServiceClient(idssConn)

	os.Exit(m.Run())
}

// TestOrderStatistics 测试订单统计接口
func TestOrderStatistics(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.OrderStatisticsReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   sid,
		},
		YearOnYear:     true,
		PeriodOnPeriod: true,
		Lid:            lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.OrderStatisticsResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcReq := &basic_analyze.OrderStatisticsRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		YearOnYear:     true,
		PeriodOnPeriod: true,
		Lid:            lid,
	}

	grpcResp, err := reportConn.OrderStatistics(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "OrderStatistics")
}

// TestOrderSummary 测试按年月日订单汇总接口
func TestOrderSummary2(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.OrderSummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Limit:     10,
		OrderType: common.OrderType(orderType),
		Lid:       lid,
	}
	grpcReq := &basic_analyze.OrderSummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      sid,
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Limit:     10,
		OrderType: idsscommon.OrderType(orderType),
		Lid:       lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.OrderSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := reportConn.OrderSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "OrderSummaryByYMD")
}

// TestOrderSummaryByProduct 测试按产品订单汇总接口
func TestOrderSummaryByProduct(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.OrderSummaryByProductReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   sid,
		},
		Limit:     10,
		Lid:       lid,
		OrderType: (orderType),
	}
	grpcReq := &basic_analyze.OrderSummaryByProductRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      sid,
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Lid:       lid,
		Limit:     10,
		OrderType: idsscommon.OrderType(orderType),
	}

	// API 调用
	var apiResp basicAnalyzeV5.OrderSummaryByProductResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := reportConn.OrderSummaryByProduct(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "OrderSummaryByProduct")
}

// TestOrderSummaryByTicket 测试按票订单汇总接口
func TestOrderSummaryByTicket(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.OrderSummaryByTicketReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Limit:     10,
		TicketId:  ticketId,
		Lid:       lid,
		OrderType: orderType,
	}
	grpcReq := &basic_analyze.OrderSummaryByTicketRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Limit:     10,
		OrderType: idsscommon.OrderType(orderType),
		TicketId:  ticketId,
		Lid:       lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.OrderSummaryByTicketResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := reportConn.OrderSummaryByTicket(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "OrderSummaryByTicket")
}

// // TestAnnualCardSummary 测试年卡汇总接口
// func TestAnnualCardSummary(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.AnnualCardSummaryReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   []int64{int64(sid[0])},
// 		},
// 		Limit: 10,
// 		Lid:   lid,
// 	}
// 	grpcReq := &basic_analyze.AnnualCardSummaryRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Lid:   lid,
// 		Limit: 10,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.AnnualCardSummaryResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := reportConn.AnnualCardSummary(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "AnnualCardSummary")
// }

// // TestDistributorSummary 测试分销商汇总接口
// func TestDistributorSummary(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.DistributorSummaryReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   sid,
// 		},
// 		Page:     1,
// 		PageSize: 10,
// 	}
// 	grpcReq := &basic_analyze.DistributorSummaryRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Page:     1,
// 		PageSize: 10,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.DistributorSummaryResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := reportConn.DistributorSummary(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "DistributorSummary")
// }

// // TestTerminalSummary 测试终端汇总接口
// func TestTerminalSummary(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.TerminalSummaryReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   []int64{int64(sid[0])},
// 		},
// 		Limit: 10,
// 		Lid:   lid,
// 	}
// 	grpcReq := &basic_analyze.TerminalSummaryRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Lid:   lid,
// 		Limit: 10,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.TerminalSummaryResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := reportConn.TerminalSummary(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "TerminalSummary")
// }

// TestSummary 测试通用汇总接口
func TestOrderSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.OrderSummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		// Compare:   true, // TODO: 检查为什么缺少字段
		Lid:       lid,
		OrderType: orderType,
	}
	grpcReq := &basic_analyze.OrderSummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		OrderType: idsscommon.OrderType(orderType),
		Lid:       lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.OrderSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := reportConn.OrderSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "OrderSummary")
}

// // TestSummaryByHour 测试小时维度票数据统计接口
// func TestSummaryByHour(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.SummaryByHourReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   []int64{int64(sid[0])},
// 		},
// 		Lid: lid,
// 	}
// 	grpcReq := &basic_analyze.SummaryByHourRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Lid: lid,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.SummaryByHourResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := productConn.SummaryByHour(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByHour")
// }

// TestSummaryByTicket 测试票维度票数据列表统计接口
func TestSummaryByTicket(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByTicketReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Page:      1,
		PageSize:  10,
		Lid:       lid,
		OrderType: orderType,
	}
	grpcReq := &basic_analyze.SummaryByTicketRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		OrderType: idsscommon.OrderType(orderType),
		Lid:       lid,
		Page:      1,
		PageSize:  10,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByTicketResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := productConn.SummaryByTicket(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTicket")
}

// TestSummaryByLevel1 测试分销渠道数据统计接口
func TestSummaryByLevel1(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByLevel1Req{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Limit:     10,
		OrderType: common.OrderType(orderType),
		Lid:       lid,
	}
	grpcReq := &basic_analyze.SummaryByLevel1Request{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Limit:     10,
		OrderType: idsscommon.OrderType(orderType),
		Lid:       lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByLevel1Resp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := channelConn.SummaryByLevel1(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByLevel1")
}

// TestSummaryByLevel2 测试二级渠道订单数据接口
func TestSummaryByLevel2(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByLevel2Req{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Limit:           10,
		OrderType:       common.OrderType(orderType),
		ChannelLevel1Id: 1,
		Lid:             lid,
	}
	grpcReq := &basic_analyze.SummaryByLevel2Request{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Limit:           10,
		OrderType:       idsscommon.OrderType(orderType),
		Lid:             lid,
		ChannelLevel1Id: 1,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByLevel2Resp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := channelConn.SummaryByLevel2(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByLevel2")
}

// TestSummaryByLevel1AndLevel1Name 测试分销渠道数据统计（包含渠道名称）接口
func TestSummaryByLevel1AndLevel1Name(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByLevel1AndLevel1NameReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Limit:     10,
		Lid:       lid,
		OrderType: common.OrderType(orderType),
	}
	grpcReq := &basic_analyze.SummaryByLevel1AndLevel1NameRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Limit:     10,
		OrderType: idsscommon.OrderType(orderType),
		Lid:       lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByLevel1AndLevel1NameResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := channelConn.SummaryByLevel1AndLevel1Name(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByLevel1AndLevel1Name")
}

// TestToiSummary 测试TOI汇总接口
func TestToiSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.ToiSummaryReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Lid:       lid,
		OrderType: common.OrderType(orderType),
	}
	grpcReq := &basic_analyze.ToiSummaryRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Lid:       lid,
		OrderType: idsscommon.OrderType(orderType),
	}

	// API 调用
	var apiResp basicAnalyzeV5.ToiSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := orderConn.ToiSummary(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "ToiSummary")
}

// TestToiStatistics 测试TOI统计接口
func TestToiStatistics(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	req := &basicAnalyzeV5.ToiStatisticsReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		Lid: lid,
	}
	grpcReq := &basic_analyze.ToiStatisticsRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		Lid:       lid,
		OrderType: idsscommon.OrderType(orderType),
	}

	// API 调用
	var apiResp basicAnalyzeV5.ToiStatisticsResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := orderConn.ToiStatistics(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "ToiStatistics")
}

// // TestRefundSummaryByLevel2 测试二级渠道退款汇总接口
// func TestRefundSummaryByLevel2(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.RefundSummaryByLevel2Req{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   []int64{int64(sid[0])},
// 		},
// 		Limit:     10,
// 		Lid:       lid,
// 		OrderType: common.OrderType(orderType),
// 	}
// 	grpcReq := &basic_analyze.RefundSummaryByLevel2Request{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Limit:     10,
// 		OrderType: idsscommon.OrderType(orderType),
// 		Lid:       lid,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.RefundSummaryByLevel2Resp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := channelConn.RefundSummaryByLevel2(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "RefundSummaryByLevel2")
// }

// // TestSummaryByLevel2AndTicket 测试二级渠道票数据统计接口
// func TestSummaryByLevel2AndTicket(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.SummaryByLevel2AndTicketReq{
// 		Request: &common.PassedTimeSpanRequest{
// 			Start: startTime,
// 			End:   endTime,
// 			Type:  0,
// 			Sid:   []int64{int64(sid[0])},
// 		},
// 		Lid:             lid,
// 		Limit:           10,
// 		OrderType:       common.OrderType(orderType),
// 		ChannelLevel1Id: 1,
// 	}
// 	grpcReq := &basic_analyze.SummaryByLevel2AndTicketRequest{
// 		Request: &idsscommon.PassedTimeSpanRequest{
// 			Sid:      []int64{int64(sid[0])},
// 			Start:    timestamppb.New(startTime),
// 			End:      timestamppb.New(endTime),
// 			DateType: idsscommon.DateType_DAILY,
// 		},
// 		Limit:           10,
// 		ChannelLevel1Id: 1,
// 		OrderType:       idsscommon.OrderType(orderType),
// 		Lid:             lid,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.SummaryByLevel2AndTicketResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := channelConn.SummaryByLevel2AndTicket(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(grpcResp.List) == 0 {
// 		t.Fatalf("测试失败，数据返回为空")
// 	}

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByLevel2AndTicket")
// }

// TestSummaryByTicketAndChannel 测试票渠道数据统计接口
func TestSummaryByTicketAndChannel(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByTicketAndChannelReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		TicketId: ticketId,
		Lid:      lid,
	}
	grpcReq := &basic_analyze.SummaryByTicketAndChannelRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		TicketId: ticketId,
		Lid:      lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByTicketAndChannelResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := productConn.SummaryByTicketAndChannel(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTicketAndChannel")
}

// TestSummaryByTicketAndDay 测试票日期数据统计接口
func TestSummaryByTicketAndDay(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &basicAnalyzeV5.SummaryByTicketAndDayReq{
		Request: &common.PassedTimeSpanRequest{
			Start: startTime,
			End:   endTime,
			Type:  0,
			Sid:   []int64{int64(sid[0])},
		},
		TicketId: ticketId,
		Lid:      lid,
	}
	grpcReq := &basic_analyze.SummaryByTicketAndDayRequest{
		Request: &idsscommon.PassedTimeSpanRequest{
			Sid:      []int64{int64(sid[0])},
			Start:    timestamppb.New(startTime),
			End:      timestamppb.New(endTime),
			DateType: idsscommon.DateType_DAILY,
		},
		TicketId: ticketId,
		Lid:      lid,
	}

	// API 调用
	var apiResp basicAnalyzeV5.SummaryByTicketAndDayResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	// gRPC 调用
	grpcResp, err := productConn.SummaryByTicketAndDay(ctx, grpcReq)
	if err != nil {
		t.Fatal(err)
	}

	if len(grpcResp.List) == 0 {
		t.Fatalf("测试失败，数据返回为空")
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "SummaryByTicketAndDay")
}

// // TestTodayTicketingDetail 测试今日出票明细接口
// func TestTodayTicketingDetail(t *testing.T) {
// 	iam := odas.NewIAM(accessId, accessKey)
// 	req := &basicAnalyzeV5.TodayTicketingReq{
// 		Lid: lid,
// 	}
// 	grpcReq := &basic_analyze.TodayTicketingRequest{
// 		Lid: lid,
// 	}

// 	// API 调用
// 	var apiResp basicAnalyzeV5.TodayTicketingResp
// 	err := iam.Do(req, &apiResp, odas.WithToken(token))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// gRPC 调用
// 	grpcResp, err := timeSpanTicketingConn.TodayTicketingDetail(ctx, grpcReq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// TODO: 测试数据为空的情况
// 	// if len(grpcResp.List) == 0 {
// 	// 	t.Fatalf("测试失败，数据返回为空")
// 	// }

// 	// 直接比较两个响应结构体
// 	utils.CompareResponses(t, &apiResp, grpcResp, "TodayTicketing")
// }
