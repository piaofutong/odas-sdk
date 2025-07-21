package compare

import (
	"reflect"
	"testing"
)

// TestRegressionIssues 回归测试，确保历史问题不再出现
func TestRegressionIssues(t *testing.T) {
	// 回归测试1：指针类型信息保留问题
	// 历史问题：在parseRecursive中解引用指针时丢失了原始指针类型信息
	t.Run("PointerTypePreservation", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int    `json:"int_ptr"`
			StrPtr *string `json:"str_ptr"`
		}

		intVal := 42
		strVal := "hello"
		obj := StructWithPtr{
			IntPtr: &intVal,
			StrPtr: &strVal,
		}

		// 解析结构体
		parser := NewStructParser()
		parser.SetDebugMode(true)
		fieldMap := parser.ParseToFieldMap(obj)

		// 验证指针类型信息被正确保留
		intPtrField, exists := fieldMap["IntPtr"]
		if !exists {
			t.Fatal("IntPtr字段不存在")
		}
		if intPtrField.Type.String() != "*int" {
			t.Errorf("IntPtr类型错误，期望 '*int'，实际 '%s'", intPtrField.Type.String())
		}

		strPtrField, exists := fieldMap["StrPtr"]
		if !exists {
			t.Fatal("StrPtr字段不存在")
		}
		if strPtrField.Type.String() != "*string" {
			t.Errorf("StrPtr类型错误，期望 '*string'，实际 '%s'", strPtrField.Type.String())
		}
	})

	// 回归测试2：字段映射方向错误问题
	// 历史问题：在applyMapping中查找的是TargetPath而不是SourcePath
	t.Run("FieldMappingDirection", func(t *testing.T) {
		type StructA struct {
			FieldA int `json:"field_a"`
		}
		type StructB struct {
			FieldB int `json:"field_b"`
		}

		obj1 := StructA{FieldA: 42}
		obj2 := StructB{FieldB: 42}

		// 配置字段映射：将B的FieldB映射到A的FieldA
		fieldMappings := []FieldMapping{
			{SourcePath: "FieldB", TargetPath: "FieldA"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponsesWithDebug(obj1, obj2, "FieldMappingDirection", options)
		if result.HasDiff {
			t.Errorf("期望无差异（字段映射应该生效），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 回归测试3：字段映射后原始字段未移除问题
	// 历史问题：字段映射后没有从mappedFieldMap中删除原始字段
	t.Run("OriginalFieldRemoval", func(t *testing.T) {
		type StructA struct {
			FieldA int `json:"field_a"`
		}
		type StructB struct {
			FieldB int `json:"field_b"`
		}

		_ = StructA{FieldA: 42}
		obj2 := StructB{FieldB: 42}

		// 解析结构体B
		parser := NewStructParser()
		sourceFieldMap := parser.ParseToFieldMap(obj2)

		// 应用字段映射
		fieldMappings := []FieldMapping{
			{SourcePath: "FieldB", TargetPath: "FieldA"},
		}
		mappedFieldMap := parser.ApplyFieldMappings(sourceFieldMap, fieldMappings)

		// 验证原始字段FieldB被移除，新字段FieldA被添加
		if _, exists := mappedFieldMap["FieldB"]; exists {
			t.Error("原始字段FieldB应该被移除，但仍然存在")
		}
		if _, exists := mappedFieldMap["FieldA"]; !exists {
			t.Error("映射后的字段FieldA应该存在，但不存在")
		}
	})

	// 回归测试4：comparePointerAndValue逻辑错误问题
	// 历史问题：comparePointerAndValue方法的比较逻辑不正确
	t.Run("ComparePointerAndValueLogic", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		// 测试场景1：非nil指针与相同值
		intVal := 42
		obj1 := StructWithPtr{IntPtr: &intVal}
		obj2 := StructWithVal{IntVal: 42}

		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		// 忽略指针值差异时应该无差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "PointerValueSame", options)
		if result.HasDiff {
			t.Errorf("期望无差异（忽略指针值差异，值相同），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}

		// 不忽略指针值差异时应该有差异
		options.IgnorePointerValueDiff = false
		result = CompareResponses(obj1, obj2, "PointerValueSame", options)
		if !result.HasDiff {
			t.Error("期望有差异（不忽略指针值差异），但没有发现差异")
		}

		// 测试场景2：nil指针与零值
		obj3 := StructWithPtr{IntPtr: nil}
		obj4 := StructWithVal{IntVal: 0}

		options.IgnorePointerValueDiff = true
		result = CompareResponses(obj3, obj4, "NilPointerZeroValue", options)
		if result.HasDiff {
			t.Errorf("期望无差异（nil指针vs零值），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}

		// 测试场景3：nil指针与非零值
		obj5 := StructWithPtr{IntPtr: nil}
		obj6 := StructWithVal{IntVal: 42}

		result = CompareResponses(obj5, obj6, "NilPointerNonZeroValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（nil指针vs非零值），但没有发现差异")
		}
	})

	// 回归测试5：isPointerValueTypePair判断错误问题
	// 历史问题：isPointerValueTypePair方法可能判断错误
	t.Run("IsPointerValueTypePairLogic", func(t *testing.T) {
		comparator := NewFieldComparator()

		// 测试指针类型和对应值类型
		intPtrType := reflect.TypeOf((*int)(nil))
		intType := reflect.TypeOf(int(0))
		if !comparator.isPointerValueTypePair(intPtrType, intType) {
			t.Error("*int和int应该被识别为指针值类型对")
		}
		if !comparator.isPointerValueTypePair(intType, intPtrType) {
			t.Error("int和*int应该被识别为指针值类型对")
		}

		// 测试非指针值类型对
		stringType := reflect.TypeOf(string(""))
		if comparator.isPointerValueTypePair(intType, stringType) {
			t.Error("int和string不应该被识别为指针值类型对")
		}

		// 测试两个指针类型
		stringPtrType := reflect.TypeOf((*string)(nil))
		if comparator.isPointerValueTypePair(intPtrType, stringPtrType) {
			t.Error("*int和*string不应该被识别为指针值类型对")
		}

		// 测试两个值类型
		if comparator.isPointerValueTypePair(intType, stringType) {
			t.Error("int和string不应该被识别为指针值类型对")
		}
	})

	// 回归测试6：字段映射calculateMappedPath参数顺序错误问题
	// 历史问题：calculateMappedPath方法的参数顺序可能错误
	t.Run("CalculateMappedPathParameterOrder", func(t *testing.T) {
		parser := NewStructParser()

		// 测试简单路径映射
		sourcePath := "FieldB"
		sourcePattern := "FieldB"
		targetPattern := "FieldA"
		mappedPath := parser.calculateMappedPath(sourcePath, sourcePattern, targetPattern)
		if mappedPath != "FieldA" {
			t.Errorf("映射路径错误，期望 'FieldA'，实际 '%s'", mappedPath)
		}

		// 测试嵌套路径映射
		sourcePath = "Nested.FieldB"
		sourcePattern = "Nested.FieldB"
		targetPattern = "Inner.FieldA"
		mappedPath = parser.calculateMappedPath(sourcePath, sourcePattern, targetPattern)
		if mappedPath != "Inner.FieldA" {
			t.Errorf("嵌套映射路径错误，期望 'Inner.FieldA'，实际 '%s'", mappedPath)
		}
	})

	// 回归测试7：compareFields方法中IgnorePointerValueDiff处理错误问题
	// 历史问题：在compareFields中处理IgnorePointerValueDiff的逻辑可能有误
	t.Run("CompareFieldsIgnorePointerValueDiffHandling", func(t *testing.T) {
		comparator := NewFieldComparator()

		// 创建指针字段和值字段
		intVal := 42
		ptrField := &FieldInfo{
			Path:  "IntPtr",
			Value: intVal, // 注意：这里是解引用后的值
			Type:  reflect.TypeOf((*int)(nil)), // 但类型保留为指针类型
		}
		valField := &FieldInfo{
			Path:  "IntVal",
			Value: 42,
			Type:  reflect.TypeOf(int(0)),
		}

		// 测试忽略指针值差异
		options := &ComparisonOptions{
			IgnorePointerValueDiff: true,
		}
		diff := comparator.compareFields(ptrField, valField, "IntPtr", "IntVal", options)
		if diff != nil {
			t.Errorf("期望无差异（忽略指针值差异），但发现差异: %s", diff.Message)
		}

		// 测试不忽略指针值差异
		options.IgnorePointerValueDiff = false
		diff = comparator.compareFields(ptrField, valField, "IntPtr", "IntVal", options)
		if diff == nil {
			t.Error("期望有差异（不忽略指针值差异），但没有发现差异")
		} else if diff.DiffType != DiffTypeTypeDifferent {
			t.Errorf("期望差异类型为 %s，但得到 %s", DiffTypeTypeDifferent, diff.DiffType)
		}
	})
}

// TestPointerValueDiffScenarios 测试各种指针值差异场景
func TestPointerValueDiffScenarios(t *testing.T) {
	// 测试场景1：指针指向相同值
	t.Run("PointerPointsToSameValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		intVal := 42
		obj1 := StructWithPtr{IntPtr: &intVal}
		obj2 := StructWithVal{IntVal: 42}

		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		// 忽略指针值差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "PointerSameValue", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
		}

		// 不忽略指针值差异
		options.IgnorePointerValueDiff = false
		result = CompareResponses(obj1, obj2, "PointerSameValue", options)
		if !result.HasDiff {
			t.Error("期望有差异，但没有发现差异")
		}
	})

	// 测试场景2：指针指向不同值
	t.Run("PointerPointsToDifferentValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		intVal := 42
		obj1 := StructWithPtr{IntPtr: &intVal}
		obj2 := StructWithVal{IntVal: 100}

		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		// 即使忽略指针值差异，值不同时也应该有差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "PointerDifferentValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（值不同），但没有发现差异")
		}
	})

	// 测试场景3：nil指针与零值
	t.Run("NilPointerVsZeroValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		obj1 := StructWithPtr{IntPtr: nil}
		obj2 := StructWithVal{IntVal: 0}

		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		// 忽略指针值差异时，nil指针和零值应该相等
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "NilPointerZeroValue", options)
		if result.HasDiff {
			t.Errorf("期望无差异（nil指针vs零值），但发现了 %d 个差异", len(result.Differences))
		}
	})

	// 测试场景4：nil指针与非零值
	t.Run("NilPointerVsNonZeroValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		obj1 := StructWithPtr{IntPtr: nil}
		obj2 := StructWithVal{IntVal: 42}

		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		// 即使忽略指针值差异，nil指针和非零值也应该不相等
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "NilPointerNonZeroValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（nil指针vs非零值），但没有发现差异")
		}
	})

	// 测试场景5：字符串指针场景
	t.Run("StringPointerScenarios", func(t *testing.T) {
		type StructWithPtr struct {
			StrPtr *string `json:"str_ptr"`
		}
		type StructWithVal struct {
			StrVal string `json:"str_val"`
		}

		fieldMappings := []FieldMapping{
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		// 子测试：非nil指针与相同字符串
		t.Run("NonNilPointerSameString", func(t *testing.T) {
			strVal := "hello"
			obj1 := StructWithPtr{StrPtr: &strVal}
			obj2 := StructWithVal{StrVal: "hello"}

			options := &ComparisonOptions{
				FieldMappings:          fieldMappings,
				IgnorePointerValueDiff: true,
			}
			result := CompareResponses(obj1, obj2, "StringPointerSame", options)
			if result.HasDiff {
				t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			}
		})

		// 子测试：nil指针与空字符串
		t.Run("NilPointerEmptyString", func(t *testing.T) {
			obj1 := StructWithPtr{StrPtr: nil}
			obj2 := StructWithVal{StrVal: ""}

			options := &ComparisonOptions{
				FieldMappings:          fieldMappings,
				IgnorePointerValueDiff: true,
			}
			result := CompareResponses(obj1, obj2, "StringPointerNilEmpty", options)
			if result.HasDiff {
				t.Errorf("期望无差异（nil指针vs空字符串），但发现了 %d 个差异", len(result.Differences))
			}
		})
	})
}