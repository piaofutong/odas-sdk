package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// ExampleUsage 演示面向对象比较器的使用方法
func ExampleUsage() {
	// 这个函数展示了如何使用新的面向对象比较器
	// 在实际测试中，你可以参考这些用法
}

// ExampleCompareStructsOO 演示基本结构体比较
func ExampleCompareStructsOO(t *testing.T) {
	// 创建比较器
	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	// 定义测试数据
	type TestStruct struct {
		ID   string
		Name string
		Age  int
	}

	expected := TestStruct{ID: "1", Name: "Alice", Age: 25}
	actual := TestStruct{ID: "1", Name: "Alice", Age: 25}

	// 基本比较
	result := comparator.Compare(expected, actual, "TestStruct")

	// 使用默认打印器
	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "TestStruct")

	if hasErrors {
		t.Errorf("TestStruct 比较失败，发现差异")
	} else {
		t.Logf("TestStruct 所有字段比较成功")
	}

	// 使用选项进行比较
	options := ComparisonOptions{
		Logger:       logger,
		IgnoreFields: []string{"Age"}, // 忽略年龄字段
	}

	actualDifferent := TestStruct{ID: "1", Name: "Alice", Age: 30} // 年龄不同但会被忽略
	result2 := comparator.CompareWithOptions(expected, actualDifferent, "TestStruct", options)

	hasErrors2 := printer.PrintResult(result2, "TestStruct with ignored fields")
	if !hasErrors2 {
		t.Logf("忽略字段功能正常工作")
	}
}

// ExampleCustomTypeComparer 演示自定义类型比较器
func ExampleCustomTypeComparer(t *testing.T) {
	// 自定义时间比较器，只比较日期部分
	timeComparer := &TimeComparer{}

	comparator := NewStructComparator()
	logger := NewTestLogger(t)

	type EventStruct struct {
		Name      string
		EventTime time.Time
	}

	time1 := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC) // 同一天但不同时间

	expected := EventStruct{Name: "Event1", EventTime: time1}
	actual := EventStruct{Name: "Event1", EventTime: time2}

	// 使用自定义比较器
	options := ComparisonOptions{
		Logger: logger,
		CustomComparers: map[string]TypeComparer{
			"EventTime": timeComparer,
		},
	}

	result := comparator.CompareWithOptions(expected, actual, "EventStruct", options)

	printer := NewDefaultResultPrinter()
	hasErrors := printer.PrintResult(result, "EventStruct with custom comparer")

	if !hasErrors {
		t.Logf("自定义时间比较器正常工作")
	}
}

// TimeComparer 自定义时间比较器，只比较日期部分
type TimeComparer struct{}

func (c *TimeComparer) Compare(expected, actual reflect.Value, fieldName string) ComparisonResult {
	return c.CompareWithOptions(expected, actual, fieldName, ComparisonOptions{})
}

func (c *TimeComparer) CompareWithOptions(expected, actual reflect.Value, fieldName string, options ComparisonOptions) ComparisonResult {
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

// ExampleBatchComparison 演示批量比较
func ExampleBatchComparison(t *testing.T) {
	type UserStruct struct {
		ID   string
		Name string
		Age  int
	}

	// 创建批量比较器
	batchComparator := NewBatchComparator()
	logger := NewTestLogger(t)

	// 添加多个比较项
	batchComparator.AddComparison("User1",
		UserStruct{ID: "1", Name: "Alice", Age: 25},
		UserStruct{ID: "1", Name: "Alice", Age: 25})

	batchComparator.AddComparison("User2",
		UserStruct{ID: "2", Name: "Bob", Age: 30},
		UserStruct{ID: "2", Name: "Bob", Age: 30})

	// 执行批量比较
	options := ComparisonOptions{Logger: logger}
	results := batchComparator.CompareAll(options)

	// 打印汇总结果
	summaryPrinter := NewSummaryResultPrinter()
	overallHasErrors := false
	for _, result := range results {
		hasErrors := summaryPrinter.PrintResult(result, result.Name)
		if hasErrors {
			overallHasErrors = true
		}
	}

	if overallHasErrors {
		t.Errorf("批量比较发现差异")
	} else {
		t.Logf("批量比较全部成功")
	}
}

// BatchComparator 批量比较器
type BatchComparator struct {
	comparisons []BatchComparisonItem
	comparator  Comparator
}

// BatchComparisonItem 批量比较项
type BatchComparisonItem struct {
	Name     string
	Expected interface{}
	Actual   interface{}
}

// NewBatchComparator 创建批量比较器
func NewBatchComparator() *BatchComparator {
	return &BatchComparator{
		comparisons: make([]BatchComparisonItem, 0),
		comparator:  NewStructComparator(),
	}
}

// AddComparison 添加比较项
func (b *BatchComparator) AddComparison(name string, expected, actual interface{}) {
	b.comparisons = append(b.comparisons, BatchComparisonItem{
		Name:     name,
		Expected: expected,
		Actual:   actual,
	})
}

// CompareAll 执行所有比较
func (b *BatchComparator) CompareAll(options ComparisonOptions) []ComparisonResult {
	results := make([]ComparisonResult, 0, len(b.comparisons))

	for _, item := range b.comparisons {
		result := b.comparator.CompareWithOptions(item.Expected, item.Actual, item.Name, options)
		results = append(results, result)
	}

	return results
}
