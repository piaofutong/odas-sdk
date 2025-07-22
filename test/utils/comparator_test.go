package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// TestStructComparatorBasic 测试基本的结构体比较功能
func TestStructComparatorBasic(t *testing.T) {
	type TestStruct struct {
		ID   string
		Name string
		Age  int
	}

	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	// 测试相同结构体
	expected := TestStruct{ID: "1", Name: "Alice", Age: 25}
	actual := TestStruct{ID: "1", Name: "Alice", Age: 25}

	options := ComparisonOptions{Logger: logger}
	result := comparator.CompareWithOptions(expected, actual, "TestStruct", options)

	if result.Status != StatusEqual {
		t.Errorf("期望相同的结构体比较结果为相等，但得到: %v", result.Status)
	}

	// 测试不同结构体
	actualDifferent := TestStruct{ID: "2", Name: "Bob", Age: 30}
	result2 := comparator.CompareWithOptions(expected, actualDifferent, "TestStruct", options)

	if result2.Status != StatusDifferent {
		t.Errorf("期望不同的结构体比较结果为不同，但得到: %v", result2.Status)
	}
}

// TestStructComparatorWithIgnoreFields 测试忽略字段功能
func TestStructComparatorWithIgnoreFields(t *testing.T) {
	type TestStruct struct {
		ID   string
		Name string
		Age  int
	}

	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	expected := TestStruct{ID: "1", Name: "Alice", Age: 25}
	actual := TestStruct{ID: "1", Name: "Alice", Age: 30} // 年龄不同

	// 不忽略字段时应该有差异
	options1 := ComparisonOptions{Logger: logger}
	result1 := comparator.CompareWithOptions(expected, actual, "TestStruct", options1)

	if result1.Status != StatusDifferent {
		t.Errorf("期望不同年龄的结构体比较结果为不同，但得到: %v", result1.Status)
	}

	// 忽略Age字段时应该相等
	options2 := ComparisonOptions{
		Logger:       logger,
		IgnoreFields: []string{"Age"},
	}
	result2 := comparator.CompareWithOptions(expected, actual, "TestStruct", options2)

	if result2.Status != StatusEqual {
		t.Errorf("期望忽略Age字段后的结构体比较结果为相等，但得到: %v", result2.Status)
	}
}

// TestStructComparatorWithCustomComparer 测试自定义比较器功能
func TestStructComparatorWithCustomComparer(t *testing.T) {
	type EventStruct struct {
		Name      string
		EventTime time.Time
	}

	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	time1 := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC) // 同一天但不同时间

	expected := EventStruct{Name: "Event1", EventTime: time1}
	actual := EventStruct{Name: "Event1", EventTime: time2}

	// 使用默认比较器时应该有差异
	options1 := ComparisonOptions{Logger: logger}
	result1 := comparator.CompareWithOptions(expected, actual, "EventStruct", options1)

	if result1.Status != StatusDifferent {
		t.Errorf("期望不同时间的结构体比较结果为不同，但得到: %v", result1.Status)
	}

	// 使用自定义日期比较器时应该相等（只比较日期）
	timeComparer := &DateOnlyComparer{}
	options2 := ComparisonOptions{
		Logger: logger,
		CustomComparers: map[string]TypeComparer{
			"EventTime": timeComparer,
		},
	}
	result2 := comparator.CompareWithOptions(expected, actual, "EventStruct", options2)

	if result2.Status != StatusEqual {
		t.Errorf("期望使用日期比较器后的结构体比较结果为相等，但得到: %v", result2.Status)
	}
}

// DateOnlyComparer 只比较日期的自定义比较器
type DateOnlyComparer struct{}

func (c *DateOnlyComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *DateOnlyComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
	result := ComparisonResult{
		Name:   fieldName,
		Status: StatusEqual,
	}

	if expected.Type().String() != "time.Time" || actual.Type().String() != "time.Time" {
		result.Status = StatusDifferent
		result.Difference = "不是时间类型"
		return result
	}

	expectedTime := expected.Interface().(time.Time)
	actualTime := actual.Interface().(time.Time)

	// 只比较年月日
	expectedDate := expectedTime.Format("2006-01-02")
	actualDate := actualTime.Format("2006-01-02")

	if expectedDate != actualDate {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("日期不同: expected %s, actual %s", expectedDate, actualDate)
	}

	return result
}

// TestResultPrinters 测试不同的结果打印器
func TestResultPrinters(t *testing.T) {
	type TestStruct struct {
		ID   string
		Name string
	}

	comparator := NewStructComparator()
	expected := TestStruct{ID: "1", Name: "Alice"}
	actual := TestStruct{ID: "2", Name: "Bob"}

	result := comparator.Compare(expected, actual, "TestStruct")

	// 测试默认打印器
	defaultPrinter := NewDefaultResultPrinter()
	hasErrors1 := defaultPrinter.PrintResult(result, "TestStruct")
	if !hasErrors1 {
		t.Errorf("期望默认打印器检测到差异")
	}

	// 测试汇总打印器
	summaryPrinter := NewSummaryResultPrinter()
	hasErrors2 := summaryPrinter.PrintResult(result, "TestStruct")
	if !hasErrors2 {
		t.Errorf("期望汇总打印器检测到差异")
	}

	// 测试响应打印器
	responsePrinter := NewResponseResultPrinter()
	hasErrors3 := responsePrinter.PrintResult(result, "TestStruct")
	if !hasErrors3 {
		t.Errorf("期望响应打印器检测到差异")
	}
}

// TestBatchComparator 测试批量比较器
func TestBatchComparator(t *testing.T) {
	type UserStruct struct {
		ID   string
		Name string
	}

	batchComparator := NewBatchComparator()
	logger := NewTestLogger(t)

	// 添加比较项
	batchComparator.AddComparison("User1",
		UserStruct{ID: "1", Name: "Alice"},
		UserStruct{ID: "1", Name: "Alice"})

	batchComparator.AddComparison("User2",
		UserStruct{ID: "2", Name: "Bob"},
		UserStruct{ID: "2", Name: "Charlie"}) // 名字不同

	// 执行批量比较
	options := ComparisonOptions{Logger: logger}
	results := batchComparator.CompareAll(options)

	if len(results) != 2 {
		t.Errorf("期望2个比较结果，但得到: %d", len(results))
	}

	// 第一个应该相等
	if results[0].Status != StatusEqual {
		t.Errorf("期望第一个比较结果为相等，但得到: %v", results[0].Status)
	}

	// 第二个应该不同
	if results[1].Status != StatusDifferent {
		t.Errorf("期望第二个比较结果为不同，但得到: %v", results[1].Status)
	}
}

// TestCompatibilityFunctions 测试兼容性函数
func TestCompatibilityFunctions(t *testing.T) {
	type TestStruct struct {
		ID   string
		Name string
	}

	expected := TestStruct{ID: "1", Name: "Alice"}
	actual := TestStruct{ID: "1", Name: "Alice"}

	// 测试 CompareResponsesOO 函数
	CompareResponses(t, expected, actual, "TestResponse")

	// 测试期望有差异的情况
	actualDifferent := TestStruct{ID: "2", Name: "Bob"}
	CompareResponsesWithExpectation(t, expected, actualDifferent, "DifferentResponse", true)
}
