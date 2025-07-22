package terminal

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/piaofutong/odas-sdk/odas"
	terminalV5 "github.com/piaofutong/odas-sdk/odas/terminal"
	"github.com/piaofutong/odas-sdk/test/utils"
	idsssdk "gitlab.12301.test/gopkg/idss-go-sdk"
	"gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss/terminal"
	"google.golang.org/grpc"
)

// 测试配置变量
var (
	accessId  = "abcdefg"
	accessKey = "abcdefg"
	token     = "558c0eba2e29b81091c159e7cbf931f6c22cc8db"
	sid       = []int64{3385}
	lid       = []int64{1}
)

var (
	ctx               context.Context
	idssConn          *grpc.ClientConn
	devicesConn       terminal.DevicesServiceClient
	deviceJournalConn terminal.DeviceJournalServiceClient
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
	devicesConn = terminal.NewDevicesServiceClient(idssConn)
	deviceJournalConn = terminal.NewDeviceJournalServiceClient(idssConn)

	os.Exit(m.Run())
}

// TestList 测试设备列表接口
func TestList(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &terminalV5.ListReq{
		Sids:       sid,
		Page:       1,
		PageSize:   10,
		DeviceType: 1,
	}
	var apiResp terminalV5.ListResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := devicesConn.List(ctx, &terminal.ListRequest{
		Sids:       sid,
		Page:       1,
		PageSize:   10,
		DeviceType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "DevicesList")
}

// TestDeviceJournalStat 测试设备日志统计接口
func TestDeviceJournalStat(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &terminalV5.DeviceJournalStatReq{
		Sid:    sid[0],
		NodeId: []int32{1, 2},
		Hour:   []int32{9, 10, 11},
	}
	var apiResp terminalV5.DeviceJournalStatResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := deviceJournalConn.Stat(ctx, &terminal.DeviceJournalStatRequest{
		Sid:    sid[0],
		NodeId: []int32{1, 2},
		Hour:   []int32{9, 10, 11},
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "DeviceJournalStat")
}

// TestDeviceSummaryByCategory 测试设备分类汇总接口
func TestDeviceSummaryByCategory(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &terminalV5.DeviceSummaryByCategoryReq{
		Sids: sid,
	}
	var apiResp terminalV5.DeviceSummaryByCategoryResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := devicesConn.DeviceSummaryByCategory(ctx, &terminal.DeviceSummaryByCategoryRequest{
		Sids: sid,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "DeviceSummaryByCategory")
}

// TestDeviceSummaryByNetStatus 测试设备网络状态汇总接口
func TestDeviceSummaryByNetStatus(t *testing.T) {
	iam := odas.NewIAM(accessId, accessKey)
	req := &terminalV5.DeviceSummaryByNetStatusReq{
		Sids: sid,
	}
	var apiResp terminalV5.DeviceSummaryByNetStatusResp
	err := iam.Do(req, &apiResp, odas.WithToken(token))
	if err != nil {
		t.Fatal(err)
	}

	grpcResp, err := devicesConn.DeviceSummaryByNetStatus(ctx, &terminal.DeviceSummaryByNetStatusRequest{
		Sids: sid,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 直接比较两个响应结构体
	utils.CompareResponses(t, &apiResp, grpcResp, "DeviceSummaryByNetStatus")
}
