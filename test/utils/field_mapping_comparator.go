package utils

import (
	"fmt"
	"reflect"
)

// FieldMappingComparator 字段映射比较器
type FieldMappingComparator struct {
	logger Logger
}

// NewFieldMappingComparator 创建新的字段映射比较器
func NewFieldMappingComparator() *FieldMappingComparator {
	return &FieldMappingComparator{}
}

// CompareWithMappings 使用字段映射进行比较
func (c *FieldMappingComparator) CompareWithMappings(sourceResp, targetResp interface{}, testName string, options FieldMappingOptions) ComparisonResult {
	c.logger = options.Logger

	// 将结构体解析为字段映射表
	sourceFieldMap := c.parseStructToFieldMap(sourceResp, "")
	targetFieldMap := c.parseStructToFieldMap(targetResp, "")

	if c.logger != nil {
		c.logger.Logf("源结构解析得到 %d 个字段", len(sourceFieldMap))
		c.logger.Logf("目标结构解析得到 %d 个字段", len(targetFieldMap))
	}

	// 执行字段映射比较
	result := ComparisonResult{
		Name:     testName,
		Status:   StatusEqual,
		Children: make([]ComparisonResult, 0),
	}

	// 遍历所有映射配置进行比较
	for _, mapping := range options.Mappings {
		mappingResult := c.compareMappedFields(sourceFieldMap, targetFieldMap, mapping)
		result.Children = append(result.Children, mappingResult)
		if mappingResult.Status == StatusDifferent {
			result.Status = StatusDifferent
		}
	}

	return result
}

// parseStructToFieldMap 将结构体解析为字段映射表
// 返回的 map 的 key 是字段路径，value 是该路径对应的值
func (c *FieldMappingComparator) parseStructToFieldMap(data interface{}, basePath string) map[string]interface{} {
	fieldMap := make(map[string]interface{})
	c.parseStructRecursive(reflect.ValueOf(data), "", fieldMap)
	return fieldMap
}

// parseStructRecursive 递归解析结构体
func (c *FieldMappingComparator) parseStructRecursive(value reflect.Value, currentPath string, fieldMap map[string]interface{}) {
	// 处理指针
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}

	// 处理接口
	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		c.parseStruct(value, currentPath, fieldMap)
	case reflect.Slice, reflect.Array:
		c.parseSlice(value, currentPath, fieldMap)
	default:
		// 基本类型，添加到字段映射表
		if currentPath != "" {
			fieldMap[currentPath] = value.Interface()
		}
	}
}

// parseStruct 解析结构体
func (c *FieldMappingComparator) parseStruct(value reflect.Value, currentPath string, fieldMap map[string]interface{}) {
	valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := valueType.Field(i)
		
		// 跳过未导出的字段
		if !field.CanInterface() {
			continue
		}

		fieldPath := c.buildFieldPath(currentPath, fieldType.Name)
		c.parseStructRecursive(field, fieldPath, fieldMap)
	}
}

// parseSlice 解析切片或数组
func (c *FieldMappingComparator) parseSlice(value reflect.Value, currentPath string, fieldMap map[string]interface{}) {
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		elementPath := fmt.Sprintf("%s[%d]", currentPath, i)
		c.parseStructRecursive(element, elementPath, fieldMap)
	}
}

// buildFieldPath 构建字段路径
func (c *FieldMappingComparator) buildFieldPath(basePath, fieldName string) string {
	if basePath == "" {
		return fieldName
	}
	return basePath + "." + fieldName
}

// compareMappedFields 比较映射的字段
func (c *FieldMappingComparator) compareMappedFields(sourceFieldMap, targetFieldMap map[string]interface{}, mapping FieldMapping) ComparisonResult {
	result := ComparisonResult{
		Name:   fmt.Sprintf("%s -> %s", mapping.SourcePath, mapping.TargetPath),
		Status: StatusEqual,
	}

	// 获取源字段值
	sourceValue, sourceExists := c.getValueByPath(sourceFieldMap, mapping.SourcePath)
	if !sourceExists {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("源字段路径 %s 不存在", mapping.SourcePath)
		return result
	}

	// 获取目标字段值
	targetValue, targetExists := c.getValueByPath(targetFieldMap, mapping.TargetPath)
	if !targetExists {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("目标字段路径 %s 不存在", mapping.TargetPath)
		return result
	}

	// 比较值
	if !reflect.DeepEqual(sourceValue, targetValue) {
		result.Status = StatusDifferent
		result.Difference = fmt.Sprintf("字段值不匹配: 源=%v, 目标=%v", sourceValue, targetValue)
	}

	if c.logger != nil {
		c.logger.Logf("映射比较: %s -> %s, 结果: %s", mapping.SourcePath, mapping.TargetPath, result.Status.String())
	}

	return result
}

// getValueByPath 根据路径获取字段值
func (c *FieldMappingComparator) getValueByPath(fieldMap map[string]interface{}, path string) (interface{}, bool) {
	// 直接查找路径
	if value, exists := fieldMap[path]; exists {
		return value, true
	}
	
	// 如果直接查找失败，尝试解析路径中的数组索引
	// 例如："A.B[1].C[2].D" 需要找到对应的实际路径
	for key, value := range fieldMap {
		if c.pathMatches(key, path) {
			return value, true
		}
	}
	
	return nil, false
}

// pathMatches 检查两个路径是否匹配（考虑数组索引）
func (c *FieldMappingComparator) pathMatches(actualPath, expectedPath string) bool {
	// 简化实现：直接比较路径
	// 在实际应用中，这里需要更复杂的路径匹配逻辑
	return actualPath == expectedPath
}
