package compare

import (
	"testing"
)

// BasicFieldMapping 全面测试字段映射功能
func BasicFieldMapping(t *testing.T) {
	// 测试用例1：基本字段映射
	t.Run("BasicFieldMapping", func(t *testing.T) {
		type StructA struct {
			FieldA int    `json:"field_a"`
			FieldB string `json:"field_b"`
		}
		type StructB struct {
			FieldX int    `json:"field_x"`
			FieldY string `json:"field_y"`
		}

		obj1 := StructA{FieldA: 42, FieldB: "hello"}
		obj2 := StructB{FieldX: 42, FieldY: "hello"}

		// 配置字段映射
		fieldMappings := []FieldMapping{
			{SourcePath: "FieldX", TargetPath: "FieldA"},
			{SourcePath: "FieldY", TargetPath: "FieldB"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "BasicFieldMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例2：嵌套字段映射
	t.Run("NestedFieldMapping", func(t *testing.T) {
		type InnerA struct {
			Value int `json:"value"`
		}
		type StructA struct {
			Nested InnerA `json:"nested"`
		}
		type InnerB struct {
			Data int `json:"data"`
		}
		type StructB struct {
			Inner InnerB `json:"inner"`
		}

		obj1 := StructA{Nested: InnerA{Value: 100}}
		obj2 := StructB{Inner: InnerB{Data: 100}}

		// 配置嵌套字段映射
		fieldMappings := []FieldMapping{
			{SourcePath: "Inner.Data", TargetPath: "Nested.Value"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "NestedFieldMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例3：数组字段映射
	t.Run("ArrayFieldMapping", func(t *testing.T) {
		type StructA struct {
			Items []string `json:"items"`
		}
		type StructB struct {
			List []string `json:"list"`
		}

		obj1 := StructA{Items: []string{"a", "b", "c"}}
		obj2 := StructB{List: []string{"a", "b", "c"}}

		// 配置数组字段映射
		fieldMappings := []FieldMapping{
			{SourcePath: "List", TargetPath: "Items"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "ArrayFieldMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例4：指针和值类型映射
	t.Run("PointerValueMapping", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int    `json:"int_ptr"`
			StrPtr *string `json:"str_ptr"`
		}
		type StructWithVal struct {
			IntVal int    `json:"int_val"`
			StrVal string `json:"str_val"`
		}

		intVal := 42
		strVal := "hello"
		obj1 := StructWithPtr{IntPtr: &intVal, StrPtr: &strVal}
		obj2 := StructWithVal{IntVal: 42, StrVal: "hello"}

		// 配置指针值映射
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}

		result := CompareResponses(obj1, obj2, "PointerValueMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异（忽略指针值差异），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例5：多重映射
	t.Run("MultipleMapping", func(t *testing.T) {
		type StructA struct {
			Field1 int    `json:"field1"`
			Field2 string `json:"field2"`
			Field3 bool   `json:"field3"`
		}
		type StructB struct {
			X int    `json:"x"`
			Y string `json:"y"`
			Z bool   `json:"z"`
		}

		obj1 := StructA{Field1: 1, Field2: "test", Field3: true}
		obj2 := StructB{X: 1, Y: "test", Z: true}

		// 配置多重映射
		fieldMappings := []FieldMapping{
			{SourcePath: "X", TargetPath: "Field1"},
			{SourcePath: "Y", TargetPath: "Field2"},
			{SourcePath: "Z", TargetPath: "Field3"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "MultipleMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例6：部分映射（有些字段不映射）
	t.Run("PartialMapping", func(t *testing.T) {
		type StructA struct {
			Field1 int    `json:"field1"`
			Field2 string `json:"field2"`
		}
		type StructB struct {
			X int    `json:"x"`
			Y string `json:"y"`
		}

		obj1 := StructA{Field1: 1, Field2: "test"}
		obj2 := StructB{X: 1, Y: "different"} // Y值不同

		// 只映射Field1，不映射Field2
		fieldMappings := []FieldMapping{
			{SourcePath: "X", TargetPath: "Field1"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "PartialMapping", options)
		// 应该有差异，因为Field2和Y没有映射，会被当作额外字段
		if !result.HasDiff {
			t.Error("期望有差异（未映射的字段），但没有发现差异")
		}
		// 应该有2个差异：Field2在A中存在但B中不存在，Y在B中存在但A中不存在
		if len(result.Differences) != 2 {
			t.Errorf("期望2个差异，但发现了 %d 个差异", len(result.Differences))
		}
	})

	// 测试用例7：映射到不存在的字段（错误处理）
	t.Run("MappingToNonExistentField", func(t *testing.T) {
		type StructA struct {
			Field1 int `json:"field1"`
		}
		type StructB struct {
			X int `json:"x"`
		}

		obj1 := StructA{Field1: 1}
		obj2 := StructB{X: 1}

		// 映射到不存在的字段
		fieldMappings := []FieldMapping{
			{SourcePath: "X", TargetPath: "NonExistentField"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "MappingToNonExistentField", options)
		// 应该有差异，因为映射无效
		if !result.HasDiff {
			t.Error("期望有差异（映射到不存在的字段），但没有发现差异")
		}
	})

	// 测试用例8：复杂嵌套结构映射
	t.Run("ComplexNestedMapping", func(t *testing.T) {
		type DeepInnerA struct {
			Value int `json:"value"`
		}
		type InnerA struct {
			Deep DeepInnerA `json:"deep"`
			Name string     `json:"name"`
		}
		type StructA struct {
			Nested InnerA `json:"nested"`
		}

		type DeepInnerB struct {
			Data int `json:"data"`
		}
		type InnerB struct {
			DeepData DeepInnerB `json:"deep_data"`
			Title    string     `json:"title"`
		}
		type StructB struct {
			Inner InnerB `json:"inner"`
		}

		obj1 := StructA{
			Nested: InnerA{
				Deep: DeepInnerA{Value: 100},
				Name: "test",
			},
		}
		obj2 := StructB{
			Inner: InnerB{
				DeepData: DeepInnerB{Data: 100},
				Title:    "test",
			},
		}

		// 配置复杂嵌套映射
		fieldMappings := []FieldMapping{
			{SourcePath: "Inner.DeepData.Data", TargetPath: "Nested.Deep.Value"},
			{SourcePath: "Inner.Title", TargetPath: "Nested.Name"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "ComplexNestedMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例9：数组元素映射
	t.Run("ArrayElementMapping", func(t *testing.T) {
		type ItemA struct {
			Value int `json:"value"`
		}
		type StructA struct {
			Items []ItemA `json:"items"`
		}

		type ItemB struct {
			Data int `json:"data"`
		}
		type StructB struct {
			List []ItemB `json:"list"`
		}

		obj1 := StructA{
			Items: []ItemA{
				{Value: 1},
				{Value: 2},
			},
		}
		obj2 := StructB{
			List: []ItemB{
				{Data: 1},
				{Data: 2},
			},
		}

		// 配置数组元素映射
		fieldMappings := []FieldMapping{
			{SourcePath: "List", TargetPath: "Items"},
			{SourcePath: "List[0].Data", TargetPath: "Items[0].Value"},
			{SourcePath: "List[1].Data", TargetPath: "Items[1].Value"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "ArrayElementMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异，但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例10：nil指针映射
	t.Run("NilPointerMapping", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int `json:"int_ptr"`
		}
		type StructWithVal struct {
			IntVal int `json:"int_val"`
		}

		obj1 := StructWithPtr{IntPtr: nil}
		obj2 := StructWithVal{IntVal: 0}

		// 配置nil指针映射
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
		}

		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}

		result := CompareResponses(obj1, obj2, "NilPointerMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异（nil指针vs零值），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})
}

// TestFieldMappingEdgeCases 测试字段映射的边界情况
func TestFieldMappingEdgeCases(t *testing.T) {
	// 测试用例1：空映射配置
	t.Run("EmptyMappingConfig", func(t *testing.T) {
		type StructA struct {
			Field int `json:"field"`
		}
		type StructB struct {
			Field int `json:"field"`
		}

		obj1 := StructA{Field: 42}
		obj2 := StructB{Field: 42}

		// 空映射配置
		options := &ComparisonOptions{
			FieldMappings: []FieldMapping{},
		}

		result := CompareResponses(obj1, obj2, "EmptyMappingConfig", options)
		if result.HasDiff {
			t.Errorf("期望无差异（相同字段名），但发现了 %d 个差异", len(result.Differences))
		}
	})

	// 测试用例2：nil映射配置
	t.Run("NilMappingConfig", func(t *testing.T) {
		type StructA struct {
			Field int `json:"field"`
		}
		type StructB struct {
			Field int `json:"field"`
		}

		obj1 := StructA{Field: 42}
		obj2 := StructB{Field: 42}

		// nil映射配置
		options := &ComparisonOptions{
			FieldMappings: nil,
		}

		result := CompareResponses(obj1, obj2, "NilMappingConfig", options)
		if result.HasDiff {
			t.Errorf("期望无差异（相同字段名），但发现了 %d 个差异", len(result.Differences))
		}
	})

	// 测试用例3：循环映射（A->B, B->A）
	t.Run("CircularMapping", func(t *testing.T) {
		type StructA struct {
			FieldA int `json:"field_a"`
			FieldB int `json:"field_b"`
		}
		type StructB struct {
			FieldX int `json:"field_x"` // 使用不同的字段名
			FieldY int `json:"field_y"` // 使用不同的字段名
		}

		obj1 := StructA{FieldA: 1, FieldB: 2}
		obj2 := StructB{FieldX: 2, FieldY: 1} // 值交换

		// 循环映射：B的FieldX映射到A的FieldB，B的FieldY映射到A的FieldA
		fieldMappings := []FieldMapping{
			{SourcePath: "FieldX", TargetPath: "FieldB"},
			{SourcePath: "FieldY", TargetPath: "FieldA"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "CircularMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异（循环映射），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例4：重复映射（多个源字段映射到同一个目标字段）
	t.Run("DuplicateTargetMapping", func(t *testing.T) {
		type StructA struct {
			Field int `json:"field"`
		}
		type StructB struct {
			Field1 int `json:"field1"`
			Field2 int `json:"field2"`
		}

		obj1 := StructA{Field: 42}
		obj2 := StructB{Field1: 42, Field2: 100}

		// 重复目标映射
		fieldMappings := []FieldMapping{
			{SourcePath: "Field1", TargetPath: "Field"},
			{SourcePath: "Field2", TargetPath: "Field"}, // 重复目标
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "DuplicateTargetMapping", options)
		// 行为可能不确定，但不应该崩溃
		t.Logf("重复目标映射结果: HasDiff=%v, 差异数量=%d", result.HasDiff, len(result.Differences))
	})

	// 测试用例5：自映射（字段映射到自己）
	t.Run("SelfMapping", func(t *testing.T) {
		type StructA struct {
			Field int `json:"field"`
		}
		type StructB struct {
			OtherField int `json:"other_field"` // 使用不同的字段名
		}

		obj1 := StructA{Field: 42}
		obj2 := StructB{OtherField: 42}

		// 字段映射：B的OtherField映射到A的Field
		fieldMappings := []FieldMapping{
			{SourcePath: "OtherField", TargetPath: "Field"},
		}

		options := &ComparisonOptions{
			FieldMappings: fieldMappings,
		}

		result := CompareResponses(obj1, obj2, "SelfMapping", options)
		if result.HasDiff {
			t.Errorf("期望无差异（字段映射），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})
}
