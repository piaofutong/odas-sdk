package compare

import (
	"reflect"
)

// DifferenceType 差异类型枚举
type DifferenceType int

const (
	DiffTypeNone               DifferenceType = iota // 无差异
	DiffTypeExtraFieldA                              // A多余字段
	DiffTypeExtraFieldB                              // B多余字段
	DiffTypeValueDifferent                           // 值不相同
	DiffTypeTypeDifferent                            // 类型不同
	DiffTypeCustomComparerDiff                       // 自定义类型比较器返回差异
)

func (d DifferenceType) String() string {
	switch d {
	case DiffTypeNone:
		return "无差异"
	case DiffTypeExtraFieldA:
		return "A多余字段"
	case DiffTypeExtraFieldB:
		return "B多余字段"
	case DiffTypeValueDifferent:
		return "值不相同"
	case DiffTypeTypeDifferent:
		return "类型不同"
	case DiffTypeCustomComparerDiff:
		return "自定义类型比较器返回差异"
	default:
		return "未知差异类型"
	}
}

// ArrayIndex 数组索引信息
type ArrayIndex struct {
	FieldName string // 数组字段名称
	Index     int    // 数组索引
}

// DifferenceDetail 差异详情
type DifferenceDetail struct {
	DiffType     DifferenceType // 差异类型
	FieldNameA   string         // A结构的字段名称
	FieldNameB   string         // B结构的字段名称
	ArrayIndices []ArrayIndex   // 差异数组索引
	ValueA       interface{}    // A结构的值
	ValueB       interface{}    // B结构的值
	TypeA        string         // A结构的类型
	TypeB        string         // B结构的类型
	Message      string         // 差异描述信息
}

// ComparisonResult 比较结果
type ComparisonResult struct {
	TestCaseName string             // 比较用例名称
	Differences  []DifferenceDetail // 差异内容列表
	HasDiff      bool               // 是否存在差异
}

// FieldMapping 字段映射配置
type FieldMapping struct {
	SourcePath string // 源字段路径，如 "A.B[]"
	TargetPath string // 目标字段路径，如 "BA.BB.BC[].BD3"
}

// CustomTypeComparer 自定义类型比较器
type CustomTypeComparer struct {
	SupportedTypes        []string                                                    // 支持比较的类型
	SupportDifferentTypes bool                                                        // 是否支持比较不同类型
	CompareFunc           func(valueA, valueB interface{}) (hasDiff bool, msg string) // 比较方法
}

// ComparisonOptions 比较选项配置
type ComparisonOptions struct {
	FieldMappings       []FieldMapping                 // 字段映射配置
	CustomTypeComparers map[string]*CustomTypeComparer // 自定义类型比较器，key为类型名称
	IgnorePointerValueDiff bool                       // 是否忽略指针和值类型的差异
}

// FieldInfo 字段信息
type FieldInfo struct {
	Path         string       // 字段路径
	Value        interface{}  // 字段值
	Type         reflect.Type // 字段类型
	ArrayIndices []ArrayIndex // 数组索引路径
}
