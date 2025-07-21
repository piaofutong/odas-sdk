package compare

import (
	"fmt"
	"reflect"
	"strings"
)

// StructParser 结构体解析器
type StructParser struct {
	debugMode bool
	visited   map[uintptr]bool // 用于检测循环引用
}

// NewStructParser 创建新的结构体解析器
func NewStructParser() *StructParser {
	return &StructParser{
		debugMode: false,
		visited:   make(map[uintptr]bool),
	}
}

// SetDebugMode 设置调试模式
func (p *StructParser) SetDebugMode(debug bool) {
	p.debugMode = debug
}

// ParseToFieldMap 将结构体解析为字段映射表
func (p *StructParser) ParseToFieldMap(data interface{}) map[string]*FieldInfo {
	// 重置visited map，确保每次解析都是干净的状态
	p.visited = make(map[uintptr]bool)

	fieldMap := make(map[string]*FieldInfo)
	value := reflect.ValueOf(data)
	if p.debugMode {
		fmt.Printf("[DEBUG] 开始解析结构体，类型: %s, Kind: %s\n", value.Type(), value.Kind())
	}
	p.parseRecursive(value, "", []ArrayIndex{}, fieldMap)
	if p.debugMode {
		fmt.Printf("[DEBUG] 解析完成，共得到 %d 个字段\n", len(fieldMap))
		for path, field := range fieldMap {
			fmt.Printf("[DEBUG] 字段: %s = %v (类型: %s)\n", path, field.Value, field.Type)
		}
	}
	return fieldMap
}

// parseRecursive 递归解析结构体
func (p *StructParser) parseRecursive(value reflect.Value, currentPath string, arrayIndices []ArrayIndex, fieldMap map[string]*FieldInfo) {
	// 保存原始类型（包括指针类型）
	originalType := value.Type()

	// 处理指针
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			// 对于 nil 指针，记录为 nil 值，但保留原始指针类型
			if currentPath != "" {
				fieldMap[currentPath] = &FieldInfo{
					Path:         currentPath,
					Value:        nil,
					Type:         originalType, // 保留指针类型
					ArrayIndices: copyArrayIndices(arrayIndices),
				}
			}
			return
		}

		// 检测循环引用
		ptr := value.Pointer()
		if p.visited[ptr] {
			// 发现循环引用，记录为特殊值并停止递归
			if currentPath != "" {
				fieldMap[currentPath] = &FieldInfo{
					Path:         currentPath,
					Value:        "<circular_reference>",
					Type:         originalType, // 保留指针类型
					ArrayIndices: copyArrayIndices(arrayIndices),
				}
				if p.debugMode {
					fmt.Printf("[DEBUG] 检测到循环引用: %s\n", currentPath)
				}
			}
			return
		}

		// 标记当前指针为已访问
		p.visited[ptr] = true
		// 在函数结束时清理标记（允许在不同路径上重新访问同一对象）
		defer func() {
			delete(p.visited, ptr)
		}()

		value = value.Elem()
	}

	// 处理接口
	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			if currentPath != "" {
				fieldMap[currentPath] = &FieldInfo{
					Path:         currentPath,
					Value:        nil,
					Type:         value.Type(),
					ArrayIndices: copyArrayIndices(arrayIndices),
				}
			}
			return
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		// 检查是否为特殊类型（如 time.Time），如果是，直接作为值处理
		if p.isSpecialType(value.Type()) {
			if currentPath != "" {
				fieldMap[currentPath] = &FieldInfo{
					Path:         currentPath,
					Value:        value.Interface(),
					Type:         originalType, // 保留原始类型（可能是指针）
					ArrayIndices: copyArrayIndices(arrayIndices),
				}
				if p.debugMode {
					fmt.Printf("[DEBUG] 添加特殊类型字段: %s = %v (类型: %s)\n", currentPath, value.Interface(), originalType)
				}
			}
		} else {
			p.parseStruct(value, currentPath, arrayIndices, fieldMap)
		}
	case reflect.Slice, reflect.Array:
		// 对于切片和数组，即使是空的也要添加字段信息，这样可以比较类型差异
		if currentPath != "" {
			var fieldValue interface{}
			// 如果是空切片或nil切片，字段值设为nil
			if value.Len() == 0 {
				fieldValue = nil
			} else {
				fieldValue = value.Interface()
			}

			fieldMap[currentPath] = &FieldInfo{
				Path:         currentPath,
				Value:        fieldValue,
				Type:         originalType, // 保留原始类型（可能是指针）
				ArrayIndices: copyArrayIndices(arrayIndices),
			}
			if p.debugMode {
				fmt.Printf("[DEBUG] 添加切片/数组字段: %s = %v (类型: %s, 长度: %d)\n", currentPath, fieldValue, originalType, value.Len())
			}
		}
		// 如果切片/数组不为空，继续解析元素
		if value.Len() > 0 {
			p.parseSlice(value, currentPath, arrayIndices, fieldMap)
		}
	default:
		// 基本类型，添加到字段映射表
		if currentPath != "" {
			fieldMap[currentPath] = &FieldInfo{
				Path:         currentPath,
				Value:        value.Interface(),
				Type:         originalType, // 保留原始类型（可能是指针）
				ArrayIndices: copyArrayIndices(arrayIndices),
			}
			if p.debugMode {
				fmt.Printf("[DEBUG] 添加基本类型字段: %s = %v (类型: %s)\n", currentPath, value.Interface(), originalType)
			}
		}
	}
}

// parseStruct 解析结构体
func (p *StructParser) parseStruct(value reflect.Value, currentPath string, arrayIndices []ArrayIndex, fieldMap map[string]*FieldInfo) {
	valueType := value.Type()
	if p.debugMode {
		fmt.Printf("[DEBUG] 解析结构体: %s, 字段数量: %d\n", valueType, value.NumField())
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := valueType.Field(i)

		// 跳过未导出的字段
		if !field.CanInterface() {
			if p.debugMode {
				fmt.Printf("[DEBUG] 跳过未导出字段: %s\n", fieldType.Name)
			}
			continue
		}

		fieldPath := p.buildFieldPath(currentPath, fieldType.Name)
		if p.debugMode {
			fmt.Printf("[DEBUG] 处理字段: %s -> %s\n", fieldType.Name, fieldPath)
		}
		p.parseRecursive(field, fieldPath, arrayIndices, fieldMap)
	}
}

// parseSlice 解析切片或数组
func (p *StructParser) parseSlice(value reflect.Value, currentPath string, arrayIndices []ArrayIndex, fieldMap map[string]*FieldInfo) {
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		elementPath := fmt.Sprintf("%s[%d]", currentPath, i)

		// 创建新的数组索引路径
		newArrayIndices := copyArrayIndices(arrayIndices)
		newArrayIndices = append(newArrayIndices, ArrayIndex{
			FieldName: currentPath,
			Index:     i,
		})

		p.parseRecursive(element, elementPath, newArrayIndices, fieldMap)
	}
}

// buildFieldPath 构建字段路径
func (p *StructParser) buildFieldPath(basePath, fieldName string) string {
	if basePath == "" {
		return fieldName
	}
	return basePath + "." + fieldName
}

// copyArrayIndices 复制数组索引切片
func copyArrayIndices(indices []ArrayIndex) []ArrayIndex {
	if len(indices) == 0 {
		return nil
	}
	result := make([]ArrayIndex, len(indices))
	copy(result, indices)
	return result
}

// isSpecialType 检查是否为特殊类型（应该作为整体值处理而不是递归解析字段）
func (p *StructParser) isSpecialType(t reflect.Type) bool {
	// time.Time 类型
	if t.String() == "time.Time" {
		return true
	}

	// 检查是否所有字段都是未导出的
	if t.Kind() == reflect.Struct {
		hasExportedField := false
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.IsExported() {
				hasExportedField = true
				break
			}
		}
		// 如果没有导出字段，作为特殊类型处理
		if !hasExportedField {
			return true
		}
	}

	return false
}

// ApplyFieldMappings 应用字段映射配置，将源字段映射表转换为目标字段映射表
func (p *StructParser) ApplyFieldMappings(sourceFieldMap map[string]*FieldInfo, mappings []FieldMapping) map[string]*FieldInfo {
	mappedFieldMap := make(map[string]*FieldInfo)

	// 首先复制所有原始字段
	for path, fieldInfo := range sourceFieldMap {
		mappedFieldMap[path] = &FieldInfo{
			Path:         fieldInfo.Path,
			Value:        fieldInfo.Value,
			Type:         fieldInfo.Type,
			ArrayIndices: copyArrayIndices(fieldInfo.ArrayIndices),
		}
	}

	// 应用字段映射
	for _, mapping := range mappings {
		p.applyMapping(sourceFieldMap, mappedFieldMap, mapping)
	}

	return mappedFieldMap
}

// applyMapping 应用单个字段映射
func (p *StructParser) applyMapping(sourceFieldMap, mappedFieldMap map[string]*FieldInfo, mapping FieldMapping) {
	if p.debugMode {
		fmt.Printf("[DEBUG] 应用字段映射: %s -> %s\n", mapping.SourcePath, mapping.TargetPath)
	}

	// 查找匹配的源字段
	matchedFields := p.findMatchingFields(sourceFieldMap, mapping.SourcePath)

	for _, sourceField := range matchedFields {
		// 计算映射后的目标路径
		mappedPath := p.calculateMappedPath(sourceField.Path, mapping.SourcePath, mapping.TargetPath)

		if mappedPath != "" {
			// 创建映射后的字段信息
			mappedField := &FieldInfo{
				Path:         mappedPath,
				Value:        sourceField.Value,
				Type:         sourceField.Type,
				ArrayIndices: p.calculateMappedArrayIndices(sourceField.ArrayIndices, mapping),
			}

			mappedFieldMap[mappedPath] = mappedField
			// 移除原始字段
			delete(mappedFieldMap, sourceField.Path)

			if p.debugMode {
				fmt.Printf("[DEBUG] 字段映射成功: %s -> %s\n", sourceField.Path, mappedPath)
			}
		}
	}
}

// findMatchingFields 查找匹配指定路径模式的字段
func (p *StructParser) findMatchingFields(fieldMap map[string]*FieldInfo, pathPattern string) []*FieldInfo {
	var matchedFields []*FieldInfo

	for _, fieldInfo := range fieldMap {
		if p.pathMatches(fieldInfo.Path, pathPattern) {
			matchedFields = append(matchedFields, fieldInfo)
		}
	}

	return matchedFields
}

// pathMatches 检查字段路径是否匹配路径模式
func (p *StructParser) pathMatches(fieldPath, pathPattern string) bool {
	// 将路径模式中的 [] 替换为实际的数组索引进行匹配
	// 例如："BA.BB.BC[].BD3" 应该匹配 "BA.BB.BC[0].BD3", "BA.BB.BC[1].BD3" 等

	// 如果模式中没有数组标记，直接比较
	if !strings.Contains(pathPattern, "[]") {
		return fieldPath == pathPattern
	}

	// 使用正则表达式风格的匹配
	patternParts := strings.Split(pathPattern, "[]")
	if len(patternParts) < 2 {
		return fieldPath == pathPattern
	}

	// 检查字段路径是否以模式的前缀开始，以模式的后缀结束
	prefix := patternParts[0]
	suffix := patternParts[len(patternParts)-1]

	if !strings.HasPrefix(fieldPath, prefix) {
		return false
	}

	if suffix != "" && !strings.HasSuffix(fieldPath, suffix) {
		return false
	}

	// 检查中间部分是否包含数组索引
	remaining := fieldPath[len(prefix):]
	if suffix != "" {
		remaining = remaining[:len(remaining)-len(suffix)]
	}

	// 简化检查：确保剩余部分包含数组索引格式
	return strings.Contains(remaining, "[")
}

// calculateMappedPath 计算映射后的路径
func (p *StructParser) calculateMappedPath(originalPath, sourcePattern, targetPattern string) string {
	// 如果源模式和目标模式都不包含数组，直接替换
	if !strings.Contains(sourcePattern, "[]") && !strings.Contains(targetPattern, "[]") {
		if originalPath == sourcePattern {
			return targetPattern
		}
		return ""
	}

	// 处理数组映射的情况
	// 例如：originalPath="BA.BB.BC[0].BD3", targetPattern="BA.BB.BC[].BD3", sourcePattern="A.B[]"
	// 应该返回 "A.B[0]"

	if strings.Contains(targetPattern, "[]") {
		// 提取数组索引
		indices := p.extractArrayIndices(originalPath, targetPattern)

		// 将索引应用到源模式
		return p.applyArrayIndices(sourcePattern, indices)
	}

	return ""
}

// extractArrayIndices 从原始路径中提取数组索引
func (p *StructParser) extractArrayIndices(originalPath, pattern string) []int {
	var indices []int

	// 简化实现：查找所有 [数字] 模式
	for i := 0; i < len(originalPath); i++ {
		if originalPath[i] == '[' {
			j := i + 1
			for j < len(originalPath) && originalPath[j] != ']' {
				j++
			}
			if j < len(originalPath) {
				indexStr := originalPath[i+1 : j]
				if index := p.parseIndex(indexStr); index >= 0 {
					indices = append(indices, index)
				}
				i = j
			}
		}
	}

	return indices
}

// parseIndex 解析索引字符串
func (p *StructParser) parseIndex(indexStr string) int {
	index := 0
	for _, char := range indexStr {
		if char >= '0' && char <= '9' {
			index = index*10 + int(char-'0')
		} else {
			return -1
		}
	}
	return index
}

// applyArrayIndices 将数组索引应用到模式中
func (p *StructParser) applyArrayIndices(pattern string, indices []int) string {
	result := pattern
	indexPos := 0

	for i := 0; i < len(result); i++ {
		if i+1 < len(result) && result[i:i+2] == "[]" {
			if indexPos < len(indices) {
				replacement := fmt.Sprintf("[%d]", indices[indexPos])
				result = result[:i] + replacement + result[i+2:]
				i += len(replacement) - 1
				indexPos++
			}
		}
	}

	return result
}

// calculateMappedArrayIndices 计算映射后的数组索引
func (p *StructParser) calculateMappedArrayIndices(originalIndices []ArrayIndex, mapping FieldMapping) []ArrayIndex {
	// 简化实现：保持原有的数组索引，但更新字段名称
	var mappedIndices []ArrayIndex

	for _, index := range originalIndices {
		// 根据映射更新字段名称
		mappedFieldName := p.mapFieldName(index.FieldName, mapping)
		mappedIndices = append(mappedIndices, ArrayIndex{
			FieldName: mappedFieldName,
			Index:     index.Index,
		})
	}

	return mappedIndices
}

// mapFieldName 映射字段名称
func (p *StructParser) mapFieldName(originalFieldName string, mapping FieldMapping) string {
	// 简化实现：如果字段名称匹配目标模式的一部分，则映射到源模式的对应部分
	if strings.Contains(mapping.TargetPath, originalFieldName) {
		// 尝试找到对应的源字段名称
		targetParts := strings.Split(mapping.TargetPath, ".")
		sourceParts := strings.Split(mapping.SourcePath, ".")

		for i, part := range targetParts {
			if strings.HasPrefix(part, originalFieldName) && i < len(sourceParts) {
				return sourceParts[i]
			}
		}
	}

	return originalFieldName
}
