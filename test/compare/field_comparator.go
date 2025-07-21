package compare

import (
	"fmt"
	"reflect"
)

// FieldComparator 字段比较器
type FieldComparator struct {
	debugMode bool
}

// NewFieldComparator 创建新的字段比较器
func NewFieldComparator() *FieldComparator {
	return &FieldComparator{
		debugMode: false,
	}
}

// SetDebugMode 设置调试模式
func (c *FieldComparator) SetDebugMode(debug bool) {
	c.debugMode = debug
}

// CompareFieldMaps 比较两个字段映射表
func (c *FieldComparator) CompareFieldMaps(fieldMapA, fieldMapB map[string]*FieldInfo, options *ComparisonOptions) []DifferenceDetail {
	var differences []DifferenceDetail

	// 创建已处理字段的集合
	processedFields := make(map[string]bool)

	// 比较 A 中的字段
	for pathA, fieldInfoA := range fieldMapA {
		processedFields[pathA] = true

		// 查找 B 中对应的字段
		fieldInfoB, existsInB := fieldMapB[pathA]

		if !existsInB {
			// A 中存在但 B 中不存在的字段
			differences = append(differences, DifferenceDetail{
				DiffType:     DiffTypeExtraFieldA,
				FieldNameA:   pathA,
				FieldNameB:   "",
				ArrayIndices: copyArrayIndices(fieldInfoA.ArrayIndices),
				ValueA:       fieldInfoA.Value,
				ValueB:       nil,
				TypeA:        c.getTypeString(fieldInfoA.Type),
				TypeB:        "",
				Message:      fmt.Sprintf("字段 %s 在结构A中存在，但在结构B中不存在", pathA),
			})
			continue
		}

		// 比较字段值和类型
		diff := c.compareFields(fieldInfoA, fieldInfoB, pathA, pathA, options)
		if diff != nil {
			differences = append(differences, *diff)
		}
	}

	// 检查 B 中存在但 A 中不存在的字段
	for pathB, fieldInfoB := range fieldMapB {
		if !processedFields[pathB] {
			// B 中存在但 A 中不存在的字段
			differences = append(differences, DifferenceDetail{
				DiffType:     DiffTypeExtraFieldB,
				FieldNameA:   "",
				FieldNameB:   pathB,
				ArrayIndices: copyArrayIndices(fieldInfoB.ArrayIndices),
				ValueA:       nil,
				ValueB:       fieldInfoB.Value,
				TypeA:        "",
				TypeB:        c.getTypeString(fieldInfoB.Type),
				Message:      fmt.Sprintf("字段 %s 在结构B中存在，但在结构A中不存在", pathB),
			})
		}
	}

	return differences
}

// compareFields 比较两个字段
func (c *FieldComparator) compareFields(fieldA, fieldB *FieldInfo, pathA, pathB string, options *ComparisonOptions) *DifferenceDetail {
	// 获取完整类型信息（保留指针）
	typeA := c.getTypeString(fieldA.Type)
	typeB := c.getTypeString(fieldB.Type)

	if c.debugMode {
		fmt.Printf("[DEBUG] 比较字段: %s (类型: %s) vs %s (类型: %s)\n", pathA, typeA, pathB, typeB)
	}

	// 检查是否有自定义类型比较器
	if options != nil && options.CustomTypeComparers != nil {
		customResult, handled := c.tryCustomComparers(fieldA, fieldB, pathA, pathB, options.CustomTypeComparers)
		if handled {
			// 自定义比较器已处理，直接返回结果
			return customResult
		}
	}

	// 检查是否应该比较类型
	// 根据需求：只判断最末级字段的值的类型，不判断中间结构类型
	shouldCompareType := c.isLeafField(fieldA) || c.isLeafField(fieldB)

	// 比较类型（仅对最末级字段）
	if shouldCompareType && typeA != typeB {
		// 检查是否忽略指针和值类型的差异
		if options != nil && options.IgnorePointerValueDiff {
			if c.isPointerValueTypePair(fieldA.Type, fieldB.Type) {
				// 这是指针和值类型的组合，使用特殊的比较逻辑
				if c.comparePointerAndValue(fieldA, fieldB) {
					// 指针和值类型被认为相等，跳过类型差异检查
					return nil
				} else {
					// 指针和值类型不相等，报告值差异
					return &DifferenceDetail{
						DiffType:     DiffTypeValueDifferent,
						FieldNameA:   pathA,
						FieldNameB:   pathB,
						ArrayIndices: copyArrayIndices(fieldA.ArrayIndices),
						ValueA:       fieldA.Value,
						ValueB:       fieldB.Value,
						TypeA:        typeA,
						TypeB:        typeB,
						Message:      fmt.Sprintf("指针和值类型的值不同: %s=%v vs %s=%v", pathA, fieldA.Value, pathB, fieldB.Value),
					}
				}
			} else {
				// 不是指针和值类型的组合，正常报告类型差异
				return &DifferenceDetail{
					DiffType:     DiffTypeTypeDifferent,
					FieldNameA:   pathA,
					FieldNameB:   pathB,
					ArrayIndices: copyArrayIndices(fieldA.ArrayIndices),
					ValueA:       fieldA.Value,
					ValueB:       fieldB.Value,
					TypeA:        typeA,
					TypeB:        typeB,
					Message:      fmt.Sprintf("字段类型不同: %s (%s) vs %s (%s)", pathA, typeA, pathB, typeB),
				}
			}
		} else {
			// 不忽略指针和值类型差异，直接报告类型差异（包括指针和值类型的组合）
			return &DifferenceDetail{
				DiffType:     DiffTypeTypeDifferent,
				FieldNameA:   pathA,
				FieldNameB:   pathB,
				ArrayIndices: copyArrayIndices(fieldA.ArrayIndices),
				ValueA:       fieldA.Value,
				ValueB:       fieldB.Value,
				TypeA:        typeA,
				TypeB:        typeB,
				Message:      fmt.Sprintf("字段类型不同: %s (%s) vs %s (%s)", pathA, typeA, pathB, typeB),
			}
		}
	}

	// 比较值
	if !c.valuesEqual(fieldA.Value, fieldB.Value) {
		return &DifferenceDetail{
			DiffType:     DiffTypeValueDifferent,
			FieldNameA:   pathA,
			FieldNameB:   pathB,
			ArrayIndices: copyArrayIndices(fieldA.ArrayIndices),
			ValueA:       fieldA.Value,
			ValueB:       fieldB.Value,
			TypeA:        typeA,
			TypeB:        typeB,
			Message:      fmt.Sprintf("字段值不同: %s=%v vs %s=%v", pathA, fieldA.Value, pathB, fieldB.Value),
		}
	}

	return nil
}

// tryCustomComparers 尝试使用自定义类型比较器
// 返回值：(差异详情, 是否已处理)
func (c *FieldComparator) tryCustomComparers(fieldA, fieldB *FieldInfo, pathA, pathB string, customComparers map[string]*CustomTypeComparer) (*DifferenceDetail, bool) {
	typeA := c.getValueType(fieldA.Type)
	typeB := c.getValueType(fieldB.Type)

	if c.debugMode {
		fmt.Printf("[DEBUG] 尝试自定义比较器: typeA=%s, typeB=%s\n", typeA, typeB)
	}

	// 查找匹配的自定义比较器
	for name, comparer := range customComparers {
		if c.debugMode {
			fmt.Printf("[DEBUG] 检查自定义比较器: %s, 支持类型: %v\n", name, comparer.SupportedTypes)
		}

		if c.isTypeSupported(typeA, comparer.SupportedTypes) {
			// 检查是否支持不同类型比较
			if typeA == typeB || comparer.SupportDifferentTypes {
				if c.isTypeSupported(typeB, comparer.SupportedTypes) || comparer.SupportDifferentTypes {
					if c.debugMode {
						fmt.Printf("[DEBUG] 使用自定义比较器: %s\n", name)
					}
					// 调用自定义比较器
					hasDiff, msg := comparer.CompareFunc(fieldA.Value, fieldB.Value)
					if hasDiff {
						return &DifferenceDetail{
							DiffType:     DiffTypeCustomComparerDiff,
							FieldNameA:   pathA,
							FieldNameB:   pathB,
							ArrayIndices: copyArrayIndices(fieldA.ArrayIndices),
							ValueA:       fieldA.Value,
							ValueB:       fieldB.Value,
							TypeA:        typeA,
							TypeB:        typeB,
							Message:      fmt.Sprintf("自定义比较器检测到差异: %s", msg),
						}, true
					}
					// 自定义比较器认为相同，返回 nil 但标记为已处理
					return nil, true
				}
			}
		}
	}

	// 没有找到匹配的自定义比较器
	return nil, false
}

// isTypeSupported 检查类型是否被支持
func (c *FieldComparator) isTypeSupported(typeName string, supportedTypes []string) bool {
	for _, supportedType := range supportedTypes {
		if typeName == supportedType {
			return true
		}
	}
	return false
}

// getValueType 获取值类型（去除指针）
func (c *FieldComparator) getValueType(t reflect.Type) string {
	// 去除指针
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

// getTypeString 获取类型字符串（保留指针信息）
func (c *FieldComparator) getTypeString(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	return t.String()
}

// valuesEqual 比较两个值是否相等
func (c *FieldComparator) valuesEqual(valueA, valueB interface{}) bool {
	// 处理 nil 值
	if valueA == nil && valueB == nil {
		return true
	}
	if valueA == nil || valueB == nil {
		return false
	}

	// 处理指针值
	valA := reflect.ValueOf(valueA)
	valB := reflect.ValueOf(valueB)

	// 如果是指针，需要比较指针指向的值
	if valA.Kind() == reflect.Ptr && valB.Kind() == reflect.Ptr {
		if valA.IsNil() && valB.IsNil() {
			return true
		}
		if valA.IsNil() || valB.IsNil() {
			return false
		}
		return c.valuesEqual(valA.Elem().Interface(), valB.Elem().Interface())
	}

	// 如果一个是指针一个不是，比较指针指向的值和非指针值
	if valA.Kind() == reflect.Ptr {
		if valA.IsNil() {
			return c.isZeroValue(valB)
		}
		return c.valuesEqual(valA.Elem().Interface(), valueB)
	}
	if valB.Kind() == reflect.Ptr {
		if valB.IsNil() {
			return c.isZeroValue(valA)
		}
		return c.valuesEqual(valueA, valB.Elem().Interface())
	}

	// 使用 reflect.DeepEqual 进行深度比较
	return reflect.DeepEqual(valueA, valueB)
}

// isZeroValue 检查值是否为零值
func (c *FieldComparator) isZeroValue(val reflect.Value) bool {
	return val.IsZero()
}

// isPointerValueTypePair 检查两个类型是否是指针和值类型的组合
func (c *FieldComparator) isPointerValueTypePair(typeA, typeB reflect.Type) bool {
	// 检查一个是指针类型，另一个是对应的值类型
	if typeA.Kind() == reflect.Ptr && typeB.Kind() != reflect.Ptr {
		// typeA是指针，typeB是值类型，检查指针指向的类型是否与typeB相同
		return typeA.Elem() == typeB
	}
	if typeB.Kind() == reflect.Ptr && typeA.Kind() != reflect.Ptr {
		// typeB是指针，typeA是值类型，检查指针指向的类型是否与typeA相同
		return typeB.Elem() == typeA
	}
	return false
}

// IsLeafField 判断字段是否是最末级字段（公开方法，用于测试）
func (c *FieldComparator) IsLeafField(field *FieldInfo) bool {
	return c.isLeafField(field)
}

// isLeafField 判断字段是否是最末级字段
// 最末级字段是指其值不是结构体、切片或数组的字段
func (c *FieldComparator) isLeafField(field *FieldInfo) bool {
	if field.Type == nil {
		return true // nil 类型视为叶子节点
	}

	// 获取去除指针后的类型
	t := field.Type
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 检查是否为基本类型或特殊类型
	switch t.Kind() {
	case reflect.Struct:
		// 对于结构体，检查是否为特殊类型（如 time.Time）
		// 特殊类型被视为叶子节点
		return c.isSpecialTypeByType(t)
	case reflect.Slice, reflect.Array:
		// 检查切片/数组的元素类型
		// 如果元素是叶子节点，则视为叶子节点
		elemType := t.Elem()
		// 去除指针
		for elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
		}

		return c.isLeafField(&FieldInfo{Type: elemType, Value: nil})
	case reflect.Map:
		// Map 不是叶子节点
		return false
	case reflect.Interface:
		// 接口类型需要检查实际值
		if field.Value == nil {
			return true
		}
		// 检查接口实际值的类型
		return c.isLeafField(&FieldInfo{Type: reflect.TypeOf(field.Value), Value: field.Value})
	default:
		// 基本类型（int, string, bool 等）都是叶子节点
		return true
	}
}

// hasSliceElementTypeDifference 检查两个切片/数组类型的元素类型是否不同
func (c *FieldComparator) hasSliceElementTypeDifference(typeA, typeB reflect.Type) bool {
	if typeA == nil || typeB == nil {
		return typeA != typeB
	}

	// 去除指针
	for typeA.Kind() == reflect.Ptr {
		typeA = typeA.Elem()
	}
	for typeB.Kind() == reflect.Ptr {
		typeB = typeB.Elem()
	}

	// 检查是否都是切片或数组
	if (typeA.Kind() != reflect.Slice && typeA.Kind() != reflect.Array) ||
		(typeB.Kind() != reflect.Slice && typeB.Kind() != reflect.Array) {
		return false
	}

	// 比较元素类型
	elemTypeA := typeA.Elem()
	elemTypeB := typeB.Elem()

	// 如果类型完全相同，则没有差异
	if elemTypeA == elemTypeB {
		return false
	}

	// 检查是否为相同名称但不同包路径的类型（这种情况认为是兼容的）
	if c.areTypesCompatible(elemTypeA, elemTypeB) {
		return false
	}

	return true
}

// areTypesCompatible 检查两个类型是否兼容
// 兼容的定义：相同的类型名称，即使来自不同的包路径
func (c *FieldComparator) areTypesCompatible(typeA, typeB reflect.Type) bool {
	if typeA == nil || typeB == nil {
		return false
	}

	// 去除指针
	for typeA.Kind() == reflect.Ptr {
		typeA = typeA.Elem()
	}
	for typeB.Kind() == reflect.Ptr {
		typeB = typeB.Elem()
	}

	// 如果Kind不同，则不兼容
	if typeA.Kind() != typeB.Kind() {
		return false
	}

	// 对于结构体类型，检查类型名称是否相同
	if typeA.Kind() == reflect.Struct {
		nameA := typeA.Name()
		nameB := typeB.Name()

		// 如果类型名称相同，认为是兼容的
		if nameA != "" && nameB != "" && nameA == nameB {
			return true
		}
	}

	return false
}

// isSpecialTypeByType 检查类型是否为特殊类型（如 time.Time）
func (c *FieldComparator) isSpecialTypeByType(t reflect.Type) bool {
	// time.Time 被视为特殊类型，应该作为叶子节点处理
	if t.String() == "time.Time" {
		return true
	}

	// 检查是否为不包含导出字段的结构体
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.IsExported() {
				// 包含导出字段，不是特殊类型
				return false
			}
		}
		// 不包含导出字段，视为特殊类型
		return true
	}

	return false
}

// comparePointerAndValue 比较指针和值类型
// 返回true表示相等，false表示不相等
func (c *FieldComparator) comparePointerAndValue(fieldA, fieldB *FieldInfo) bool {
	valueA := fieldA.Value
	valueB := fieldB.Value
	typeA := fieldA.Type

	var ptrValue interface{}
	var nonPtrValue interface{}

	// 确定哪个是指针，哪个是值
	if typeA.Kind() == reflect.Ptr {
		ptrValue = valueA
		nonPtrValue = valueB
	} else {
		ptrValue = valueB
		nonPtrValue = valueA
	}

	// 处理指针值
	if ptrValue == nil {
		// 指针为nil，检查非指针值是否为零值
		nonPtrVal := reflect.ValueOf(nonPtrValue)
		return nonPtrVal.IsZero()
	}

	// 指针不为nil，由于在解析时指针已经被解引用，
	// ptrValue实际上是指针指向的值，直接比较即可
	return reflect.DeepEqual(ptrValue, nonPtrValue)
}

// PrintDifferences 打印差异信息到控制台（调试用）
func (c *FieldComparator) PrintDifferences(testCaseName string, differences []DifferenceDetail) {
	fmt.Printf("\n=== 比较结果: %s ===\n", testCaseName)
	if len(differences) == 0 {
		fmt.Println("✅ 无差异")
		return
	}

	fmt.Printf("❌ 发现 %d 个差异:\n", len(differences))
	for i, diff := range differences {
		fmt.Printf("\n差异 %d:\n", i+1)
		fmt.Printf("  类型: %s\n", diff.DiffType.String())
		fmt.Printf("  字段A: %s\n", diff.FieldNameA)
		fmt.Printf("  字段B: %s\n", diff.FieldNameB)
		fmt.Printf("  值A: %v\n", diff.ValueA)
		fmt.Printf("  值B: %v\n", diff.ValueB)
		fmt.Printf("  类型A: %s\n", diff.TypeA)
		fmt.Printf("  类型B: %s\n", diff.TypeB)
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
