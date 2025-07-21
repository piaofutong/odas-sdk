package compare

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

// 测试用的结构体定义
type SimpleStruct struct {
	Name  string
	Value int
}

type NestedStruct struct {
	ID     int
	Simple SimpleStruct
	Items  []string
}

type SelfReferencingStruct struct {
	ID       int
	Name     string
	Children []*SelfReferencingStruct
	Parent   *SelfReferencingStruct
}

type MultiLevelArray struct {
	Matrix2D [][]int
	Matrix3D [][][]string
	Mixed    [][]SimpleStruct
}

type InterfaceStruct struct {
	Data interface{}
	Any  any
}

type PointerStruct struct {
	IntPtr    *int
	StringPtr *string
	StructPtr *SimpleStruct
	NilPtr    *int
}

type EmptyStruct struct{}

type MapStruct struct {
	StringMap map[string]int
	IntMap    map[int]string
	NestedMap map[string]map[string]int
}

type UncomparableStruct struct {
	Channel chan int
	Mutex   sync.Mutex
	Func    func()
}

type TimeStruct struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type ArrayStruct struct {
	FixedArray [3]int
	Slice      []int
}

type MixedTypesStruct struct {
	String  string
	Int     int
	Float   float64
	Bool    bool
	Bytes   []byte
	Rune    rune
	UintPtr uintptr
}

type unexportedFieldStruct struct {
	privateField string
	PrivateField string
}

// TestStructParser_NewStructParser 测试构造函数
func TestStructParser_NewStructParser(t *testing.T) {
	parser := NewStructParser()
	if parser == nil {
		t.Fatal("NewStructParser 返回 nil")
	}
	if parser.debugMode {
		t.Error("新创建的解析器应该默认关闭调试模式")
	}
}

// TestStructParser_SetDebugMode 测试调试模式设置
func TestStructParser_SetDebugMode(t *testing.T) {
	parser := NewStructParser()

	// 测试开启调试模式
	parser.SetDebugMode(true)
	if !parser.debugMode {
		t.Error("调试模式应该被开启")
	}

	// 测试关闭调试模式
	parser.SetDebugMode(false)
	if parser.debugMode {
		t.Error("调试模式应该被关闭")
	}
}

// TestStructParser_ParseToFieldMap_SimpleStruct 测试简单结构体解析
func TestStructParser_ParseToFieldMap_SimpleStruct(t *testing.T) {
	parser := NewStructParser()
	struct1 := SimpleStruct{
		Name:  "test",
		Value: 123,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	if len(fieldMap) != 2 {
		t.Errorf("期望解析出 2 个字段，实际得到 %d 个", len(fieldMap))
	}

	// 检查 Name 字段
	nameField, exists := fieldMap["Name"]
	if !exists {
		t.Error("Name 字段不存在")
	}
	if nameField.Value != "test" {
		t.Errorf("Name 字段值错误，期望 'test'，实际 %v", nameField.Value)
	}
	if nameField.Type.String() != "string" {
		t.Errorf("Name 字段类型错误，期望 'string'，实际 %s", nameField.Type.String())
	}

	// 检查 Value 字段
	valueField, exists := fieldMap["Value"]
	if !exists {
		t.Error("Value 字段不存在")
	}
	if valueField.Value != 123 {
		t.Errorf("Value 字段值错误，期望 123，实际 %v", valueField.Value)
	}
	if valueField.Type.String() != "int" {
		t.Errorf("Value 字段类型错误，期望 'int'，实际 %s", valueField.Type.String())
	}
}

// TestStructParser_ParseToFieldMap_NestedStruct 测试嵌套结构体解析
func TestStructParser_ParseToFieldMap_NestedStruct(t *testing.T) {
	parser := NewStructParser()
	struct1 := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name:  "nested",
			Value: 456,
		},
		Items: []string{"item1", "item2"},
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 检查顶级字段
	if _, exists := fieldMap["ID"]; !exists {
		t.Error("ID 字段不存在")
	}

	// 检查嵌套字段
	if _, exists := fieldMap["Simple.Name"]; !exists {
		t.Error("Simple.Name 字段不存在")
	}
	if _, exists := fieldMap["Simple.Value"]; !exists {
		t.Error("Simple.Value 字段不存在")
	}

	// 检查数组字段
	if _, exists := fieldMap["Items[0]"]; !exists {
		t.Error("Items[0] 字段不存在")
	}
	if _, exists := fieldMap["Items[1]"]; !exists {
		t.Error("Items[1] 字段不存在")
	}

	// 验证嵌套字段值
	if fieldMap["Simple.Name"].Value != "nested" {
		t.Errorf("Simple.Name 值错误，期望 'nested'，实际 %v", fieldMap["Simple.Name"].Value)
	}
	if fieldMap["Items[0]"].Value != "item1" {
		t.Errorf("Items[0] 值错误，期望 'item1'，实际 %v", fieldMap["Items[0]"].Value)
	}
}

// TestStructParser_ParseToFieldMap_SelfReferencing 测试自引用结构体
func TestStructParser_ParseToFieldMap_SelfReferencing(t *testing.T) {
	parser := NewStructParser()

	// 创建简单的自引用结构，避免循环引用
	child1 := &SelfReferencingStruct{
		ID:   2,
		Name: "child1",
		// 不设置 Parent 避免循环引用
	}
	child2 := &SelfReferencingStruct{
		ID:   3,
		Name: "child2",
		// 不设置 Parent 避免循环引用
	}
	parent := &SelfReferencingStruct{
		ID:       1,
		Name:     "parent",
		Children: []*SelfReferencingStruct{child1, child2},
		// 不设置 Parent 避免循环引用
	}

	fieldMap := parser.ParseToFieldMap(parent)

	// 检查基本字段
	if fieldMap["ID"].Value != 1 {
		t.Errorf("ID 值错误，期望 1，实际 %v", fieldMap["ID"].Value)
	}
	if fieldMap["Name"].Value != "parent" {
		t.Errorf("Name 值错误，期望 'parent'，实际 %v", fieldMap["Name"].Value)
	}

	// 检查子节点字段
	if _, exists := fieldMap["Children[0].ID"]; !exists {
		t.Error("Children[0].ID 字段不存在")
	}
	if _, exists := fieldMap["Children[0].Name"]; !exists {
		t.Error("Children[0].Name 字段不存在")
	}
	if _, exists := fieldMap["Children[1].ID"]; !exists {
		t.Error("Children[1].ID 字段不存在")
	}

	// 检查 nil 父节点引用
	if _, exists := fieldMap["Children[0].Parent"]; !exists {
		t.Error("Children[0].Parent 字段不存在")
	}
	if fieldMap["Children[0].Parent"].Value != nil {
		t.Error("Children[0].Parent 应该是 nil")
	}
}

// TestStructParser_ParseToFieldMap_MultiLevelArray 测试多级数组
func TestStructParser_ParseToFieldMap_MultiLevelArray(t *testing.T) {
	parser := NewStructParser()
	struct1 := MultiLevelArray{
		Matrix2D: [][]int{{1, 2}, {3, 4}},
		Matrix3D: [][][]string{{{"a", "b"}, {"c", "d"}}, {{"e", "f"}, {"g", "h"}}},
		Mixed: [][]SimpleStruct{
			{{Name: "s1", Value: 1}, {Name: "s2", Value: 2}},
			{{Name: "s3", Value: 3}},
		},
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 检查二维数组
	if _, exists := fieldMap["Matrix2D[0][0]"]; !exists {
		t.Error("Matrix2D[0][0] 字段不存在")
	}
	if _, exists := fieldMap["Matrix2D[1][1]"]; !exists {
		t.Error("Matrix2D[1][1] 字段不存在")
	}
	if fieldMap["Matrix2D[0][0]"].Value != 1 {
		t.Errorf("Matrix2D[0][0] 值错误，期望 1，实际 %v", fieldMap["Matrix2D[0][0]"].Value)
	}

	// 检查三维数组
	if _, exists := fieldMap["Matrix3D[0][0][0]"]; !exists {
		t.Error("Matrix3D[0][0][0] 字段不存在")
	}
	if fieldMap["Matrix3D[0][0][0]"].Value != "a" {
		t.Errorf("Matrix3D[0][0][0] 值错误，期望 'a'，实际 %v", fieldMap["Matrix3D[0][0][0]"].Value)
	}

	// 检查混合数组
	if _, exists := fieldMap["Mixed[0][0].Name"]; !exists {
		t.Error("Mixed[0][0].Name 字段不存在")
	}
	if fieldMap["Mixed[0][0].Name"].Value != "s1" {
		t.Errorf("Mixed[0][0].Name 值错误，期望 's1'，实际 %v", fieldMap["Mixed[0][0].Name"].Value)
	}
}

// TestStructParser_ParseToFieldMap_Interface 测试接口类型
func TestStructParser_ParseToFieldMap_Interface(t *testing.T) {
	parser := NewStructParser()
	struct1 := InterfaceStruct{
		Data: "string data",
		Any:  123,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	if _, exists := fieldMap["Data"]; !exists {
		t.Error("Data 字段不存在")
	}
	if _, exists := fieldMap["Any"]; !exists {
		t.Error("Any 字段不存在")
	}

	if fieldMap["Data"].Value != "string data" {
		t.Errorf("Data 值错误，期望 'string data'，实际 %v", fieldMap["Data"].Value)
	}
	if fieldMap["Any"].Value != 123 {
		t.Errorf("Any 值错误，期望 123，实际 %v", fieldMap["Any"].Value)
	}
}

// TestStructParser_ParseToFieldMap_NilInterface 测试 nil 接口
func TestStructParser_ParseToFieldMap_NilInterface(t *testing.T) {
	parser := NewStructParser()
	struct1 := InterfaceStruct{
		Data: nil,
		Any:  nil,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	if _, exists := fieldMap["Data"]; !exists {
		t.Error("Data 字段不存在")
	}
	if _, exists := fieldMap["Any"]; !exists {
		t.Error("Any 字段不存在")
	}

	if fieldMap["Data"].Value != nil {
		t.Errorf("Data 值错误，期望 nil，实际 %v", fieldMap["Data"].Value)
	}
	if fieldMap["Any"].Value != nil {
		t.Errorf("Any 值错误，期望 nil，实际 %v", fieldMap["Any"].Value)
	}
}

// TestStructParser_ParseToFieldMap_Pointers 测试指针类型
func TestStructParser_ParseToFieldMap_Pointers(t *testing.T) {
	parser := NewStructParser()

	intVal := 42
	stringVal := "hello"
	structVal := SimpleStruct{Name: "test", Value: 100}

	struct1 := PointerStruct{
		IntPtr:    &intVal,
		StringPtr: &stringVal,
		StructPtr: &structVal,
		NilPtr:    nil,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 检查非 nil 指针
	if _, exists := fieldMap["IntPtr"]; !exists {
		t.Error("IntPtr 字段不存在")
	}
	if fieldMap["IntPtr"].Value != 42 {
		t.Errorf("IntPtr 值错误，期望 42，实际 %v", fieldMap["IntPtr"].Value)
	}

	if _, exists := fieldMap["StringPtr"]; !exists {
		t.Error("StringPtr 字段不存在")
	}
	if fieldMap["StringPtr"].Value != "hello" {
		t.Errorf("StringPtr 值错误，期望 'hello'，实际 %v", fieldMap["StringPtr"].Value)
	}

	// 检查结构体指针的嵌套字段
	if _, exists := fieldMap["StructPtr.Name"]; !exists {
		t.Error("StructPtr.Name 字段不存在")
	}
	if fieldMap["StructPtr.Name"].Value != "test" {
		t.Errorf("StructPtr.Name 值错误，期望 'test'，实际 %v", fieldMap["StructPtr.Name"].Value)
	}

	// 检查 nil 指针
	if _, exists := fieldMap["NilPtr"]; !exists {
		t.Error("NilPtr 字段不存在")
	}
	if fieldMap["NilPtr"].Value != nil {
		t.Errorf("NilPtr 值错误，期望 nil，实际 %v", fieldMap["NilPtr"].Value)
	}
}

// TestStructParser_ParseToFieldMap_EmptyStruct 测试空结构体
func TestStructParser_ParseToFieldMap_EmptyStruct(t *testing.T) {
	parser := NewStructParser()
	struct1 := EmptyStruct{}

	fieldMap := parser.ParseToFieldMap(struct1)

	if len(fieldMap) != 0 {
		t.Errorf("空结构体应该解析出 0 个字段，实际得到 %d 个", len(fieldMap))
	}
}

// TestStructParser_ParseToFieldMap_Map 测试 map 类型
func TestStructParser_ParseToFieldMap_Map(t *testing.T) {
	parser := NewStructParser()
	struct1 := MapStruct{
		StringMap: map[string]int{"key1": 1, "key2": 2},
		IntMap:    map[int]string{1: "value1", 2: "value2"},
		NestedMap: map[string]map[string]int{
			"outer1": {"inner1": 10, "inner2": 20},
		},
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// Map 类型应该被当作基本类型处理
	if _, exists := fieldMap["StringMap"]; !exists {
		t.Error("StringMap 字段不存在")
	}
	if _, exists := fieldMap["IntMap"]; !exists {
		t.Error("IntMap 字段不存在")
	}
	if _, exists := fieldMap["NestedMap"]; !exists {
		t.Error("NestedMap 字段不存在")
	}

	// 验证 map 值
	if reflect.TypeOf(fieldMap["StringMap"].Value).Kind() != reflect.Map {
		t.Error("StringMap 应该是 map 类型")
	}
}

// TestStructParser_ParseToFieldMap_Time 测试时间类型
func TestStructParser_ParseToFieldMap_Time(t *testing.T) {
	parser := NewStructParser()
	now := time.Now()
	struct1 := TimeStruct{
		CreatedAt: now,
		UpdatedAt: &now,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// time.Time 应该被当作特殊类型处理
	if _, exists := fieldMap["CreatedAt"]; !exists {
		t.Error("CreatedAt 字段不存在")
	}
	if _, exists := fieldMap["UpdatedAt"]; !exists {
		t.Error("UpdatedAt 字段不存在")
	}

	// 验证时间值
	if !fieldMap["CreatedAt"].Value.(time.Time).Equal(now) {
		t.Error("CreatedAt 时间值不正确")
	}
	if !fieldMap["UpdatedAt"].Value.(time.Time).Equal(now) {
		t.Error("UpdatedAt 时间值不正确")
	}
}

// TestStructParser_ParseToFieldMap_Array 测试数组类型
func TestStructParser_ParseToFieldMap_Array(t *testing.T) {
	parser := NewStructParser()
	struct1 := ArrayStruct{
		FixedArray: [3]int{1, 2, 3},
		Slice:      []int{4, 5, 6, 7},
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 检查固定数组元素
	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("FixedArray[%d]", i)
		if _, exists := fieldMap[key]; !exists {
			t.Errorf("%s 字段不存在", key)
		}
		if fieldMap[key].Value != i+1 {
			t.Errorf("%s 值错误，期望 %d，实际 %v", key, i+1, fieldMap[key].Value)
		}
	}

	// 检查切片元素
	for i := 0; i < 4; i++ {
		key := fmt.Sprintf("Slice[%d]", i)
		if _, exists := fieldMap[key]; !exists {
			t.Errorf("%s 字段不存在", key)
		}
		if fieldMap[key].Value != i+4 {
			t.Errorf("%s 值错误，期望 %d，实际 %v", key, i+4, fieldMap[key].Value)
		}
	}
}

// TestStructParser_ParseToFieldMap_MixedTypes 测试混合类型
func TestStructParser_ParseToFieldMap_MixedTypes(t *testing.T) {
	parser := NewStructParser()
	struct1 := MixedTypesStruct{
		String:  "test",
		Int:     42,
		Float:   3.14,
		Bool:    true,
		Bytes:   []byte{1, 2, 3},
		Rune:    'A',
		UintPtr: uintptr(0x1000), // 使用固定值避免循环引用
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	tests := []struct {
		field    string
		expected interface{}
		typeStr  string
	}{
		{"String", "test", "string"},
		{"Int", 42, "int"},
		{"Float", 3.14, "float64"},
		{"Bool", true, "bool"},
		{"Rune", 'A', "int32"},
	}

	for _, test := range tests {
		if _, exists := fieldMap[test.field]; !exists {
			t.Errorf("%s 字段不存在", test.field)
			continue
		}
		if fieldMap[test.field].Value != test.expected {
			t.Errorf("%s 值错误，期望 %v，实际 %v", test.field, test.expected, fieldMap[test.field].Value)
		}
		if fieldMap[test.field].Type.String() != test.typeStr {
			t.Errorf("%s 类型错误，期望 %s，实际 %s", test.field, test.typeStr, fieldMap[test.field].Type.String())
		}
	}

	// 检查字节数组
	if _, exists := fieldMap["Bytes[0]"]; !exists {
		t.Error("Bytes[0] 字段不存在")
	}
	if fieldMap["Bytes[0]"].Value != byte(1) {
		t.Errorf("Bytes[0] 值错误，期望 1，实际 %v", fieldMap["Bytes[0]"].Value)
	}
}

// TestStructParser_ParseToFieldMap_UnexportedFields 测试未导出字段
func TestStructParser_ParseToFieldMap_UnexportedFields(t *testing.T) {
	parser := NewStructParser()
	struct1 := unexportedFieldStruct{
		privateField: "private",
		PrivateField: "public",
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 只有导出字段应该被解析
	if _, exists := fieldMap["privateField"]; exists {
		t.Error("未导出字段 privateField 不应该被解析")
	}
	if _, exists := fieldMap["PrivateField"]; !exists {
		t.Error("导出字段 PrivateField 应该被解析")
	}
	if fieldMap["PrivateField"].Value != "public" {
		t.Errorf("PrivateField 值错误，期望 'public'，实际 %v", fieldMap["PrivateField"].Value)
	}
}

// TestStructParser_ApplyFieldMappings 测试字段映射
func TestStructParser_ApplyFieldMappings(t *testing.T) {
	parser := NewStructParser()
	struct1 := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name:  "test",
			Value: 123,
		},
		Items: []string{"item1", "item2"},
	}

	sourceFieldMap := parser.ParseToFieldMap(struct1)

	// 定义字段映射：将现有字段映射到新字段名
	mappings := []FieldMapping{
		{
			SourcePath: "Simple.Name",
			TargetPath: "NewName",
		},
		{
			SourcePath: "Items[0]",
			TargetPath: "NewItems[0]",
		},
		{
			SourcePath: "Items[1]",
			TargetPath: "NewItems[1]",
		},
	}

	mappedFieldMap := parser.ApplyFieldMappings(sourceFieldMap, mappings)

	// 检查原始字段被移除
	if _, exists := mappedFieldMap["Simple.Name"]; exists {
		t.Error("原始字段 Simple.Name 应该被移除")
	}

	// 检查映射后的字段
	if _, exists := mappedFieldMap["NewName"]; !exists {
		t.Error("映射字段 NewName 不存在")
	} else if mappedFieldMap["NewName"].Value != "test" {
		t.Errorf("映射字段 NewName 值错误，期望 'test'，实际 %v", mappedFieldMap["NewName"].Value)
	}

	// 检查数组映射
	if _, exists := mappedFieldMap["NewItems[0]"]; !exists {
		t.Error("映射字段 NewItems[0] 不存在")
	} else if mappedFieldMap["NewItems[0]"].Value != "item1" {
		t.Errorf("映射字段 NewItems[0] 值错误，期望 'item1'，实际 %v", mappedFieldMap["NewItems[0]"].Value)
	}

	if _, exists := mappedFieldMap["NewItems[1]"]; !exists {
		t.Error("映射字段 NewItems[1] 不存在")
	} else if mappedFieldMap["NewItems[1]"].Value != "item2" {
		t.Errorf("映射字段 NewItems[1] 值错误，期望 'item2'，实际 %v", mappedFieldMap["NewItems[1]"].Value)
	}
}

// TestStructParser_ParseToFieldMap_EmptySlice 测试空切片
func TestStructParser_ParseToFieldMap_EmptySlice(t *testing.T) {
	parser := NewStructParser()
	struct1 := struct {
		EmptySlice []string
		NilSlice   []int
	}{
		EmptySlice: []string{},
		NilSlice:   nil,
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 空切片应该产生字段，字段值应该为nil
	if field, exists := fieldMap["EmptySlice"]; !exists {
		t.Error("空切片应该产生字段")
	} else if field.Value != nil {
		t.Errorf("空切片字段值应该为nil，实际为 %v", field.Value)
	}
	
	// nil切片应该产生字段，字段值应该为nil
	if field, exists := fieldMap["NilSlice"]; !exists {
		t.Error("nil 切片应该产生字段")
	} else if field.Value != nil {
		t.Errorf("nil切片字段值应该为nil，实际为 %v", field.Value)
	}
}

// TestStructParser_ParseToFieldMap_ComplexNesting 测试复杂嵌套
func TestStructParser_ParseToFieldMap_ComplexNesting(t *testing.T) {
	parser := NewStructParser()

	type Level3 struct {
		Value string
	}
	type Level2 struct {
		Level3s []Level3
	}
	type Level1 struct {
		Level2s []Level2
	}

	struct1 := Level1{
		Level2s: []Level2{
			{
				Level3s: []Level3{
					{Value: "deep1"},
					{Value: "deep2"},
				},
			},
			{
				Level3s: []Level3{
					{Value: "deep3"},
				},
			},
		},
	}

	fieldMap := parser.ParseToFieldMap(struct1)

	// 检查深层嵌套字段
	expectedFields := []string{
		"Level2s[0].Level3s[0].Value",
		"Level2s[0].Level3s[1].Value",
		"Level2s[1].Level3s[0].Value",
	}

	for _, field := range expectedFields {
		if _, exists := fieldMap[field]; !exists {
			t.Errorf("深层嵌套字段 %s 不存在", field)
		}
	}

	if fieldMap["Level2s[0].Level3s[0].Value"].Value != "deep1" {
		t.Error("深层嵌套字段值不正确")
	}
}

// TestStructParser_ParseToFieldMap_WithDebug 测试调试模式
func TestStructParser_ParseToFieldMap_WithDebug(t *testing.T) {
	parser := NewStructParser()
	parser.SetDebugMode(true)

	struct1 := SimpleStruct{
		Name:  "debug_test",
		Value: 999,
	}

	// 这个测试主要是确保调试模式不会导致崩溃
	fieldMap := parser.ParseToFieldMap(struct1)

	if len(fieldMap) != 2 {
		t.Errorf("调试模式下解析结果不正确，期望 2 个字段，实际 %d 个", len(fieldMap))
	}
}

// BenchmarkStructParser_ParseToFieldMap 性能基准测试
func BenchmarkStructParser_ParseToFieldMap(b *testing.B) {
	parser := NewStructParser()
	struct1 := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name:  "benchmark",
			Value: 123,
		},
		Items: []string{"item1", "item2", "item3", "item4", "item5"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parser.ParseToFieldMap(struct1)
	}
}

// BenchmarkStructParser_ApplyFieldMappings 字段映射性能基准测试
func BenchmarkStructParser_ApplyFieldMappings(b *testing.B) {
	parser := NewStructParser()
	struct1 := NestedStruct{
		ID: 1,
		Simple: SimpleStruct{
			Name:  "benchmark",
			Value: 123,
		},
		Items: []string{"item1", "item2", "item3"},
	}

	sourceFieldMap := parser.ParseToFieldMap(struct1)
	mappings := []FieldMapping{
		{SourcePath: "NewName", TargetPath: "Simple.Name"},
		{SourcePath: "NewItems[]", TargetPath: "Items[]"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parser.ApplyFieldMappings(sourceFieldMap, mappings)
	}
}
