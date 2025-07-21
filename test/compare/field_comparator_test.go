package compare

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// TestFieldComparator_CompareFieldMaps 测试字段映射表比较
func TestFieldComparator_CompareFieldMaps(t *testing.T) {
	comparator := NewFieldComparator()

	// 测试用例1：相同字段映射表
	fieldMapA := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
		"Value": {
			Path:  "Value",
			Value: 100,
			Type:  reflect.TypeOf(0),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
		"Value": {
			Path:  "Value",
			Value: 100,
			Type:  reflect.TypeOf(0),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 0 {
		t.Errorf("期望无差异，但发现了 %d 个差异", len(differences))
	}
}

// TestFieldComparator_CompareFieldMaps_ValueDifferent 测试值不同的情况
func TestFieldComparator_CompareFieldMaps_ValueDifferent(t *testing.T) {
	comparator := NewFieldComparator()

	fieldMapA := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test1",
			Type:  reflect.TypeOf(""),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test2",
			Type:  reflect.TypeOf(""),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if differences[0].DiffType != DiffTypeValueDifferent {
		t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeValueDifferent, differences[0].DiffType)
	}
}

// TestFieldComparator_CompareFieldMaps_TypeDifferent 测试类型不同的情况
func TestFieldComparator_CompareFieldMaps_TypeDifferent(t *testing.T) {
	comparator := NewFieldComparator()

	fieldMapA := map[string]*FieldInfo{
		"Value": {
			Path:  "Value",
			Value: 100,
			Type:  reflect.TypeOf(0),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Value": {
			Path:  "Value",
			Value: "100",
			Type:  reflect.TypeOf(""),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if differences[0].DiffType != DiffTypeTypeDifferent {
		t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeTypeDifferent, differences[0].DiffType)
	}
}

// TestFieldComparator_CompareFieldMaps_ExtraFieldA 测试A中存在B中不存在的字段
func TestFieldComparator_CompareFieldMaps_ExtraFieldA(t *testing.T) {
	comparator := NewFieldComparator()

	fieldMapA := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
		"ExtraField": {
			Path:  "ExtraField",
			Value: "extra",
			Type:  reflect.TypeOf(""),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if differences[0].DiffType != DiffTypeExtraFieldA {
		t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeExtraFieldA, differences[0].DiffType)
	}
}

// TestFieldComparator_CompareFieldMaps_ExtraFieldB 测试B中存在A中不存在的字段
func TestFieldComparator_CompareFieldMaps_ExtraFieldB(t *testing.T) {
	comparator := NewFieldComparator()

	fieldMapA := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
		"ExtraField": {
			Path:  "ExtraField",
			Value: "extra",
			Type:  reflect.TypeOf(""),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if differences[0].DiffType != DiffTypeExtraFieldB {
		t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeExtraFieldB, differences[0].DiffType)
	}
}

// TestFieldComparator_CompareFieldMaps_WithCustomComparer 测试自定义比较器
func TestFieldComparator_CompareFieldMaps_WithCustomComparer(t *testing.T) {
	comparator := NewFieldComparator()

	now := time.Now()
	later := now.Add(30 * time.Second)

	fieldMapA := map[string]*FieldInfo{
		"CreatedAt": {
			Path:  "CreatedAt",
			Value: now,
			Type:  reflect.TypeOf(time.Time{}),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"CreatedAt": {
			Path:  "CreatedAt",
			Value: later,
			Type:  reflect.TypeOf(time.Time{}),
		},
	}

	// 不使用自定义比较器，应该有差异
	differences1 := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences1) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences1))
	}

	// 使用自定义比较器，允许1分钟内的差异
	options := &ComparisonOptions{
		CustomTypeComparers: map[string]*CustomTypeComparer{
			"time_comparer": {
				SupportedTypes:        []string{"time.Time"},
				SupportDifferentTypes: false,
				CompareFunc: func(valueA, valueB interface{}) (bool, string) {
					timeA, okA := valueA.(time.Time)
					timeB, okB := valueB.(time.Time)
					if !okA || !okB {
						return true, "无法转换为时间类型"
					}

					diff := timeA.Sub(timeB)
					if diff < 0 {
						diff = -diff
					}

					if diff > time.Minute {
						return true, "时间差异超过1分钟"
					}

					return false, "时间差异在允许范围内"
				},
			},
		},
	}

	differences2 := comparator.CompareFieldMaps(fieldMapA, fieldMapB, options)
	if len(differences2) != 0 {
		t.Errorf("期望无差异（自定义比较器允许1分钟内差异），但发现了 %d 个差异", len(differences2))
	}
}

// TestFieldComparator_CompareFieldMaps_WithCustomComparerDiff 测试自定义比较器检测到差异
func TestFieldComparator_CompareFieldMaps_WithCustomComparerDiff(t *testing.T) {
	comparator := NewFieldComparator()

	now := time.Now()
	later := now.Add(2 * time.Minute) // 2分钟后，超过允许范围

	fieldMapA := map[string]*FieldInfo{
		"CreatedAt": {
			Path:  "CreatedAt",
			Value: now,
			Type:  reflect.TypeOf(time.Time{}),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"CreatedAt": {
			Path:  "CreatedAt",
			Value: later,
			Type:  reflect.TypeOf(time.Time{}),
		},
	}

	options := &ComparisonOptions{
		CustomTypeComparers: map[string]*CustomTypeComparer{
			"time_comparer": {
				SupportedTypes:        []string{"time.Time"},
				SupportDifferentTypes: false,
				CompareFunc: func(valueA, valueB interface{}) (bool, string) {
					timeA, okA := valueA.(time.Time)
					timeB, okB := valueB.(time.Time)
					if !okA || !okB {
						return true, "无法转换为时间类型"
					}

					diff := timeA.Sub(timeB)
					if diff < 0 {
						diff = -diff
					}

					if diff > time.Minute {
						return true, "时间差异超过1分钟"
					}

					return false, "时间差异在允许范围内"
				},
			},
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, options)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if differences[0].DiffType != DiffTypeCustomComparerDiff {
		t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeCustomComparerDiff, differences[0].DiffType)
	}
}

// TestFieldComparator_CompareFieldMaps_WithArrayIndices 测试数组索引
func TestFieldComparator_CompareFieldMaps_WithArrayIndices(t *testing.T) {
	comparator := NewFieldComparator()

	fieldMapA := map[string]*FieldInfo{
		"Items[0]": {
			Path:  "Items[0]",
			Value: 100,
			Type:  reflect.TypeOf(0),
			ArrayIndices: []ArrayIndex{
				{FieldName: "Items", Index: 0},
			},
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Items[0]": {
			Path:  "Items[0]",
			Value: 200,
			Type:  reflect.TypeOf(0),
			ArrayIndices: []ArrayIndex{
				{FieldName: "Items", Index: 0},
			},
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 1 {
		t.Errorf("期望1个差异，但发现了 %d 个差异", len(differences))
	}

	if len(differences[0].ArrayIndices) != 1 {
		t.Errorf("期望1个数组索引，但得到 %d 个", len(differences[0].ArrayIndices))
	}

	if differences[0].ArrayIndices[0].FieldName != "Items" || differences[0].ArrayIndices[0].Index != 0 {
		t.Errorf("数组索引不正确，期望 Items[0]，但得到 %s[%d]", differences[0].ArrayIndices[0].FieldName, differences[0].ArrayIndices[0].Index)
	}
}

// TestFieldComparator_CompareFieldMaps_NilValues 测试nil值比较
func TestFieldComparator_CompareFieldMaps_NilValues(t *testing.T) {
	comparator := NewFieldComparator()

	// 测试两个nil值
	fieldMapA := map[string]*FieldInfo{
		"Pointer": {
			Path:  "Pointer",
			Value: nil,
			Type:  reflect.TypeOf((*string)(nil)),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Pointer": {
			Path:  "Pointer",
			Value: nil,
			Type:  reflect.TypeOf((*string)(nil)),
		},
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	if len(differences) != 0 {
		t.Errorf("期望无差异（两个nil值），但发现了 %d 个差异", len(differences))
	}

	// 测试一个nil一个非nil
	str := "test"
	fieldMapC := map[string]*FieldInfo{
		"Pointer": {
			Path:  "Pointer",
			Value: &str,
			Type:  reflect.TypeOf((*string)(nil)),
		},
	}

	differences2 := comparator.CompareFieldMaps(fieldMapA, fieldMapC, nil)
	if len(differences2) != 1 {
		t.Errorf("期望1个差异（nil vs 非nil），但发现了 %d 个差异", len(differences2))
	}
}

// TestFieldComparator_SetDebugMode 测试调试模式设置
func TestFieldComparator_SetDebugMode(t *testing.T) {
	comparator := NewFieldComparator()

	// 测试设置调试模式
	comparator.SetDebugMode(true)
	// 这里主要测试方法不会崩溃，具体的调试输出在其他测试中验证

	comparator.SetDebugMode(false)
	// 测试关闭调试模式
}

// TestFieldComparator_PrintDifferences 测试打印差异功能
func TestFieldComparator_PrintDifferences(t *testing.T) {
	comparator := NewFieldComparator()

	// 测试空差异列表
	comparator.PrintDifferences("空差异测试", []DifferenceDetail{})

	// 测试有差异的情况
	differences := []DifferenceDetail{
		{
			DiffType:   DiffTypeValueDifferent,
			FieldNameA: "Name",
			FieldNameB: "Name",
			ValueA:     "test1",
			ValueB:     "test2",
			TypeA:      "string",
			TypeB:      "string",
			Message:    "字段值不同",
		},
	}

	comparator.PrintDifferences("有差异测试", differences)
}

// TestFieldComparator_ComplexScenario 测试复杂场景
func TestFieldComparator_ComplexScenario(t *testing.T) {
	comparator := NewFieldComparator()
	comparator.SetDebugMode(true)

	// 创建复杂的字段映射表
	fieldMapA := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test",
			Type:  reflect.TypeOf(""),
		},
		"Value": {
			Path:  "Value",
			Value: 100,
			Type:  reflect.TypeOf(0),
		},
		"Items[0]": {
			Path:  "Items[0]",
			Value: "item1",
			Type:  reflect.TypeOf(""),
			ArrayIndices: []ArrayIndex{
				{FieldName: "Items", Index: 0},
			},
		},
		"ExtraFieldA": {
			Path:  "ExtraFieldA",
			Value: "extra",
			Type:  reflect.TypeOf(""),
		},
	}

	fieldMapB := map[string]*FieldInfo{
		"Name": {
			Path:  "Name",
			Value: "test_modified", // 值不同
			Type:  reflect.TypeOf(""),
		},
		"Value": {
			Path:  "Value",
			Value: "100", // 类型不同
			Type:  reflect.TypeOf(""),
		},
		"Items[0]": {
			Path:  "Items[0]",
			Value: "item1", // 相同
			Type:  reflect.TypeOf(""),
			ArrayIndices: []ArrayIndex{
				{FieldName: "Items", Index: 0},
			},
		},
		"ExtraFieldB": {
			Path:  "ExtraFieldB",
			Value: "extra_b",
			Type:  reflect.TypeOf(""),
		},
		// 缺少 ExtraFieldA
	}

	differences := comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)

	// 应该有4个差异：
	// 1. Name值不同
	// 2. Value类型不同
	// 3. ExtraFieldA在A中存在B中不存在
	// 4. ExtraFieldB在B中存在A中不存在
	if len(differences) != 4 {
		t.Errorf("期望4个差异，但发现了 %d 个差异", len(differences))
	}

	// 验证差异类型
	diffTypes := make(map[DifferenceType]int)
	for _, diff := range differences {
		diffTypes[diff.DiffType]++
	}

	expectedTypes := map[DifferenceType]int{
		DiffTypeValueDifferent: 1,
		DiffTypeTypeDifferent:  1,
		DiffTypeExtraFieldA:    1,
		DiffTypeExtraFieldB:    1,
	}

	for expectedType, expectedCount := range expectedTypes {
		if diffTypes[expectedType] != expectedCount {
			t.Errorf("期望差异类型 %s 有 %d 个，但实际有 %d 个", expectedType, expectedCount, diffTypes[expectedType])
		}
	}
}

// BenchmarkFieldComparator_CompareFieldMaps 基准测试
func BenchmarkFieldComparator_CompareFieldMaps(b *testing.B) {
	comparator := NewFieldComparator()

	// 创建大型字段映射表
	fieldMapA := make(map[string]*FieldInfo)
	fieldMapB := make(map[string]*FieldInfo)

	for i := 0; i < 1000; i++ {
		path := fmt.Sprintf("Field%d", i)
		fieldMapA[path] = &FieldInfo{
			Path:  path,
			Value: i,
			Type:  reflect.TypeOf(0),
		}
		fieldMapB[path] = &FieldInfo{
			Path:  path,
			Value: i,
			Type:  reflect.TypeOf(0),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comparator.CompareFieldMaps(fieldMapA, fieldMapB, nil)
	}
}