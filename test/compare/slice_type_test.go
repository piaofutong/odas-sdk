package compare

import (
	"reflect"
	"testing"
	"time"
)

// TestIsLeafField_BasicSlices 测试基础类型切片的叶子字段判断
func TestIsLeafField_BasicSlices(t *testing.T) {
	comparator := NewFieldComparator()

	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"int切片", []int{1, 2, 3}, true},
		{"string切片", []string{"a", "b"}, true},
		{"bool切片", []bool{true, false}, true},
		{"float64切片", []float64{1.1, 2.2}, true},
		{"空int切片", []int{}, true},
		{"nil int切片", ([]int)(nil), true},
		{"int指针切片", []*int{}, true},          // 基础类型指针的切片也是叶子
		{"time.Time切片", []time.Time{}, true}, // 特殊类型的切片是叶子
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldInfo := &FieldInfo{
				Value: tt.value,
				Type:  reflect.TypeOf(tt.value),
			}
			result := comparator.IsLeafField(fieldInfo)
			if result != tt.expected {
				t.Errorf("IsLeafField() = %v, expected %v for %s", result, tt.expected, tt.name)
			}
		})
	}
}

// TestIsLeafField_ComplexSlices 测试复杂类型切片的叶子字段判断
func TestIsLeafField_ComplexSlices(t *testing.T) {
	comparator := NewFieldComparator()

	type ComplexStruct struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"结构体切片", []ComplexStruct{}, false}, // 复杂结构体的切片不是叶子
		{"结构体指针切片", []*ComplexStruct{}, false},
		{"嵌套切片", [][]int{}, true},            // 嵌套切片不是叶子
		{"map切片", []map[string]int{}, false}, // map切片不是叶子
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldInfo := &FieldInfo{
				Value: tt.value,
				Type:  reflect.TypeOf(tt.value),
			}
			result := comparator.IsLeafField(fieldInfo)
			if result != tt.expected {
				t.Errorf("IsLeafField() = %v, expected %v for %s", result, tt.expected, tt.name)
			}
		})
	}
}

// TestEmptySliceTypeDifference 测试空切片类型差异检测
func TestEmptySliceTypeDifference(t *testing.T) {
	type StructA struct {
		Name string
		Age  int
	}

	type StructB struct {
		Title string
		Value int
	}

	type ResponseA struct {
		List []*StructA
	}

	type ResponseB struct {
		List []*StructB
	}

	tests := []struct {
		name         string
		structA      interface{}
		structB      interface{}
		expectedDiff bool
		description  string
	}{
		{
			name:         "空切片不同元素类型",
			structA:      &ResponseA{List: []*StructA{}},
			structB:      &ResponseB{List: []*StructB{}},
			expectedDiff: true,
			description:  "空切片应该能检测到元素类型差异",
		},
		{
			name:         "nil切片不同元素类型",
			structA:      &ResponseA{List: nil},
			structB:      &ResponseB{List: nil},
			expectedDiff: true,
			description:  "nil切片应该能检测到元素类型差异",
		},
		{
			name:         "相同类型空切片",
			structA:      &ResponseA{List: []*StructA{}},
			structB:      &ResponseA{List: []*StructA{}},
			expectedDiff: false,
			description:  "相同类型的空切片不应该有差异",
		},
		{
			name:         "基础类型切片差异",
			structA:      struct{ List []int }{List: []int{}},
			structB:      struct{ List []string }{List: []string{}},
			expectedDiff: true,
			description:  "不同基础类型的切片应该能检测到差异",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareResponses(tt.structA, tt.structB, tt.name, nil)
			if result.HasDiff != tt.expectedDiff {
				t.Errorf("%s: HasDiff = %v, expected %v\n%s",
					tt.name, result.HasDiff, tt.expectedDiff, tt.description)
				if len(result.Differences) > 0 {
					t.Logf("发现的差异:")
					for i, diff := range result.Differences {
						t.Logf("  差异%d: %s - %s", i+1, diff.DiffType.String(), diff.Message)
					}
				}
			}
		})
	}
}

// TestBasicSliceComparison 测试基础类型切片的比较
func TestBasicSliceComparison(t *testing.T) {
	tests := []struct {
		name         string
		structA      interface{}
		structB      interface{}
		expectedDiff bool
	}{
		{
			name:         "相同int切片",
			structA:      struct{ List []int }{List: []int{1, 2, 3}},
			structB:      struct{ List []int }{List: []int{1, 2, 3}},
			expectedDiff: false,
		},
		{
			name:         "不同int切片值",
			structA:      struct{ List []int }{List: []int{1, 2, 3}},
			structB:      struct{ List []int }{List: []int{1, 2, 4}},
			expectedDiff: true,
		},
		{
			name:         "int vs string切片",
			structA:      struct{ List []int }{List: []int{1, 2, 3}},
			structB:      struct{ List []string }{List: []string{"1", "2", "3"}},
			expectedDiff: true,
		},
		{
			name:         "空string切片 vs nil",
			structA:      struct{ List []string }{List: []string{}},
			structB:      struct{ List []string }{List: nil},
			expectedDiff: false, // 空切片和nil在Go中被认为是相等的
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareResponses(tt.structA, tt.structB, tt.name, nil)
			if result.HasDiff != tt.expectedDiff {
				t.Errorf("%s: HasDiff = %v, expected %v", tt.name, result.HasDiff, tt.expectedDiff)
				if len(result.Differences) > 0 {
					t.Logf("发现的差异:")
					for i, diff := range result.Differences {
						t.Logf("  差异%d: %s - %s", i+1, diff.DiffType.String(), diff.Message)
					}
				}
			}
		})
	}
}

// TestSliceElementTypeDifference 测试切片元素类型差异检测的辅助方法
func TestSliceElementTypeDifference(t *testing.T) {
	comparator := NewFieldComparator()

	tests := []struct {
		name     string
		typeA    reflect.Type
		typeB    reflect.Type
		expected bool
	}{
		{
			name:     "相同int切片类型",
			typeA:    reflect.TypeOf([]int{}),
			typeB:    reflect.TypeOf([]int{}),
			expected: false,
		},
		{
			name:     "int vs string切片",
			typeA:    reflect.TypeOf([]int{}),
			typeB:    reflect.TypeOf([]string{}),
			expected: true,
		},
		{
			name:     "结构体指针切片不同类型",
			typeA:    reflect.TypeOf([]*struct{ A int }{}),
			typeB:    reflect.TypeOf([]*struct{ B string }{}),
			expected: true,
		},
		{
			name:     "nil类型",
			typeA:    nil,
			typeB:    reflect.TypeOf([]int{}),
			expected: true,
		},
		{
			name:     "非切片类型",
			typeA:    reflect.TypeOf("string"),
			typeB:    reflect.TypeOf([]int{}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := comparator.hasSliceElementTypeDifference(tt.typeA, tt.typeB)
			if result != tt.expected {
				t.Errorf("hasSliceElementTypeDifference() = %v, expected %v for %s",
					result, tt.expected, tt.name)
			}
		})
	}
}
