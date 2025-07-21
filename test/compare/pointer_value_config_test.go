package compare

import (
	"testing"
)

// TestIgnorePointerValueDiff 测试IgnorePointerValueDiff配置
func TestIgnorePointerValueDiff(t *testing.T) {
	// 测试结构体
	type TestStruct struct {
		IntPtr  *int `json:"int_ptr"`
		IntVal  int  `json:"int_val"`
		StrPtr  *string `json:"str_ptr"`
		StrVal  string  `json:"str_val"`
	}

	// 测试用例1：指针为nil，值为零值 - 应该相等
	t.Run("NilPointerVsZeroValue", func(t *testing.T) {
		// 使用不同的结构体类型来测试指针和值类型的比较
		type StructWithPtr struct {
			IntPtr *int    `json:"int_field"`
			StrPtr *string `json:"str_field"`
		}
		type StructWithVal struct {
			IntVal int    `json:"int_field"`
			StrVal string `json:"str_field"`
		}

		obj1 := StructWithPtr{
			IntPtr: nil,
			StrPtr: nil,
		}
		obj2 := StructWithVal{
			IntVal: 0,
			StrVal: "",
		}

		// 配置字段映射，将值字段映射到指针字段
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		// 不忽略指针值差异时，应该有差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: false,
		}
		result := CompareResponses(obj1, obj2, "NilPointerVsZeroValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（不忽略指针值差异），但没有发现差异")
		}

		// 忽略指针值差异时，应该没有差异
		options.IgnorePointerValueDiff = true
		result = CompareResponses(obj1, obj2, "NilPointerVsZeroValue", options)
		if result.HasDiff {
			t.Errorf("期望没有差异（忽略指针值差异），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例2：指针不为nil，值相同 - 应该相等
	t.Run("NonNilPointerVsSameValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int    `json:"int_field"`
			StrPtr *string `json:"str_field"`
		}
		type StructWithVal struct {
			IntVal int    `json:"int_field"`
			StrVal string `json:"str_field"`
		}

		intVal := 42
		strVal := "hello"
		obj1 := StructWithPtr{
			IntPtr: &intVal,
			StrPtr: &strVal,
		}
		obj2 := StructWithVal{
			IntVal: 42,
			StrVal: "hello",
		}

		// 配置字段映射，将值字段映射到指针字段
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		// 不忽略指针值差异时，应该有差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: false,
		}
		result := CompareResponses(obj1, obj2, "NonNilPointerVsSameValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（不忽略指针值差异），但没有发现差异")
		}

		// 忽略指针值差异时，应该没有差异
		options.IgnorePointerValueDiff = true
		result = CompareResponses(obj1, obj2, "NonNilPointerVsSameValue", options)
		if result.HasDiff {
			t.Errorf("期望没有差异（忽略指针值差异），但发现了 %d 个差异", len(result.Differences))
			for _, diff := range result.Differences {
				t.Logf("差异: %s", diff.Message)
			}
		}
	})

	// 测试用例3：指针不为nil，值不同 - 应该不相等
	t.Run("NonNilPointerVsDifferentValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int    `json:"int_field"`
			StrPtr *string `json:"str_field"`
		}
		type StructWithVal struct {
			IntVal int    `json:"int_field"`
			StrVal string `json:"str_field"`
		}

		intVal := 42
		strVal := "hello"
		obj1 := StructWithPtr{
			IntPtr: &intVal,
			StrPtr: &strVal,
		}
		obj2 := StructWithVal{
			IntVal: 100,
			StrVal: "world",
		}

		// 配置字段映射，将值字段映射到指针字段
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		// 即使忽略指针值差异，值不同时也应该有差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "NonNilPointerVsDifferentValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（值不同），但没有发现差异")
		}
	})

	// 测试用例4：指针为nil，值不为零值 - 应该不相等
	t.Run("NilPointerVsNonZeroValue", func(t *testing.T) {
		type StructWithPtr struct {
			IntPtr *int    `json:"int_field"`
			StrPtr *string `json:"str_field"`
		}
		type StructWithVal struct {
			IntVal int    `json:"int_field"`
			StrVal string `json:"str_field"`
		}

		obj1 := StructWithPtr{
			IntPtr: nil,
			StrPtr: nil,
		}
		obj2 := StructWithVal{
			IntVal: 42,
			StrVal: "hello",
		}

		// 配置字段映射，将值字段映射到指针字段
		fieldMappings := []FieldMapping{
			{SourcePath: "IntVal", TargetPath: "IntPtr"},
			{SourcePath: "StrVal", TargetPath: "StrPtr"},
		}

		// 即使忽略指针值差异，nil指针和非零值也应该有差异
		options := &ComparisonOptions{
			FieldMappings:          fieldMappings,
			IgnorePointerValueDiff: true,
		}
		result := CompareResponses(obj1, obj2, "NilPointerVsNonZeroValue", options)
		if !result.HasDiff {
			t.Error("期望有差异（nil指针vs非零值），但没有发现差异")
		}
	})
}