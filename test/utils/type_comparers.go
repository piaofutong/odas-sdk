package utils

import (
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// BasicTypeComparer 基本类型比较器
type BasicTypeComparer struct{}

func (c *BasicTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *BasicTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	// 尝试转换为相同类型进行比较
	if expected.Type().ConvertibleTo(actual.Type()) {
		convertedExpected := expected.Convert(actual.Type())
		if !reflect.DeepEqual(convertedExpected.Interface(), actual.Interface()) {
			result.Status = StatusDifferent
			result.Difference = fmt.Sprintf("'%v' != '%v'", expected.Interface(), actual.Interface())
		}
	} else if actual.Type().ConvertibleTo(expected.Type()) {
		convertedActual := actual.Convert(expected.Type())
		if !reflect.DeepEqual(expected.Interface(), convertedActual.Interface()) {
			result.Status = StatusDifferent
			result.Difference = fmt.Sprintf("'%v' != '%v'", expected.Interface(), actual.Interface())
		}
	} else {
		// 类型不可转换，直接比较字符串表示
		expectedStr := fmt.Sprintf("%v", expected.Interface())
		actualStr := fmt.Sprintf("%v", actual.Interface())
		if expectedStr != actualStr {
			result.Status = StatusDifferent
			result.Difference = fmt.Sprintf("'%v' != '%v'", expected.Interface(), actual.Interface())
		}
	}

	return result
}

// StructTypeComparer 结构体类型比较器
type StructTypeComparer struct{}

func (c *StructTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *StructTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:     fieldName,
		Status:   StatusEqual,
		Children: []ComparisonResult{},
	}

	// 检查是否都是结构体
	if expected.Kind() != reflect.Struct || actual.Kind() != reflect.Struct {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("类型不匹配: expected %v, actual %v", expected.Kind(), actual.Kind())
		return result
	}

	expectedType := expected.Type()
	actualType := actual.Type()

	// 创建字段映射
	actualFields := make(map[string]reflect.Value)
	for i := 0; i < actual.NumField(); i++ {
		field := actualType.Field(i)
		if field.IsExported() {
			actualFields[field.Name] = actual.Field(i)
		}
	}

	// 比较expected中的每个字段
	for i := 0; i < expected.NumField(); i++ {
		field := expectedType.Field(i)
		if !field.IsExported() {
			continue
		}

		fieldName := field.Name

		// 检查是否在忽略列表中
		ignored := false
		for _, ignoredField := range options.IgnoreFields {
			if fieldName == ignoredField {
				ignored = true
				break
			}
		}
		if ignored {
			continue
		}

		expectedFieldVal := expected.Field(i)

		actualFieldVal, exists := actualFields[fieldName]
		if !exists {
			fieldResult := ComparisonResult{
				Name:       fmt.Sprintf("+%s", fieldName),
				Status:     StatusDifferent,
				Difference: "字段在actual中不存在",
			}
			result.Children = append(result.Children, fieldResult)
			result.Status = StatusDifferent
			continue
		}

		// 检查自定义比较器
		var fieldResult ComparisonResult
		if options.CustomComparers != nil {
			if customComparer, exists := options.CustomComparers[fieldName]; exists {
				fieldResult = customComparer.Compare(expectedFieldVal, actualFieldVal, fieldName)
			} else {
				// 检查类型注册器
				if options.TypeRegistry != nil {
					if typeComparer, found := options.TypeRegistry.GetTypeComparer(expectedFieldVal.Type(), actualFieldVal.Type()); found {
						fieldResult = typeComparer.CompareWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
					} else {
						fieldResult = compareFieldRecursiveWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
					}
				} else {
					fieldResult = compareFieldRecursiveWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
				}
			}
		} else {
			// 检查类型注册器
			if options.TypeRegistry != nil {
				if typeComparer, found := options.TypeRegistry.GetTypeComparer(expectedFieldVal.Type(), actualFieldVal.Type()); found {
					fieldResult = typeComparer.CompareWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
				} else {
					fieldResult = compareFieldRecursiveWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
				}
			} else {
				fieldResult = compareFieldRecursiveWithOptions(expectedFieldVal, actualFieldVal, fieldName, options)
			}
		}

		result.Children = append(result.Children, fieldResult)
		if fieldResult.Status == StatusDifferent {
			result.Status = StatusDifferent
		}
	}

	// 检查actual中是否有expected中没有的字段
	expectedFields := make(map[string]bool)
	for i := 0; i < expected.NumField(); i++ {
		field := expectedType.Field(i)
		if field.IsExported() {
			expectedFields[field.Name] = true
		}
	}

	for i := 0; i < actual.NumField(); i++ {
		field := actualType.Field(i)
		if field.IsExported() && !expectedFields[field.Name] {
			// 检查是否在忽略列表中
			ignored := false
			for _, ignoredField := range options.IgnoreFields {
				if field.Name == ignoredField {
					ignored = true
					break
				}
			}
			if ignored {
				continue
			}

			fieldResult := ComparisonResult{
				Name:       fmt.Sprintf("-%s", field.Name),
				Status:     StatusDifferent,
				Difference: "字段在expected中不存在",
			}
			result.Children = append(result.Children, fieldResult)
			result.Status = StatusDifferent
		}
	}

	return result
}

// SliceTypeComparer 切片类型比较器
type SliceTypeComparer struct{}

func (c *SliceTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *SliceTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:     fieldName,
		Status:   StatusEqual,
		Children: []ComparisonResult{},
	}

	expectedLen := expected.Len()
	actualLen := actual.Len()

	// 检查切片元素类型是否兼容
	expectedElemType := expected.Type().Elem()
	actualElemType := actual.Type().Elem()
	
	// 如果元素类型完全不同且不能相互转换，则认为是不同的
	if expectedElemType != actualElemType && 
		!expectedElemType.ConvertibleTo(actualElemType) && 
		!actualElemType.ConvertibleTo(expectedElemType) {
		// 特殊处理：如果都是结构体类型，允许进行字段级比较
		if expectedElemType.Kind() != reflect.Struct || actualElemType.Kind() != reflect.Struct {
			result.Status = StatusDifferent
			result.Difference = fmt.Sprintf("切片元素类型不兼容: expected %v, actual %v", expectedElemType, actualElemType)
			return result
		}
	}

	if expectedLen != actualLen {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("长度不匹配: expected %d, actual %d", expectedLen, actualLen)
		return result
	}

	for i := 0; i < expectedLen; i++ {
		indexName := fmt.Sprintf("[%d]", i)
		indexResult := compareFieldRecursiveWithOptions(expected.Index(i), actual.Index(i), indexName, options)
		result.Children = append(result.Children, indexResult)
		if indexResult.Status == StatusDifferent {
			result.Status = StatusDifferent
		}
	}

	return result
}

// PointerTypeComparer 指针类型比较器
type PointerTypeComparer struct{}

func (c *PointerTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *PointerTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	if expected.IsNil() && actual.IsNil() {
		return result
	}
	if expected.IsNil() {
		result.Status = StatusDifferent
		result.Difference = "expected nil, actual not nil"
		return result
	}
	if actual.IsNil() {
		result.Status = StatusDifferent
		result.Difference = "expected not nil, actual nil"
		return result
	}

	return compareFieldRecursiveWithOptions(expected.Elem(), actual.Elem(), fieldName, options)
}

// InterfaceTypeComparer 接口类型比较器
type InterfaceTypeComparer struct{}

func (c *InterfaceTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *InterfaceTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	// 检查值是否有效
	if !expected.IsValid() && !actual.IsValid() {
		return result
	}
	if !expected.IsValid() {
		result.Status = StatusDifferent
		result.Difference = "expected invalid, actual valid"
		return result
	}
	if !actual.IsValid() {
		result.Status = StatusDifferent
		result.Difference = "expected valid, actual invalid"
		return result
	}

	// 检查是否可以调用IsNil()
	expectedCanBeNil := canBeNil(expected)
	actualCanBeNil := canBeNil(actual)

	if expectedCanBeNil && actualCanBeNil {
		if expected.IsNil() && actual.IsNil() {
			return result
		}
		if expected.IsNil() {
			result.Status = StatusDifferent
			result.Difference = "expected nil, actual not nil"
			return result
		}
		if actual.IsNil() {
			result.Status = StatusDifferent
			result.Difference = "expected not nil, actual nil"
			return result
		}
		return compareFieldRecursiveWithOptions(expected.Elem(), actual.Elem(), fieldName, options)
	} else if expectedCanBeNil {
		if expected.IsNil() {
			result.Status = StatusDifferent
			result.Difference = "expected nil, actual not nil"
			return result
		}
		return compareFieldRecursiveWithOptions(expected.Elem(), actual, fieldName, options)
	} else if actualCanBeNil {
		if actual.IsNil() {
			result.Status = StatusDifferent
			result.Difference = "expected not nil, actual nil"
			return result
		}
		return compareFieldRecursiveWithOptions(expected, actual.Elem(), fieldName, options)
	} else {
		// 都不能为nil，直接比较
		return compareFieldRecursiveWithOptions(expected, actual, fieldName, options)
	}
}

// canBeNil 检查值是否可以为nil
func canBeNil(val reflect.Value) bool {
	return val.Kind() == reflect.Chan || val.Kind() == reflect.Func ||
		val.Kind() == reflect.Interface || val.Kind() == reflect.Map ||
		val.Kind() == reflect.Ptr || val.Kind() == reflect.Slice
}

// TimeTypeComparer 时间类型比较器
type TimeTypeComparer struct{}

func (c *TimeTypeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *TimeTypeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	// 处理指针无关比较
	expectedVal := c.normalizeTimeValue(expected)
	actualVal := c.normalizeTimeValue(actual)

	if !expectedVal.IsValid() || !actualVal.IsValid() {
		result.Status = StatusDifferent
		result.Difference = "无效的时间值"
		return result
	}

	if expectedVal.Type().String() != "time.Time" || actualVal.Type().String() != "time.Time" {
		result.Status = StatusDifferent
		result.Difference = "不是时间类型"
		return result
	}

	expectedTime := expectedVal.Interface().(time.Time)
	actualTime := actualVal.Interface().(time.Time)

	if !expectedTime.Equal(actualTime) {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("时间不同: expected %v, actual %v", expectedTime, actualTime)
	}

	return result
}

// normalizeTimeValue 标准化时间值（处理指针）
func (c *TimeTypeComparer) normalizeTimeValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return reflect.Value{}
		}
		val = val.Elem()
	}
	return val
}

// TimeTimestampComparer time.Time 和 timestamppb.Timestamp 交叉比较器
type TimeTimestampComparer struct{}

func (c *TimeTimestampComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *TimeTimestampComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	// 处理指针无关比较
	expectedVal := c.normalizeValue(expected)
	actualVal := c.normalizeValue(actual)

	if !expectedVal.IsValid() || !actualVal.IsValid() {
		result.Status = StatusDifferent
		result.Difference = "无效的时间值"
		return result
	}

	// 转换为统一的时间格式进行比较
	expectedTime, err1 := c.extractTime(expectedVal)
	actualTime, err2 := c.extractTime(actualVal)

	if err1 != nil || err2 != nil {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("时间提取失败: %v, %v", err1, err2)
		return result
	}

	if !expectedTime.Equal(actualTime) {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("时间不同: expected %v, actual %v", expectedTime, actualTime)
	}

	return result
}

// normalizeValue 标准化值（处理指针）
func (c *TimeTimestampComparer) normalizeValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return reflect.Value{}
		}
		val = val.Elem()
	}
	return val
}

// extractTime 从不同类型中提取时间
func (c *TimeTimestampComparer) extractTime(val reflect.Value) (time.Time, error) {
	typeStr := val.Type().String()

	switch {
	case typeStr == "time.Time":
		return val.Interface().(time.Time), nil
	case typeStr == "*time.Time":
		timestamp := val.Interface().(*time.Time)
		if timestamp == nil {
			return time.Time{}, fmt.Errorf("*time.Time 指针为 nil")
		}
		return *timestamp, nil
	case typeStr == "timestamppb.Timestamp":
		// 值类型的 timestamppb.Timestamp
		if timestamp, ok := val.Interface().(timestamppb.Timestamp); ok {
			return timestamp.AsTime(), nil
		}
		return time.Time{}, fmt.Errorf("无法转换为 timestamppb.Timestamp")
	case typeStr == "*timestamppb.Timestamp":
		// 指针类型的 timestamppb.Timestamp
		if timestamp, ok := val.Interface().(*timestamppb.Timestamp); ok {
			if timestamp == nil {
				return time.Time{}, fmt.Errorf("timestamppb.Timestamp 指针为 nil")
			}
			return timestamp.AsTime(), nil
		}
		return time.Time{}, fmt.Errorf("无法转换为 *timestamppb.Timestamp")
	default:
		return time.Time{}, fmt.Errorf("不支持的时间类型: %s", typeStr)
	}
}

// compareFieldRecursiveWithOptions 使用选项递归比较字段值的辅助函数
func compareFieldRecursiveWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	// 检查是否在忽略列表中
	for _, ignoredField := range options.IgnoreFields {
		if fieldName == ignoredField {
			return ComparisonResult{
				Name:   fieldName,
				Status: StatusEqual,
			}
		}
	}

	// 检查自定义比较器
	if options.CustomComparers != nil {
		if customComparer, exists := options.CustomComparers[fieldName]; exists {
			return customComparer.Compare(expected, actual, fieldName)
		}
	}

	// 处理无效值
	if !expected.IsValid() && !actual.IsValid() {
		return ComparisonResult{Name: fieldName, Status: StatusEqual}
	}
	if !expected.IsValid() {
		return ComparisonResult{
			Name:       fieldName,
			Status:     StatusDifferent,
			Difference: fmt.Sprintf("expected <invalid>, actual %v", actual.Interface()),
		}
	}
	if !actual.IsValid() {
		return ComparisonResult{
			Name:       fieldName,
			Status:     StatusDifferent,
			Difference: fmt.Sprintf("expected %v, actual <invalid>", expected.Interface()),
		}
	}

	// 注意：这里不再有硬编码的时间类型特殊处理
	// 类型注册器的检查应该在调用此函数的地方进行

	expectedKind := expected.Kind()
	actualKind := actual.Kind()

	// 特殊处理指针与值类型之间的比较（兼容旧版本逻辑）
	if expectedKind == reflect.Struct && actualKind == reflect.Ptr {
		// expected是值类型，actual是指针类型
		if actual.IsNil() {
			// actual返回nil指针，检查expected的值是否为零值
			if expected.IsZero() {
				return ComparisonResult{Name: fieldName, Status: StatusEqual} // 都是零值，认为相等
			} else {
				return ComparisonResult{
					Name:       fieldName,
					Status:     StatusDifferent,
					Difference: fmt.Sprintf("expected有值 '%v'，但actual为nil", expected.Interface()),
				}
			}
		} else {
			// actual不为nil，比较结构体内容
			structComparer := &StructTypeComparer{}
			return structComparer.CompareWithOptions(expected, actual.Elem(), fieldName, options)
		}
	}
	if expectedKind == reflect.Ptr && actualKind == reflect.Struct {
		// expected是指针类型，actual是值类型
		if expected.IsNil() {
			// expected返回nil指针，检查actual的值是否为零值
			if actual.IsZero() {
				return ComparisonResult{Name: fieldName, Status: StatusEqual} // 都是零值，认为相等
			} else {
				return ComparisonResult{
					Name:       fieldName,
					Status:     StatusDifferent,
					Difference: fmt.Sprintf("expected为nil，但actual有值 '%v'", actual.Interface()),
				}
			}
		} else {
			// expected不为nil，比较结构体内容
			structComparer := &StructTypeComparer{}
			return structComparer.CompareWithOptions(expected.Elem(), actual, fieldName, options)
		}
	}

	// 处理相同类型
	if expectedKind == actualKind {
		switch expectedKind {
		case reflect.Struct:
			structComparer := &StructTypeComparer{}
			return structComparer.CompareWithOptions(expected, actual, fieldName, options)
		case reflect.Slice, reflect.Array:
			sliceComparer := &SliceTypeComparer{}
			return sliceComparer.CompareWithOptions(expected, actual, fieldName, options)
		case reflect.Ptr:
			ptrComparer := &PointerTypeComparer{}
			return ptrComparer.CompareWithOptions(expected, actual, fieldName, options)
		case reflect.Interface:
			interfaceComparer := &InterfaceTypeComparer{}
			return interfaceComparer.CompareWithOptions(expected, actual, fieldName, options)
		default:
			// 基本类型
			basicComparer := &BasicTypeComparer{}
			return basicComparer.CompareWithOptions(expected, actual, fieldName, options)
		}
	}

	// 类型不同，使用基本比较器
	basicComparer := &BasicTypeComparer{}
	return basicComparer.CompareWithOptions(expected, actual, fieldName, options)
}
