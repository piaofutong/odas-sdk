package compare

import (
	"fmt"
)

// ResponseComparator 响应比较器
type ResponseComparator struct {
	parser     *StructParser
	comparator *FieldComparator
	debugMode  bool
}

// NewResponseComparator 创建新的响应比较器
func NewResponseComparator() *ResponseComparator {
	return &ResponseComparator{
		parser:     NewStructParser(),
		comparator: NewFieldComparator(),
		debugMode:  false,
	}
}

// SetDebugMode 设置调试模式
func (r *ResponseComparator) SetDebugMode(debug bool) {
	r.debugMode = debug
	r.parser.SetDebugMode(debug)
	r.comparator.SetDebugMode(debug)
}

// CompareResponses 比较两个结构体响应
// 参数:
//   - structA: 结构 A
//   - structB: 结构 B
//   - testCaseName: 比较用例名称
//   - options: 可选配置（字段映射配置、自定义类型比较器）
// 返回:
//   - ComparisonResult: 比较结果，包含差异信息
func (r *ResponseComparator) CompareResponses(structA, structB interface{}, testCaseName string, options *ComparisonOptions) *ComparisonResult {
	if r.debugMode {
		fmt.Printf("[DEBUG] 开始比较响应: %s\n", testCaseName)
	}

	// 步骤1: 将结构 A、B 解析成字段映射表
	fieldMapA := r.parser.ParseToFieldMap(structA)
	fieldMapB := r.parser.ParseToFieldMap(structB)

	if r.debugMode {
		fmt.Printf("[DEBUG] 结构A解析得到 %d 个字段\n", len(fieldMapA))
		fmt.Printf("[DEBUG] 结构B解析得到 %d 个字段\n", len(fieldMapB))
	}

	// 步骤2: 复制 fieldMapB 成 fieldMapC，根据字段映射配置修改字段名称
	fieldMapC := fieldMapB
	if options != nil && len(options.FieldMappings) > 0 {
		fieldMapC = r.parser.ApplyFieldMappings(fieldMapB, options.FieldMappings)
		if r.debugMode {
			fmt.Printf("[DEBUG] 应用字段映射后，结构C有 %d 个字段\n", len(fieldMapC))
		}
	}

	// 步骤3: 比较 fieldMapA 和 fieldMapC，生成差异结果
	differences := r.comparator.CompareFieldMaps(fieldMapA, fieldMapC, options)

	// 步骤4: 构建比较结果
	result := &ComparisonResult{
		TestCaseName: testCaseName,
		Differences:  differences,
		HasDiff:      len(differences) > 0,
	}

	// 步骤5: （调试）输出差异结果到控制台
	if r.debugMode {
		r.comparator.PrintDifferences(testCaseName, differences)
	}

	return result
}

// CompareResponses 全局函数，提供便捷的比较接口
func CompareResponses(structA, structB interface{}, testCaseName string, options *ComparisonOptions) *ComparisonResult {
	comparator := NewResponseComparator()
	return comparator.CompareResponses(structA, structB, testCaseName, options)
}

// CompareResponsesWithDebug 带调试模式的比较函数
func CompareResponsesWithDebug(structA, structB interface{}, testCaseName string, options *ComparisonOptions) *ComparisonResult {
	comparator := NewResponseComparator()
	comparator.SetDebugMode(true)
	return comparator.CompareResponses(structA, structB, testCaseName, options)
}

// PrintComparisonResult 打印比较结果
func PrintComparisonResult(result *ComparisonResult) {
	if result == nil {
		fmt.Println("比较结果为空")
		return
	}

	fmt.Printf("\n=== 比较结果: %s ===\n", result.TestCaseName)
	if !result.HasDiff {
		fmt.Println("✅ 无差异")
		return
	}

	fmt.Printf("❌ 发现 %d 个差异:\n", len(result.Differences))
	for i, diff := range result.Differences {
		fmt.Printf("\n差异 %d:\n", i+1)
		fmt.Printf("  类型: %s\n", diff.DiffType.String())
		if diff.FieldNameA != "" {
			fmt.Printf("  字段A: %s\n", diff.FieldNameA)
		}
		if diff.FieldNameB != "" {
			fmt.Printf("  字段B: %s\n", diff.FieldNameB)
		}
		if diff.ValueA != nil {
			fmt.Printf("  值A: %v\n", diff.ValueA)
		}
		if diff.ValueB != nil {
			fmt.Printf("  值B: %v\n", diff.ValueB)
		}
		if diff.TypeA != "" {
			fmt.Printf("  类型A: %s\n", diff.TypeA)
		}
		if diff.TypeB != "" {
			fmt.Printf("  类型B: %s\n", diff.TypeB)
		}
		if len(diff.ArrayIndices) > 0 {
			fmt.Printf("  数组索引: ")
			for j, idx := range diff.ArrayIndices {
				if j > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s[%d]", idx.FieldName, idx.Index)
			}
			fmt.Println()
		}
		fmt.Printf("  描述: %s\n", diff.Message)
	}
	fmt.Println()
}