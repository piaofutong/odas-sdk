package utils

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/piaofutong/odas-sdk/odas/common"
	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestStructConverter 测试结构体转换器
func TestStructConverter(t *testing.T) {
	// 测试基本的结构体转换
	type SourceStruct struct {
		ID   int64
		Name string
	}

	type TargetStruct struct {
		ID   int64
		Name string
	}

	converter := NewStructConverter()
	src := SourceStruct{ID: 123, Name: "test"}
	var dst TargetStruct

	err := converter.ConvertStruct(src, &dst)
	if err != nil {
		t.Fatalf("ConvertStruct failed: %v", err)
	}

	if dst.ID != src.ID || dst.Name != src.Name {
		t.Errorf("Conversion failed: expected %+v, got %+v", src, dst)
	}
}

// TestTimeConversion 测试时间转换
func TestTimeConversion(t *testing.T) {
	type SourceStruct struct {
		Time time.Time
	}

	type TargetStruct struct {
		Time *timestamppb.Timestamp
	}

	converter := NewStructConverter()
	now := time.Now()
	src := SourceStruct{Time: now}
	var dst TargetStruct

	err := converter.ConvertStruct(src, &dst)
	if err != nil {
		t.Fatalf("ConvertStruct failed: %v", err)
	}

	if dst.Time == nil {
		t.Error("Time conversion failed: timestamp is nil")
	}

	// 验证时间是否正确转换
	convertedTime := dst.Time.AsTime()
	if !convertedTime.Equal(now) {
		t.Errorf("Time conversion failed: expected %v, got %v", now, convertedTime)
	}
}

// TestEnumConversion 测试枚举转换
func TestEnumConversion(t *testing.T) {
	type SourceStruct struct {
		Type int64
	}

	type TargetStruct struct {
		Type idsscommon.DateType
	}

	converter := NewStructConverter()
	src := SourceStruct{Type: 1}
	var dst TargetStruct

	err := converter.ConvertStruct(src, &dst)
	if err != nil {
		t.Fatalf("ConvertStruct failed: %v", err)
	}

	if dst.Type != idsscommon.DateType_DAILY {
		t.Errorf("Enum conversion failed: expected %v, got %v", idsscommon.DateType_DAILY, dst.Type)
	}
}

// TestComplexStructConversion 测试复杂结构体转换
func TestComplexStructConversion(t *testing.T) {
	// 模拟真实的API请求结构转换
	type APIRequest struct {
		Request *common.PassedTimeSpanByOrderTypeV4Request
	}

	type GRPCRequest struct {
		Request *idsscommon.PassedTimeSpanByOrderTypeV4Request
	}

	converter := NewStructConverter()
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()

	apiReq := &APIRequest{
		Request: &common.PassedTimeSpanByOrderTypeV4Request{
			Start:      startTime,
			End:        endTime,
			Type:       1,
			Sid:        []int64{3385},
			Lid:        []int64{116157, 116155},
			ExcludeLid: []int64{116157, 116156},
			OrderType:  1,
		},
	}

	var grpcReq GRPCRequest
	err := converter.ConvertStruct(apiReq, &grpcReq)
	if err != nil {
		t.Fatalf("ConvertStruct failed: %v", err)
	}

	// 验证转换结果
	if grpcReq.Request == nil {
		t.Fatal("Request field is nil")
	}

	// 验证基本字段
	if len(grpcReq.Request.Sid) != len(apiReq.Request.Sid) {
		t.Errorf("Sid conversion failed: expected %v, got %v", apiReq.Request.Sid, grpcReq.Request.Sid)
	}

	if len(grpcReq.Request.Lid) != len(apiReq.Request.Lid) {
		t.Errorf("Lid conversion failed: expected %v, got %v", apiReq.Request.Lid, grpcReq.Request.Lid)
	}

	// 验证时间转换
	if grpcReq.Request.Start == nil {
		t.Error("Start time conversion failed: timestamp is nil")
	} else {
		convertedStart := grpcReq.Request.Start.AsTime()
		if !convertedStart.Equal(startTime) {
			t.Errorf("Start time conversion failed: expected %v, got %v", startTime, convertedStart)
		}
	}

	if grpcReq.Request.End == nil {
		t.Error("End time conversion failed: timestamp is nil")
	} else {
		convertedEnd := grpcReq.Request.End.AsTime()
		if !convertedEnd.Equal(endTime) {
			t.Errorf("End time conversion failed: expected %v, got %v", endTime, convertedEnd)
		}
	}

	// 验证枚举转换
	if grpcReq.Request.DateType != idsscommon.DateType_DAILY {
		t.Errorf("DateType conversion failed: expected %v, got %v", idsscommon.DateType_DAILY, grpcReq.Request.DateType)
	}

	if grpcReq.Request.OrderType != idsscommon.OrderType_MAIN {
		t.Errorf("OrderType conversion failed: expected %v, got %v", idsscommon.OrderType_MAIN, grpcReq.Request.OrderType)
	}
}

// TestGlobalConvertFunction 测试全局转换函数
func TestGlobalConvertFunction(t *testing.T) {
	type SourceStruct struct {
		ID   int64
		Name string
	}

	type TargetStruct struct {
		ID   int64
		Name string
	}

	src := SourceStruct{ID: 456, Name: "global test"}
	var dst TargetStruct

	err := ConvertStruct(src, &dst)
	if err != nil {
		t.Fatalf("ConvertStruct failed: %v", err)
	}

	if dst.ID != src.ID || dst.Name != src.Name {
		t.Errorf("Global conversion failed: expected %+v, got %+v", src, dst)
	}
}

// TestFieldDifference 测试字段差异检测
func TestFieldDifference(t *testing.T) {
	// 测试缺失字段的情况
	t.Run("MissingFields", func(t *testing.T) {
		type SourceStruct struct {
			Name string
		}

		type TargetStruct struct {
			Name string
			Age  int
		}

		converter := NewStructConverter()
		converter.SetStrictMode(true) // 启用严格模式
		src := SourceStruct{Name: "Alice"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err == nil {
			t.Error("Expected error for missing fields, but got nil")
		}
		if err != nil && !strings.Contains(err.Error(), "字段差异导致转换失败") {
			t.Errorf("Expected field difference error, got: %v", err)
		}
	})

	// 测试多余字段的情况
	t.Run("ExtraFields", func(t *testing.T) {
		type SourceStruct struct {
			Name string
			Age  int
			City string
		}

		type TargetStruct struct {
			Name string
			Age  int
		}

		converter := NewStructConverter()
		converter.SetStrictMode(true) // 启用严格模式
		src := SourceStruct{Name: "Alice", Age: 30, City: "Beijing"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err == nil {
			t.Error("Expected error for extra fields, but got nil")
		}
		if err != nil && !strings.Contains(err.Error(), "字段差异导致转换失败") {
			t.Errorf("Expected field difference error, got: %v", err)
		}
	})

	// 测试同时存在缺失和多余字段的情况
	t.Run("MissingAndExtraFields", func(t *testing.T) {
		type SourceStruct struct {
			Name string
			City string
		}

		type TargetStruct struct {
			Name string
			Age  int
		}

		converter := NewStructConverter()
		converter.SetStrictMode(true) // 启用严格模式
		src := SourceStruct{Name: "Alice", City: "Beijing"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err == nil {
			t.Error("Expected error for field differences, but got nil")
		}
		if err != nil && !strings.Contains(err.Error(), "字段差异导致转换失败") {
			t.Errorf("Expected field difference error, got: %v", err)
		}
	})

	// 测试非严格模式（默认模式）
	t.Run("NonStrictMode", func(t *testing.T) {
		type SourceStruct struct {
			Name string
			City string // 多余字段
		}

		type TargetStruct struct {
			Name string
			Age  int // 缺失字段
		}

		converter := NewStructConverter()
		// 不设置严格模式，默认为非严格模式
		src := SourceStruct{Name: "Alice", City: "Beijing"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err != nil {
			t.Errorf("Expected no error in non-strict mode, but got: %v", err)
		}

		// 验证匹配的字段被正确转换
		if dst.Name != "Alice" {
			t.Errorf("Expected Name to be 'Alice', got '%s'", dst.Name)
		}

		// 验证缺失字段保持零值
		if dst.Age != 0 {
			t.Errorf("Expected Age to be 0 (zero value), got %d", dst.Age)
		}
	})
}

// SimpleTestConverter 简单的测试转换器
type SimpleTestConverter struct{}

func (s *SimpleTestConverter) Convert(src reflect.Value) (reflect.Value, error) {
	if src.Kind() != reflect.String {
		return reflect.Value{}, fmt.Errorf("expected string, got %s", src.Kind())
	}
	strVal := src.String()
	intVal := int64(len(strVal)) // 简单转换：字符串长度转为int64
	return reflect.ValueOf(intVal), nil
}

// TestPointerAgnosticConversion 测试指针无关转换
func TestPointerAgnosticConversion(t *testing.T) {

	converter := NewStructConverter()
	// 注册指针无关的转换器
	converter.RegisterConverterWithOptions("string->int64", &SimpleTestConverter{}, true)

	// 测试值到值的转换
	t.Run("ValueToValue", func(t *testing.T) {
		type SourceStruct struct {
			Name string
		}
		type TargetStruct struct {
			Name int64
		}

		src := SourceStruct{Name: "hello"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if dst.Name != 5 {
			t.Errorf("Expected Name to be 5, got %d", dst.Name)
		}
	})

	// 测试值到指针的转换
	t.Run("ValueToPointer", func(t *testing.T) {
		type SourceStruct struct {
			Name string
		}
		type TargetStruct struct {
			Name *int64
		}

		src := SourceStruct{Name: "world"}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if dst.Name == nil {
			t.Error("Expected Name to be non-nil")
		} else if *dst.Name != 5 {
			t.Errorf("Expected Name to be 5, got %d", *dst.Name)
		}
	})

	// 测试指针到值的转换
	t.Run("PointerToValue", func(t *testing.T) {
		type SourceStruct struct {
			Name *string
		}
		type TargetStruct struct {
			Name int64
		}

		nameVal := "test"
		src := SourceStruct{Name: &nameVal}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if dst.Name != 4 {
			t.Errorf("Expected Name to be 4, got %d", dst.Name)
		}
	})

	// 测试指针到指针的转换
	t.Run("PointerToPointer", func(t *testing.T) {
		type SourceStruct struct {
			Name *string
		}
		type TargetStruct struct {
			Name *int64
		}

		nameVal := "golang"
		src := SourceStruct{Name: &nameVal}
		var dst TargetStruct

		err := converter.ConvertStruct(src, &dst)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if dst.Name == nil {
			t.Error("Expected Name to be non-nil")
		} else if *dst.Name != 6 {
			t.Errorf("Expected Name to be 6, got %d", *dst.Name)
		}
	})

	// 测试非指针无关转换器不会被使用
	t.Run("NonPointerAgnosticConverter", func(t *testing.T) {
		converter2 := NewStructConverter()
		// 注册非指针无关的转换器
		converter2.RegisterConverter("string->int64", &SimpleTestConverter{})

		type SourceStruct struct {
			Name *string
		}
		type TargetStruct struct {
			Name int64
		}

		nameVal := "test"
		src := SourceStruct{Name: &nameVal}
		var dst TargetStruct

		err := converter2.ConvertStruct(src, &dst)
		if err == nil {
			t.Error("Expected error for non-pointer-agnostic converter, but got nil")
		}
	})
}
