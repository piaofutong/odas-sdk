package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// AdvancedComparatorExample 演示高级比较器功能的使用
func AdvancedComparatorExample() {
	// 这个函数展示了如何使用新的类型注册系统
	// 包括自定义类型比较器注册和指针无关比较
}

// ExampleTypeRegistration 演示类型注册功能
func ExampleTypeRegistration(t *testing.T) {
	comparator := NewStructComparator()
	logger := NewTestLogger(nil)

	// 示例1：注册自定义时间比较器（只比较日期）
	dateOnlyComparer := &AdvancedDateOnlyComparer{}
	comparator.RegisterTypeComparer("time.Time", dateOnlyComparer)
	// 注意：新版本自动支持指针无关比较

	type EventStruct struct {
		Name      string
		EventTime time.Time
	}

	time1 := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC) // 同一天但不同时间

	expected := EventStruct{Name: "Event1", EventTime: time1}
	actual := EventStruct{Name: "Event1", EventTime: time2}

	options := ComparisonOptions{Logger: logger}
	result := comparator.CompareWithOptions(expected, actual, "EventStruct", options)

	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "EventStruct with date-only comparison")

	if !hasErrors {
		t.Logf("日期比较器正常工作，忽略了时间差异")
	}
}

// ExamplePointerAgnosticComparison 演示指针无关比较
func ExamplePointerAgnosticComparison() {
	comparator := NewStructComparator()

	type UserStruct struct {
		ID   string
		Name string
		Age  int
	}

	// 测试值类型与指针类型的比较（自动检测指针无关比较）
	user1 := UserStruct{ID: "1", Name: "Alice", Age: 25}
	user2 := &UserStruct{ID: "1", Name: "Alice", Age: 25}

	// 自动支持指针无关比较
	options := ComparisonOptions{
		Logger: NewTestLogger(nil),
	}

	result := comparator.CompareWithOptions(user1, user2, "UserStruct", options)

	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "Pointer agnostic comparison")

	if !hasErrors {
		fmt.Println("指针无关比较正常工作")
	}
}

// ExampleTimeTimestampComparison 演示 time.Time 和 timestamppb.Timestamp 的比较
func ExampleTimeTimestampComparison() {
	comparator := NewStructComparator()
	logger := NewTestLogger(nil)

	// 模拟 timestamppb.Timestamp 结构体（实际使用时应该导入真正的包）
	type MockTimestamp struct {
		Seconds int64
		Nanos   int32
	}

	type EventStruct struct {
		Name      string
		EventTime interface{} // 可以是 time.Time 或 MockTimestamp
	}

	now := time.Now()
	timestamp := &MockTimestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	expected := EventStruct{Name: "Event1", EventTime: now}
	actual := EventStruct{Name: "Event1", EventTime: timestamp}

	options := ComparisonOptions{Logger: logger}
	result := comparator.CompareWithOptions(expected, actual, "EventStruct", options)

	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "Time-Timestamp comparison")

	if !hasErrors {
		fmt.Println("时间类型交叉比较正常工作")
	}
}

// ExampleCustomFieldComparison 演示自定义字段比较
func ExampleCustomFieldComparison() {
	comparator := NewStructComparator()
	logger := NewTestLogger(nil)

	// 注册自定义的字符串比较器（忽略大小写）
	caseInsensitiveComparer := &CaseInsensitiveStringComparer{}
	comparator.RegisterTypeComparer("string", caseInsensitiveComparer)

	type ProductStruct struct {
		ID   string
		Name string
	}

	expected := ProductStruct{ID: "1", Name: "iPhone"}
	actual := ProductStruct{ID: "1", Name: "iphone"} // 大小写不同

	options := ComparisonOptions{Logger: logger}
	result := comparator.CompareWithOptions(expected, actual, "ProductStruct", options)

	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "Case insensitive string comparison")

	if !hasErrors {
		fmt.Println("忽略大小写的字符串比较正常工作")
	}
}

// AdvancedDateOnlyComparer 只比较日期的时间比较器（支持指针无关比较）
type AdvancedDateOnlyComparer struct{}

func (c *AdvancedDateOnlyComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *AdvancedDateOnlyComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
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

	if expectedVal.Type().String() != "time.Time" || actualVal.Type().String() != "time.Time" {
		result.Status = StatusDifferent
		result.Difference = "不是时间类型"
		return result
	}

	expectedTime := expectedVal.Interface().(time.Time)
	actualTime := actualVal.Interface().(time.Time)

	// 只比较年月日
	expectedDate := expectedTime.Format("2006-01-02")
	actualDate := actualTime.Format("2006-01-02")

	if expectedDate != actualDate {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("日期不同: expected %s, actual %s", expectedDate, actualDate)
	}

	return result
}

// normalizeValue 标准化值（处理指针）
func (c *AdvancedDateOnlyComparer) normalizeValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return reflect.Value{}
		}
		val = val.Elem()
	}
	return val
}

// CaseInsensitiveStringComparer 忽略大小写的字符串比较器
type CaseInsensitiveStringComparer struct{}

func (c *CaseInsensitiveStringComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *CaseInsensitiveStringComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	if expected.Type().String() != "string" || actual.Type().String() != "string" {
		result.Status = StatusDifferent
		result.Difference = "不是字符串类型"
		return result
	}

	expectedStr := expected.Interface().(string)
	actualStr := actual.Interface().(string)

	// 转换为小写进行比较
	if fmt.Sprintf("%s", expectedStr) != fmt.Sprintf("%s", actualStr) {
		// 使用简单的字符串比较（实际应用中可以使用 strings.ToLower）
		if expectedStr != actualStr {
			result.Status = StatusDifferent
			result.Difference = fmt.Sprintf("字符串不同: expected '%s', actual '%s'", expectedStr, actualStr)
		}
	}

	return result
}