package utils

import (
	"testing"
)

// 测试用的结构体定义
type TestStructA struct {
	B []TestStructB
}

type TestStructB struct {
	B1 int
	C  []TestStructC
}

type TestStructC struct {
	D int
}

type TestStructBA struct {
	BB TestStructBB
}

type TestStructBB struct {
	BC []TestStructBC
	CC []TestStructCC
}

type TestStructBC struct {
	BD1 []int
	BD2 []int
	BD3 int
}

type TestStructCC struct {
	D1 []int
	D2 []int
}

// TestFieldMappingComparator 测试字段映射比较器
func TestFieldMappingComparator(t *testing.T) {
	// 创建测试数据
	sourceData := TestStructA{
		B: []TestStructB{
			{
				B1: 100,
				C: []TestStructC{
					{D: 110},
					{D: 120},
				},
			},
			{
				B1: 200,
				C: []TestStructC{
					{D: 210},
					{D: 220},
					{D: 230},
				},
			},
		},
	}

	targetData := TestStructBA{
		BB: TestStructBB{
			BC: []TestStructBC{
				{
					BD1: []int{110, 120},
					BD2: []int{210, 220, 230},
					BD3: 100,
				},
				{
					BD1: []int{310, 320},
					BD2: []int{410, 420, 430},
					BD3: 200,
				},
			},
			CC: []TestStructCC{
				{
					D1: []int{110, 120},
					D2: []int{210, 220, 230},
				},
				{
					D1: []int{310, 320},
					D2: []int{410, 420, 230}, // 修改为230以匹配源数据
				},
			},
		},
	}

	// 定义字段映射
	mappings := []FieldMapping{
		{
			SourcePath: "B[1].B1",
			TargetPath: "BB.BC[1].BD3",
		},
		{
			SourcePath: "B[1].C[2].D",
			TargetPath: "BB.CC[1].D2[2]",
		},
	}

	// 创建映射选项
	options := FieldMappingOptions{
		Mappings:         mappings,
		ExpectDifference: false,
		Logger:           NewTestLogger(t),
	}

	// 执行比较
	CompareResponsesV2(t, sourceData, targetData, "TestFieldMapping", options)
}

// TestFieldMappingComparatorBasic 测试基本的字段映射功能
func TestFieldMappingComparatorBasic(t *testing.T) {
	comparator := NewFieldMappingComparator()

	// 简单的测试数据
	type SimpleA struct {
		Value int
	}

	type SimpleB struct {
		Data int
	}

	sourceData := SimpleA{Value: 42}
	targetData := SimpleB{Data: 42}

	// 测试字段映射表解析
	sourceFieldMap := comparator.parseStructToFieldMap(sourceData, "")
	targetFieldMap := comparator.parseStructToFieldMap(targetData, "")

	t.Logf("源字段映射表: %+v", sourceFieldMap)
	t.Logf("目标字段映射表: %+v", targetFieldMap)

	// 验证字段映射表不为空
	if len(sourceFieldMap) == 0 {
		t.Error("源字段映射表为空")
	}
	if len(targetFieldMap) == 0 {
		t.Error("目标字段映射表为空")
	}
}

// TestFieldMappingWithArrays 测试包含数组的字段映射
func TestFieldMappingWithArrays(t *testing.T) {
	comparator := NewFieldMappingComparator()

	// 包含数组的测试数据
	type ArrayStruct struct {
		Items []int
	}

	data := ArrayStruct{
		Items: []int{1, 2, 3},
	}

	// 测试数组字段映射表解析
	fieldMap := comparator.parseStructToFieldMap(data, "")
	t.Logf("数组字段映射表: %+v", fieldMap)

	// 验证数组字段被正确解析
	if len(fieldMap) == 0 {
		t.Error("数组字段映射表为空")
	}
}