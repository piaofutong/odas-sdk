package utils

import (
	"fmt"
	"reflect"
	"testing"
)

// ComparisonResult 表示比较结果
type ComparisonResult struct {
	Name       string             // 字段名称
	Status     ComparisonStatus   // 比较状态
	Difference string             // 差异详情
	Children   []ComparisonResult // 嵌套字段的比较结果
}

// ComparisonStatus 比较状态枚举
type ComparisonStatus int

const (
	StatusEqual     ComparisonStatus = iota // 无差异
	StatusDifferent                         // 存在差异
)

func (s ComparisonStatus) String() string {
	switch s {
	case StatusEqual:
		return "无差异"
	case StatusDifferent:
		return "存在差异"
	default:
		return "未知状态"
	}
}

// ComparisonOptions 比较选项配置
type ComparisonOptions struct {
	ExpectDifference bool                    // 是否期望有差异
	IgnoreFields     []string                // 忽略的字段名列表
	CustomComparers  map[string]TypeComparer // 自定义类型比较器
	Logger           Logger                  // 日志记录器
	TypeRegistry     TypeRegistry            // 类型注册器
}

// Logger 日志接口
type Logger interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// TestLogger 测试日志实现
type TestLogger struct {
	t *testing.T
}

func NewTestLogger(t *testing.T) *TestLogger {
	return &TestLogger{t: t}
}

func (l *TestLogger) Logf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.t.Errorf(format, args...)
}

// TypeComparer 类型比较器接口
type TypeComparer interface {
	Compare(expected, actual reflect.Value, fieldName string) ComparisonResult
	CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult
}

// TypeRegistry 类型注册器接口
type TypeRegistry interface {
	RegisterTypeComparer(typePattern string, comparer TypeComparer)
	GetTypeComparer(expectedType, actualType reflect.Type) (TypeComparer, bool)
}

// DefaultTypeRegistry 默认类型注册器实现
type DefaultTypeRegistry struct {
	typeComparers map[string]TypeComparer // 类型字符串到比较器的映射
}

// NewDefaultTypeRegistry 创建默认类型注册器
func NewDefaultTypeRegistry() *DefaultTypeRegistry {
	return &DefaultTypeRegistry{
		typeComparers: make(map[string]TypeComparer),
	}
}

// RegisterTypeComparer 注册类型比较器
func (r *DefaultTypeRegistry) RegisterTypeComparer(typePattern string, comparer TypeComparer) {
	r.typeComparers[typePattern] = comparer
}

// GetTypeComparer 获取类型比较器
func (r *DefaultTypeRegistry) GetTypeComparer(expectedType, actualType reflect.Type) (TypeComparer, bool) {
	// 获取所有可能的类型表示
	expectedVariants := r.getTypeVariants(expectedType)
	actualVariants := r.getTypeVariants(actualType)
	
	// 尝试所有可能的组合匹配
	for _, expectedVariant := range expectedVariants {
		for _, actualVariant := range actualVariants {
			// 检查组合匹配
			typeKey := expectedVariant + "|" + actualVariant
			if comparer, exists := r.typeComparers[typeKey]; exists {
				return comparer, true
			}
			
			// 检查反向组合匹配
			reverseTypeKey := actualVariant + "|" + expectedVariant
			if comparer, exists := r.typeComparers[reverseTypeKey]; exists {
				return comparer, true
			}
		}
	}
	
	// 检查单一类型匹配（当两个类型相同时）
	for _, variant := range expectedVariants {
		if comparer, exists := r.typeComparers[variant]; exists {
			// 验证这个比较器是否适用于当前的类型组合
			if r.isCompatibleTypes(expectedType, actualType, variant) {
				return comparer, true
			}
		}
	}
	
	return nil, false
}

// getTypeVariants 获取类型的所有变体（包括指针和非指针形式）
func (r *DefaultTypeRegistry) getTypeVariants(t reflect.Type) []string {
	variants := make([]string, 0, 2)
	
	// 添加原始类型
	variants = append(variants, t.String())
	
	// 添加元素类型（去除指针）
	elemType := r.getElementType(t)
	if elemType.String() != t.String() {
		variants = append(variants, elemType.String())
	}
	
	return variants
}

// isCompatibleTypes 检查两个类型是否兼容给定的比较器模式
func (r *DefaultTypeRegistry) isCompatibleTypes(expectedType, actualType reflect.Type, pattern string) bool {
	// 获取元素类型
	expectedElem := r.getElementType(expectedType)
	actualElem := r.getElementType(actualType)
	
	// 如果模式匹配元素类型，则两个类型的元素类型必须相同
	if pattern == expectedElem.String() || pattern == actualElem.String() {
		return expectedElem.String() == actualElem.String()
	}
	
	return true
}

// getElementType 获取类型的元素类型（去除指针）
func (r *DefaultTypeRegistry) getElementType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}



// Comparator 比较器接口
type Comparator interface {
	Compare(expected, actual interface{}, name string) ComparisonResult
	CompareWithOptions(expected, actual interface{}, name string, options ComparisonOptions) ComparisonResult
}

// StructComparator 结构体比较器实现
type StructComparator struct {
	comparers    map[reflect.Kind]TypeComparer
	printer      ResultPrinter
	typeRegistry TypeRegistry
}

// NewStructComparator 创建新的结构体比较器
func NewStructComparator() *StructComparator {
	c := &StructComparator{
		comparers:    make(map[reflect.Kind]TypeComparer),
		printer:      NewDefaultResultPrinter(),
		typeRegistry: NewDefaultTypeRegistry(),
	}

	// 注册默认比较器
	c.RegisterComparer(reflect.Struct, &StructTypeComparer{})
	c.RegisterComparer(reflect.Slice, &SliceTypeComparer{})
	c.RegisterComparer(reflect.Array, &SliceTypeComparer{})
	c.RegisterComparer(reflect.Ptr, &PointerTypeComparer{})
	c.RegisterComparer(reflect.Interface, &InterfaceTypeComparer{})

	// 基本类型使用默认比较器
	basicComparer := &BasicTypeComparer{}
	for _, kind := range []reflect.Kind{
		reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String,
	} {
		c.RegisterComparer(kind, basicComparer)
	}

	// 注册默认的类型比较器
	c.registerDefaultTypeComparers()

	return c
}

// RegisterComparer 注册类型比较器
func (c *StructComparator) RegisterComparer(kind reflect.Kind, comparer TypeComparer) {
	c.comparers[kind] = comparer
}

// RegisterTypeComparer 注册基于类型字符串的比较器
func (c *StructComparator) RegisterTypeComparer(typePattern string, comparer TypeComparer) {
	c.typeRegistry.RegisterTypeComparer(typePattern, comparer)
}



// registerDefaultTypeComparers 注册默认的类型比较器
func (c *StructComparator) registerDefaultTypeComparers() {
	// 注册时间类型比较器（自动支持指针无关比较）
	timeComparer := &TimeTypeComparer{}
	c.RegisterTypeComparer("time.Time", timeComparer)
	
	// 注册 time.Time 和 timestamppb.Timestamp 的交叉比较
	// 新的智能匹配系统会自动处理所有指针变体和双向匹配
	timeTimestampComparer := &TimeTimestampComparer{}
	c.RegisterTypeComparer("time.Time|timestamppb.Timestamp", timeTimestampComparer)
}

// SetResultPrinter 设置结果打印器
func (c *StructComparator) SetResultPrinter(printer ResultPrinter) {
	c.printer = printer
}

// Compare 比较两个值
func (c *StructComparator) Compare(expected, actual interface{}, name string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, name, ComparisonOptions{})
}

// CompareWithOptions 使用选项比较两个值
func (c *StructComparator) CompareWithOptions(expected, actual interface{}, name string, options ComparisonOptions) ComparisonResult {
	expectedVal := c.normalizeValue(reflect.ValueOf(expected))
	actualVal := c.normalizeValue(reflect.ValueOf(actual))

	// 确保 TypeRegistry 被设置
	if options.TypeRegistry == nil {
		options.TypeRegistry = c.typeRegistry
	}

	if options.Logger != nil {
		options.Logger.Logf("开始比较 %s: expected类型=%v, actual类型=%v", name, expectedVal.Type(), actualVal.Type())
	}

	return c.compareValues(expectedVal, actualVal, name, options)
}

// normalizeValue 标准化值（处理指针）
func (c *StructComparator) normalizeValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return val
		}
		val = val.Elem()
	}
	return val
}

// compareValues 比较两个反射值
func (c *StructComparator) compareValues(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
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

	// 检查类型注册器中的自定义类型比较器
	if typeComparer, exists := c.typeRegistry.GetTypeComparer(expected.Type(), actual.Type()); exists {
		return typeComparer.CompareWithOptions(expected, actual, fieldName, options)
	}

	// 特殊处理：指针无关比较（自动检测）
	if c.shouldUsePointerAgnosticComparison(expected, actual) {
		if result := c.handlePointerAgnosticComparison(expected, actual, fieldName, options); result != nil {
			return *result
		}
	}

	// 特殊处理：值类型与指针类型的比较（兼容旧版本）
	if result := c.handlePointerValueComparison(expected, actual, fieldName, options); result != nil {
		return *result
	}

	// 使用注册的比较器
	expectedKind := expected.Kind()
	actualKind := actual.Kind()

	if expectedKind == actualKind {
		if comparer, exists := c.comparers[expectedKind]; exists {
			return comparer.CompareWithOptions(expected, actual, fieldName, options)
		}
	}

	// 默认使用基本类型比较器
	return c.comparers[reflect.String].CompareWithOptions(expected, actual, fieldName, options)
}

// shouldUsePointerAgnosticComparison 检查是否应该使用指针无关比较
func (c *StructComparator) shouldUsePointerAgnosticComparison(expected, actual reflect.Value) bool {
	// 获取元素类型
	expectedElemType := c.getElementType(expected.Type())
	actualElemType := c.getElementType(actual.Type())
	
	// 如果元素类型相同且有一个是指针类型，则使用指针无关比较
	if expectedElemType == actualElemType {
		expectedIsPtr := expected.Type().Kind() == reflect.Ptr
		actualIsPtr := actual.Type().Kind() == reflect.Ptr
		return expectedIsPtr != actualIsPtr // 一个是指针，一个不是
	}
	return false
}

// getElementType 获取类型的元素类型（去除指针）
func (c *StructComparator) getElementType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// handlePointerAgnosticComparison 处理指针无关比较
func (c *StructComparator) handlePointerAgnosticComparison(expected, actual reflect.Value, fieldName string, options ComparisonOptions) *ComparisonResult {
	// 标准化值（去除指针）
	expectedVal := c.normalizeValue(expected)
	actualVal := c.normalizeValue(actual)
	
	// 处理 nil 值
	if !expectedVal.IsValid() && !actualVal.IsValid() {
		return &ComparisonResult{Name: fieldName, Status: StatusEqual}
	}
	if !expectedVal.IsValid() {
		return &ComparisonResult{
			Name:       fieldName,
			Status:     StatusDifferent,
			Difference: "expected 为 nil，actual 不为 nil",
		}
	}
	if !actualVal.IsValid() {
		return &ComparisonResult{
			Name:       fieldName,
			Status:     StatusDifferent,
			Difference: "expected 不为 nil，actual 为 nil",
		}
	}
	
	// 递归比较标准化后的值
	result := c.compareValues(expectedVal, actualVal, fieldName, options)
	return &result
}

// handlePointerValueComparison 处理指针与值类型的比较
func (c *StructComparator) handlePointerValueComparison(expected, actual reflect.Value, fieldName string, options ComparisonOptions) *ComparisonResult {
	expectedKind := expected.Kind()
	actualKind := actual.Kind()

	if expectedKind == reflect.Struct && actualKind == reflect.Ptr {
		if actual.IsNil() {
			if expected.IsZero() {
				return &ComparisonResult{Name: fieldName, Status: StatusEqual}
			}
			return &ComparisonResult{
				Name:       fieldName,
				Status:     StatusDifferent,
				Difference: fmt.Sprintf("API有值 '%v'，但GRPC为nil", expected.Interface()),
			}
		}
		result := c.compareValues(expected, actual.Elem(), fieldName, options)
		return &result
	}

	if expectedKind == reflect.Ptr && actualKind == reflect.Struct {
		if expected.IsNil() {
			if actual.IsZero() {
				return &ComparisonResult{Name: fieldName, Status: StatusEqual}
			}
			return &ComparisonResult{
				Name:       fieldName,
				Status:     StatusDifferent,
				Difference: fmt.Sprintf("API为nil，但GRPC有值 '%v'", actual.Interface()),
			}
		}
		result := c.compareValues(expected.Elem(), actual, fieldName, options)
		return &result
	}

	return nil
}
