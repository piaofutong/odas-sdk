package utils

import (
	"testing"
)

// CompareResponses 直接比较两个响应结构体，自动提取和比较相同字段
func CompareResponses(t *testing.T, apiResp, grpcResp interface{}, testName string) {
	CompareResponsesWithExpectation(t, apiResp, grpcResp, testName, false)
}

// CompareResponsesWithExpectation 比较响应并支持指定是否期望有差异
func CompareResponsesWithExpectation(t *testing.T, apiResp, grpcResp interface{}, testName string, expectDifference bool) {
	// 使用面向对象的比较器
	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	options := ComparisonOptions{
		ExpectDifference: expectDifference,
		Logger:           logger,
	}

	result := comparator.CompareWithOptions(apiResp, grpcResp, testName, options)

	printer := NewResponseResultPrinter()
	hasDifference := printer.PrintResult(result, testName)

	if expectDifference {
		if hasDifference {
			t.Logf("%s 比较完成，如期发现差异", testName)
		} else {
			t.Errorf("%s 比较失败，期望有差异但未发现", testName)
		}
	} else {
		if hasDifference {
			t.Errorf("%s 比较失败，发现差异", testName)
		} else {
			// t.Logf("%s 所有字段比较成功", testName)
		}
	}
}

// FieldMapping 字段映射配置
type FieldMapping struct {
	SourcePath string // 源字段路径，如 "A.B[1].C[2].D"
	TargetPath string // 目标字段路径，如 "BA.BB.BC[1].BD3"
}

// FieldMappingOptions 字段映射选项
type FieldMappingOptions struct {
	Mappings         []FieldMapping // 字段映射配置列表
	ExpectDifference bool           // 是否期望有差异
	Logger           Logger         // 日志记录器
}

// CompareResponsesV2 支持字段重映射的比较方法
func CompareResponsesV2(t *testing.T, sourceResp, targetResp interface{}, testName string, options FieldMappingOptions) {
	// 创建字段映射比较器
	comparator := NewFieldMappingComparator()
	logger := options.Logger
	if logger == nil {
		logger = NewTestLogger(t)
	}

	// 执行字段映射比较
	result := comparator.CompareWithMappings(sourceResp, targetResp, testName, options)

	// 打印结果
	printer := NewResponseResultPrinter()
	hasDifference := printer.PrintResult(result, testName)

	if options.ExpectDifference {
		if hasDifference {
			t.Logf("%s 字段映射比较完成，如期发现差异", testName)
		} else {
			t.Errorf("%s 字段映射比较失败，期望有差异但未发现", testName)
		}
	} else {
		if hasDifference {
			t.Errorf("%s 字段映射比较失败，发现差异", testName)
		} else {
			t.Logf("%s 字段映射比较成功，所有映射字段无差异", testName)
		}
	}
}
