package s3rd

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas"
	s3rdV5 "github.com/piaofutong/odas-sdk/odas/s3rd"
	"github.com/piaofutong/odas-sdk/odas/types"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	"gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss/s3rd"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 测试配置变量
var (
	accessId     = "abcdefg"
	accessKey    = "abcdefg"
	token        = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	sid          = []int64{3385}
	lid          = []int64{1}
	start        = types.Time(time.Now().AddDate(0, 0, -7))
	end          = types.Time(time.Now())
	date         = types.Time(time.Now())
	compare      = &types.Time{}
	province     = "北京"
	unknown      = false
	limit        = int64(10)
	codeCategory = []string{"category1"}
)

var (
	ctx            context.Context
	idssConn       *grpc.ClientConn
	parkingConn    s3rd.ParkingServiceClient
	lvyunHotelConn s3rd.LvyunHotelServiceClient
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
	parkingConn = s3rd.NewParkingServiceClient(idssConn)
	lvyunHotelConn = s3rd.NewLvyunHotelServiceClient(idssConn)

	os.Exit(m.Run())
}

// TestInoutSummaryByHour 测试按小时进出汇总接口
func TestInoutSummaryByHour(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.InoutSummaryByHourReq{
		Sid:     sid,
		Date:    date,
		Compare: compare,
	}
	var apiResp s3rdV5.InoutSummaryByHourResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	var compareTime *timestamppb.Timestamp
	if compare != nil && !time.Time(*compare).IsZero() {
		compareTime = timestamppb.New(time.Time(*compare))
	}

	grpcResp, err := parkingConn.InoutSummaryByHour(ctx, &s3rd.InoutSummaryByHourRequest{
		Sid:     sid,
		Date:    timestamppb.New(time.Time(date)),
		Compare: compareTime,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "InoutSummaryByHour")
}

// TestLocationInSummary 测试位置进入汇总接口
func TestLocationInSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.LocationInSummaryReq{
		Sid:      sid,
		Start:    start,
		End:      end,
		Province: province,
		Unknown:  unknown,
		Limit:    limit,
	}
	var apiResp s3rdV5.LocationInSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := parkingConn.LocationInSummary(ctx, &s3rd.LocationInSummaryRequest{
		Sid:      sid,
		Start:    timestamppb.New(time.Time(start)),
		End:      timestamppb.New(time.Time(end)),
		Province: province,
		Unknown:  unknown,
		Limit:    limit,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "LocationInSummary")
}

// TestOccupancy 测试酒店入住率查询接口
func TestOccupancy(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.OccupancyReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}
	var apiResp s3rdV5.OccupancyResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := lvyunHotelConn.Occupancy(ctx, &s3rd.OccupancyRequest{
		Sid:   sid,
		Start: timestamppb.New(time.Time(start)),
		End:   timestamppb.New(time.Time(end)),
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "Occupancy")
}

// TestRevenueSummaryByCodeCategory 测试按代码分类收入汇总接口
func TestRevenueSummaryByCodeCategory(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.RevenueSummaryByCodeCategoryReq{
		Sid:          sid,
		Start:        &start,
		End:          &end,
		CodeCategory: codeCategory,
	}
	var apiResp s3rdV5.RevenueSummaryByCodeCategoryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	startTime := timestamppb.New(time.Time(start))
	endTime := timestamppb.New(time.Time(end))
	grpcResp, err := lvyunHotelConn.RevenueSummaryByCodeCategory(ctx, &s3rd.RevenueSummaryByCodeCategoryRequest{
		Sid:          sid,
		Start:        startTime,
		End:          endTime,
		CodeCategory: codeCategory,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "RevenueSummaryByCodeCategory")
}

// TestRmOrderSummary 测试客房订单汇总接口
func TestRmOrderSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.RmOrderSummaryReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}
	var apiResp s3rdV5.RmOrderSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := lvyunHotelConn.RmOrderSummary(ctx, &s3rd.RmOrderSummaryRequest{
		Sid:   sid,
		Start: timestamppb.New(time.Time(start)),
		End:   timestamppb.New(time.Time(end)),
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "RmOrderSummary")
}

// TestRmSaleSummary 测试客房销售汇总接口
func TestRmSaleSummary(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.RmSaleSummaryReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}
	var apiResp s3rdV5.RmSaleSummaryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := lvyunHotelConn.RmSaleSummary(ctx, &s3rd.RmSaleSummaryRequest{
		Sid:   sid,
		Start: timestamppb.New(time.Time(start)),
		End:   timestamppb.New(time.Time(end)),
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "RmSaleSummary")
}

// TestRmSaleSummaryByBind 测试按绑定客房销售汇总接口
func TestRmSaleSummaryByBind(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.RmSaleSummaryByBindReq{
		Sid:   sid,
		Start: start,
		End:   end,
	}
	var apiResp s3rdV5.RmSaleSummaryByBindResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := lvyunHotelConn.RmSaleSummaryByBind(ctx, &s3rd.RmSaleSummaryByBindRequest{
		Sid:   sid,
		Start: timestamppb.New(time.Time(start)),
		End:   timestamppb.New(time.Time(end)),
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "RmSaleSummaryByBind")
}

// TestSpace 测试停车位查询接口
func TestSpace(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &s3rdV5.SpaceReq{
		Sid: sid,
	}
	var apiResp s3rdV5.SpaceResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := parkingConn.Space(ctx, &s3rd.SpaceRequest{
		Sid: sid,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "Space")
}
