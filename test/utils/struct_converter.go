package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	idsscommon "gitlab.12301.test/gopkg/idss-go-sdk/proto/gen/idss"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StructConverter 结构体转换器
type StructConverter struct {
	typeConverters map[string]*ConverterConfig
	strictMode     bool // 严格模式：当存在字段差异时是否返回错误
}

// TypeConverter 类型转换器接口
type TypeConverter interface {
	Convert(src reflect.Value) (reflect.Value, error)
}

// ConverterConfig 转换器配置
type ConverterConfig struct {
	Converter       TypeConverter
	PointerAgnostic bool // 是否忽略指针差异，支持指针和值之间的自动转换
}

// NewStructConverter 创建新的结构体转换器
func NewStructConverter() *StructConverter {
	c := &StructConverter{
		typeConverters: make(map[string]*ConverterConfig),
	}

	// 注册默认的类型转换器
	c.RegisterConverter("int64->idss.DateType", &IntToDateTypeConverter{})
	c.RegisterConverter("int64->idss.OrderType", &IntToOrderTypeConverter{})
	// 兼容性注册
	c.RegisterConverter("int64->idsscommon.DateType", &IntToDateTypeConverter{})
	c.RegisterConverter("int64->idsscommon.OrderType", &IntToOrderTypeConverter{})

	c.RegisterConverterWithOptions("time.Time->timestamppb.Timestamp", &TimeToTimestampConverter{}, true)

	return c
}

// RegisterConverter 注册类型转换器
func (c *StructConverter) RegisterConverter(typeMapping string, converter TypeConverter) {
	c.typeConverters[typeMapping] = &ConverterConfig{
		Converter:       converter,
		PointerAgnostic: false,
	}
}

// RegisterConverterWithOptions 注册类型转换器（带选项）
func (c *StructConverter) RegisterConverterWithOptions(typeMapping string, converter TypeConverter, pointerAgnostic bool) {
	c.typeConverters[typeMapping] = &ConverterConfig{
		Converter:       converter,
		PointerAgnostic: pointerAgnostic,
	}
}

// SetStrictMode 设置严格模式
// 严格模式下，当存在字段差异时会返回错误并打印差异信息
// 非严格模式下，只转换匹配的字段，忽略差异
func (c *StructConverter) SetStrictMode(strict bool) {
	c.strictMode = strict
}

// ConvertStruct 转换结构体
// src: 源结构体
// dst: 目标结构体类型（必须是指针类型）
func (c *StructConverter) ConvertStruct(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 检查dst必须是指针类型
	if dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	// 获取dst指向的实际值
	dstElem := dstVal.Elem()
	if !dstElem.CanSet() {
		return fmt.Errorf("dst cannot be set")
	}

	// 处理src为指针的情况
	for srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			return fmt.Errorf("src is nil")
		}
		srcVal = srcVal.Elem()
	}

	return c.convertValue(srcVal, dstElem, "")
}

// convertValue 转换值
func (c *StructConverter) convertValue(src, dst reflect.Value, fieldPath string) error {
	srcType := src.Type()
	dstType := dst.Type()

	// 如果类型完全相同，直接赋值
	if srcType == dstType {
		dst.Set(src)
		return nil
	}

	// 检查是否有自定义转换器
	typeMapping := fmt.Sprintf("%s->%s", srcType.String(), dstType.String())
	if converterConfig, exists := c.typeConverters[typeMapping]; exists {
		convertedVal, err := converterConfig.Converter.Convert(src)
		if err != nil {
			return fmt.Errorf("failed to convert %s: %v", fieldPath, err)
		}
		dst.Set(convertedVal)
		return nil
	}

	// 如果没有找到精确匹配的转换器，尝试查找指针无关的转换器
	if err := c.tryPointerAgnosticConversion(src, dst, fieldPath); err == nil {
		return nil
	} else if err.Error() != "no pointer agnostic converter found" {
		return err
	}

	// 处理不同类型的转换
	switch {
	case src.Kind() == reflect.Struct && dst.Kind() == reflect.Struct:
		return c.convertStruct(src, dst, fieldPath)
	case src.Kind() == reflect.Ptr && dst.Kind() == reflect.Ptr:
		return c.convertPointer(src, dst, fieldPath)
	case src.Kind() == reflect.Slice && dst.Kind() == reflect.Slice:
		return c.convertSlice(src, dst, fieldPath)
	case src.Kind() != reflect.Ptr && dst.Kind() == reflect.Ptr:
		// 值类型转指针类型
		return c.convertValueToPointer(src, dst, fieldPath)
	case src.Kind() == reflect.Ptr && dst.Kind() != reflect.Ptr:
		// 指针类型转值类型
		return c.convertPointerToValue(src, dst, fieldPath)
	default:
		// 尝试直接转换
		if src.Type().ConvertibleTo(dst.Type()) {
			dst.Set(src.Convert(dst.Type()))
			return nil
		}
		return fmt.Errorf("cannot convert %s from %s to %s", fieldPath, srcType, dstType)
	}
}

// convertStruct 转换结构体
func (c *StructConverter) convertStruct(src, dst reflect.Value, fieldPath string) error {
	srcType := src.Type()
	dstType := dst.Type()

	// 收集字段差异信息
	var missingFields []string // 目标结构体中存在但源结构体中不存在的字段
	var extraFields []string   // 源结构体中存在但目标结构体中不存在的字段
	var matchedFields []string // 匹配的字段

	// 检查目标结构体的字段
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		dstFieldVal := dst.Field(i)

		// 跳过不可设置的字段
		if !dstFieldVal.CanSet() {
			continue
		}

		// 在源结构体中查找同名字段
		srcFieldVal, found := c.findField(src, dstField.Name)
		if !found {
			missingFields = append(missingFields, dstField.Name)
			continue
		}

		matchedFields = append(matchedFields, dstField.Name)

		fieldPathNew := fieldPath
		if fieldPathNew != "" {
			fieldPathNew += "."
		}
		fieldPathNew += dstField.Name

		if err := c.convertValue(srcFieldVal, dstFieldVal, fieldPathNew); err != nil {
			return err
		}
	}

	// 检查源结构体中多余的字段
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		_, found := c.findFieldInType(dstType, srcField.Name)
		if !found {
			extraFields = append(extraFields, srcField.Name)
		}
	}

	// 如果存在字段差异，根据严格模式决定是否返回错误
	if len(missingFields) > 0 || len(extraFields) > 0 {
		if c.strictMode {
			return c.buildFieldDifferenceError(srcType, dstType, missingFields, extraFields, matchedFields, fieldPath)
		}
		// 非严格模式下，只打印警告信息但不返回错误
		// c.printFieldDifferenceWarning(srcType, dstType, missingFields, extraFields, matchedFields, fieldPath)
	}

	return nil
}

// findField 在结构体中查找字段
func (c *StructConverter) findField(structVal reflect.Value, fieldName string) (reflect.Value, bool) {
	structType := structVal.Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == fieldName {
			return structVal.Field(i), true
		}
	}
	return reflect.Value{}, false
}

// findFieldInType 在结构体类型中查找字段
func (c *StructConverter) findFieldInType(structType reflect.Type, fieldName string) (reflect.StructField, bool) {
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == fieldName {
			return field, true
		}
	}
	return reflect.StructField{}, false
}

// buildFieldDifferenceInfo 构建字段差异信息
func (c *StructConverter) buildFieldDifferenceInfo(srcType, dstType reflect.Type, missingFields, extraFields, matchedFields []string, fieldPath string) string {
	var diffInfo strings.Builder

	// 构建差异信息
	diffInfo.WriteString(fmt.Sprintf("结构体字段差异：源类型=%s, 目标类型=%s\n", srcType.String(), dstType.String()))
	diffInfo.WriteString(" {\n")

	// 显示结构体名称
	structName := dstType.Name()
	if structName == "" {
		structName = "AnonymousStruct"
	}
	diffInfo.WriteString(fmt.Sprintf("\t\t%s {\n", structName))

	// 显示缺失字段（目标中有但源中没有）
	for _, field := range missingFields {
		diffInfo.WriteString(fmt.Sprintf("\t\t\t+%s: 字段在源结构体中不存在\n", field))
	}

	// 显示匹配字段
	for _, field := range matchedFields {
		diffInfo.WriteString(fmt.Sprintf("\t\t\t%s：无差异\n", field))
	}

	// 显示多余字段（源中有但目标中没有）
	for _, field := range extraFields {
		diffInfo.WriteString(fmt.Sprintf("\t\t\t-%s: 字段在目标结构体中不存在\n", field))
	}

	diffInfo.WriteString("\t\t}\n")
	diffInfo.WriteString(" }")

	return diffInfo.String()
}

// buildFieldDifferenceError 构建字段差异错误信息
func (c *StructConverter) buildFieldDifferenceError(srcType, dstType reflect.Type, missingFields, extraFields, matchedFields []string, fieldPath string) error {
	diffInfo := c.buildFieldDifferenceInfo(srcType, dstType, missingFields, extraFields, matchedFields, fieldPath)
	fmt.Println(diffInfo)
	return fmt.Errorf("字段差异导致转换失败")
}

// printFieldDifferenceWarning 打印字段差异警告信息
func (c *StructConverter) printFieldDifferenceWarning(srcType, dstType reflect.Type, missingFields, extraFields, matchedFields []string, fieldPath string) {
	diffInfo := c.buildFieldDifferenceInfo(srcType, dstType, missingFields, extraFields, matchedFields, fieldPath)
	fmt.Printf("[警告] %s\n", diffInfo)
}

// convertPointer 转换指针
func (c *StructConverter) convertPointer(src, dst reflect.Value, fieldPath string) error {
	if src.IsNil() {
		dst.Set(reflect.Zero(dst.Type()))
		return nil
	}

	// 创建目标类型的新实例
	dstElemType := dst.Type().Elem()
	newDstElem := reflect.New(dstElemType).Elem()

	// 转换元素
	if err := c.convertValue(src.Elem(), newDstElem, fieldPath); err != nil {
		return err
	}

	// 设置指针
	dst.Set(newDstElem.Addr())
	return nil
}

// convertSlice 转换切片
func (c *StructConverter) convertSlice(src, dst reflect.Value, fieldPath string) error {
	if src.IsNil() {
		dst.Set(reflect.Zero(dst.Type()))
		return nil
	}

	srcLen := src.Len()
	dstSlice := reflect.MakeSlice(dst.Type(), srcLen, srcLen)

	for i := 0; i < srcLen; i++ {
		srcElem := src.Index(i)
		dstElem := dstSlice.Index(i)
		indexPath := fmt.Sprintf("%s[%d]", fieldPath, i)
		if err := c.convertValue(srcElem, dstElem, indexPath); err != nil {
			return err
		}
	}

	dst.Set(dstSlice)
	return nil
}

// convertValueToPointer 值类型转指针类型
func (c *StructConverter) convertValueToPointer(src, dst reflect.Value, fieldPath string) error {
	dstElemType := dst.Type().Elem()
	newDstElem := reflect.New(dstElemType).Elem()

	// 检查是否有指针无关的转换器支持这种转换
	if err := c.tryPointerAgnosticConversion(src, dst, fieldPath); err == nil {
		return nil
	}

	// 如果没有指针无关转换器，只允许相同基础类型的转换
	if src.Type() == dstElemType {
		newDstElem.Set(src)
		dst.Set(newDstElem.Addr())
		return nil
	}

	// 尝试直接转换
	if src.Type().ConvertibleTo(dstElemType) {
		newDstElem.Set(src.Convert(dstElemType))
		dst.Set(newDstElem.Addr())
		return nil
	}

	return fmt.Errorf("cannot convert %s from %s to %s", fieldPath, src.Type(), dst.Type())
}

// convertPointerToValue 指针类型转值类型
func (c *StructConverter) convertPointerToValue(src, dst reflect.Value, fieldPath string) error {
	if src.IsNil() {
		return fmt.Errorf("cannot convert nil pointer to value at %s", fieldPath)
	}

	// 检查是否有指针无关的转换器支持这种转换
	if err := c.tryPointerAgnosticConversion(src, dst, fieldPath); err == nil {
		return nil
	}

	// 如果没有指针无关转换器，只允许相同基础类型的转换
	if src.Elem().Type() == dst.Type() {
		dst.Set(src.Elem())
		return nil
	}

	// 尝试直接转换
	if src.Elem().Type().ConvertibleTo(dst.Type()) {
		dst.Set(src.Elem().Convert(dst.Type()))
		return nil
	}

	return fmt.Errorf("cannot convert %s from %s to %s", fieldPath, src.Type(), dst.Type())
}

// 具体的类型转换器实现

// TimeToTimestampConverter 时间转时间戳转换器
type TimeToTimestampConverter struct{}

func (t *TimeToTimestampConverter) Convert(src reflect.Value) (reflect.Value, error) {
	if src.Type() != reflect.TypeOf(time.Time{}) {
		return reflect.Value{}, fmt.Errorf("expected time.Time, got %s", src.Type())
	}

	timeVal := src.Interface().(time.Time)
	timestamp := timestamppb.New(timeVal)
	return reflect.ValueOf(timestamp).Elem(), nil
}

// IntToDateTypeConverter int64转DateType转换器
type IntToDateTypeConverter struct{}

func (i *IntToDateTypeConverter) Convert(src reflect.Value) (reflect.Value, error) {
	if src.Kind() != reflect.Int64 {
		return reflect.Value{}, fmt.Errorf("expected int64, got %s", src.Kind())
	}

	intVal := src.Int()
	var dateType idsscommon.DateType
	switch intVal {
	case 1:
		dateType = idsscommon.DateType_DAILY // API中1对应DAILY(proto中值为0)
	case 2:
		dateType = idsscommon.DateType_MONTHLY // API中2对应MONTHLY(proto中值为1)
	case 3:
		dateType = idsscommon.DateType_YEARLY // API中3对应YEARLY(proto中值为2)
	default:
		dateType = idsscommon.DateType_DAILY
	}

	return reflect.ValueOf(dateType), nil
}

// IntToOrderTypeConverter int64转OrderType转换器
type IntToOrderTypeConverter struct{}

func (i *IntToOrderTypeConverter) Convert(src reflect.Value) (reflect.Value, error) {
	if src.Kind() != reflect.Int64 {
		return reflect.Value{}, fmt.Errorf("expected int64, got %s", src.Kind())
	}

	intVal := src.Int()
	var orderType idsscommon.OrderType
	switch intVal {
	case 1:
		orderType = idsscommon.OrderType_MAIN // API中1对应MAIN(proto中值为0)
	case 2:
		orderType = idsscommon.OrderType_SUB // API中2对应SUB(proto中值为1)
	default:
		orderType = idsscommon.OrderType_MAIN
	}

	return reflect.ValueOf(orderType), nil
}

// tryPointerAgnosticConversion 尝试指针无关的转换
func (c *StructConverter) tryPointerAgnosticConversion(src, dst reflect.Value, fieldPath string) error {
	srcType := src.Type()
	dstType := dst.Type()

	// 获取基础类型（去除指针）
	srcBaseType := srcType
	for srcBaseType.Kind() == reflect.Ptr {
		srcBaseType = srcBaseType.Elem()
	}

	dstBaseType := dstType
	for dstBaseType.Kind() == reflect.Ptr {
		dstBaseType = dstBaseType.Elem()
	}

	// 尝试查找基础类型的转换器
	baseTypeMapping := fmt.Sprintf("%s->%s", srcBaseType.String(), dstBaseType.String())
	if converterConfig, exists := c.typeConverters[baseTypeMapping]; exists && converterConfig.PointerAgnostic {
		// 准备源值（去除指针）
		srcVal := src
		for srcVal.Kind() == reflect.Ptr {
			if srcVal.IsNil() {
				return fmt.Errorf("cannot convert nil pointer at %s", fieldPath)
			}
			srcVal = srcVal.Elem()
		}

		// 执行基础类型转换
		convertedVal, err := converterConfig.Converter.Convert(srcVal)
		if err != nil {
			return fmt.Errorf("failed to convert %s: %v", fieldPath, err)
		}

		// 根据目标类型调整结果（添加或移除指针）
		resultVal := convertedVal
		dstPtrDepth := 0
		tempType := dstType
		for tempType.Kind() == reflect.Ptr {
			dstPtrDepth++
			tempType = tempType.Elem()
		}

		// 为结果添加必要的指针层级
		for i := 0; i < dstPtrDepth; i++ {
			ptrVal := reflect.New(resultVal.Type())
			ptrVal.Elem().Set(resultVal)
			resultVal = ptrVal
		}

		dst.Set(resultVal)
		return nil
	}

	return fmt.Errorf("no pointer agnostic converter found")
}

// ConvertStruct 全局便捷函数
func ConvertStruct(src interface{}, dst interface{}) error {
	converter := NewStructConverter()
	return converter.ConvertStruct(src, dst)
}
