package utils

import (
	"fmt"
	"sort"
	"strings"
)

// ResultPrinter 结果打印器接口
type ResultPrinter interface {
	PrintResult(result ComparisonResult, name string) bool
}

// DefaultResultPrinter 默认结果打印器
type DefaultResultPrinter struct{}

// NewDefaultResultPrinter 创建默认结果打印器
func NewDefaultResultPrinter() *DefaultResultPrinter {
	return &DefaultResultPrinter{}
}

// PrintResult 打印比较结果
func (p *DefaultResultPrinter) PrintResult(result ComparisonResult, structName string) bool {
	hasErrors := result.Status == StatusDifferent

	if hasErrors {
		fmt.Printf("%s比较完成，发现差异:\n", structName)
		p.printDifferencesOnly(result, 0)
	} else {
		fmt.Printf("%s比较完成，所有字段无差异\n", structName)
	}

	return hasErrors
}

// printDifferencesOnly 只打印有差异的字段
func (p *DefaultResultPrinter) printDifferencesOnly(result ComparisonResult, indent int) {
	// 生成缩进
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}

	// 如果当前字段有差异，打印差异信息
	if result.Status == StatusDifferent {
		if result.Difference != "" {
			fmt.Printf("%s- %s: %s\n", indentStr, result.Name, result.Difference)
		} else {
			fmt.Printf("%s- %s: 存在差异\n", indentStr, result.Name)
		}
	}

	// 递归处理子字段
	for _, child := range result.Children {
		if child.Status == StatusDifferent {
			p.printDifferencesOnly(child, indent+1)
		}
	}
}

// SummaryResultPrinter 汇总结果打印器
type SummaryResultPrinter struct{}

// NewSummaryResultPrinter 创建汇总结果打印器
func NewSummaryResultPrinter() *SummaryResultPrinter {
	return &SummaryResultPrinter{}
}

// PrintResult 打印汇总比较结果
func (p *SummaryResultPrinter) PrintResult(result ComparisonResult, summaryName string) bool {
	hasErrors := result.Status == StatusDifferent

	if hasErrors {
		fmt.Printf("相同结构不同包类型，比较完成：字段存在差异\n")
		fmt.Printf(" {\n")
		p.printSummaryResultRecursive(result, "\t\t")
		fmt.Printf(" }\n")
	} else {
		fmt.Printf("%s比较完成，所有字段无差异\n", summaryName)
	}

	return hasErrors
}

// printSummaryResultRecursive 递归打印汇总结果
func (p *SummaryResultPrinter) printSummaryResultRecursive(result ComparisonResult, indent string) {
	if len(result.Children) > 0 {
		// 有子字段的结构体或切片
		if strings.Contains(result.Name, "List") {
			// 切片类型
			fmt.Printf("%s%s [\n", indent, result.Name)
			for _, child := range result.Children {
				p.printSummaryResultRecursive(child, indent+"\t")
			}
			fmt.Printf("%s]\n", indent)
		} else {
			// 结构体类型
			fmt.Printf("%s%s {\n", indent, result.Name)
			for _, child := range result.Children {
				p.printSummaryResultRecursive(child, indent+"\t")
			}
			fmt.Printf("%s}\n", indent)
		}
	} else {
		// 基本字段
		if result.Status == StatusDifferent {
			fmt.Printf("%s%s\t\t存在差异\n", indent, result.Name)
			if result.Difference != "" {
				fmt.Printf("%s\t\t\t%s\n", indent, result.Difference)
			}
		} else {
			fmt.Printf("%s%s\t\t无差异\n", indent, result.Name)
		}
	}
}

// ResponseResultPrinter 响应结果打印器
type ResponseResultPrinter struct{}

// NewResponseResultPrinter 创建响应结果打印器
func NewResponseResultPrinter() *ResponseResultPrinter {
	return &ResponseResultPrinter{}
}

// PrintResult 打印响应比较结果
func (p *ResponseResultPrinter) PrintResult(result ComparisonResult, testName string) bool {
	if result.Status == StatusDifferent {
		fmt.Printf("相同结构不同包类型，比较完成：字段存在差异\n")
		fmt.Printf(" {\n")
		p.printResponseResultRecursive(result, "\t\t")
		fmt.Printf(" }\n")
		return true
	} else {
		// fmt.Printf("%s比较完成，所有字段无差异\n", testName)
		return false
	}
}

// printResponseResultRecursive 递归打印响应比较结果 - 紧凑格式
func (p *ResponseResultPrinter) printResponseResultRecursive(result ComparisonResult, indent string) {
	// 如果是根节点，直接打印子字段
	if strings.Contains(result.Name, "AuthorizationList") || strings.Contains(result.Name, "Response") {
		for _, child := range result.Children {
			p.printResponseResultRecursive(child, indent)
		}
		return
	}

	if len(result.Children) > 0 {
		// 有子字段的结构体或切片
		if strings.Contains(result.Name, "List") {
			// 切片类型 - 汇总显示所有元素的字段状态
			fmt.Printf("%s%s [\n", indent, result.Name)

			// 收集所有字段的状态
			fieldStatusMap := make(map[string][]ComparisonResult)
			for _, child := range result.Children {
				if len(child.Children) > 0 {
					// 如果子元素是结构体，收集其字段
					for _, grandChild := range child.Children {
						fieldStatusMap[grandChild.Name] = append(fieldStatusMap[grandChild.Name], grandChild)
					}
				}
			}

			// 按字段名排序并打印汇总状态
			var fieldNames []string
			for fieldName := range fieldStatusMap {
				fieldNames = append(fieldNames, fieldName)
			}
			sort.Strings(fieldNames)

			for _, fieldName := range fieldNames {
				statuses := fieldStatusMap[fieldName]
				hasDiff := false
				var diffDetails []string

				for i, status := range statuses {
					if status.Status == StatusDifferent {
						hasDiff = true
						if status.Difference != "" {
							diffDetails = append(diffDetails, fmt.Sprintf("%s[%d].%s (%s)", result.Name, i, fieldName, status.Difference))
						}
					}
				}

				if hasDiff {
					fmt.Printf("%s\t%s：存在差异\n", indent, fieldName)
					for _, detail := range diffDetails {
						fmt.Printf("%s\t\t%s\n", indent, detail)
					}
				} else {
					fmt.Printf("%s\t%s：无差异\n", indent, fieldName)
				}
			}

			fmt.Printf("%s]\n", indent)
		} else {
			// 结构体类型
			fmt.Printf("%s%s {\n", indent, result.Name)
			for _, child := range result.Children {
				p.printResponseResultRecursive(child, indent+"\t")
			}
			fmt.Printf("%s}\n", indent)
		}
	} else {
		// 基本字段
		if result.Status == StatusDifferent {
			if strings.HasPrefix(result.Name, "+") {
				// 字段在actual中不存在
				fmt.Printf("%s%s: %s\n", indent, result.Name, result.Difference)
			} else if strings.HasPrefix(result.Name, "-") {
				// 字段在expected中不存在
				fmt.Printf("%s%s: %s\n", indent, result.Name, result.Difference)
			} else if result.Difference != "" {
				fmt.Printf("%s%s：存在差异 (%s)\n", indent, result.Name, result.Difference)
			} else {
				fmt.Printf("%s%s：存在差异\n", indent, result.Name)
			}
		} else {
			fmt.Printf("%s%s：无差异\n", indent, result.Name)
		}
	}
}
