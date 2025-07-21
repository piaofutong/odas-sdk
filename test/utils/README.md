# Compare 包

这是一个通用的结构体比较工具包，专门用于测试场景中比较两个结构体的字段值。

## 功能特性

- **跨类型比较**：支持不同类型但字段结构相同的结构体比较
- **深度递归**：支持嵌套结构体、切片、数组、指针、接口等复杂类型
- **详细报告**：提供详细的差异报告，精确定位不匹配的字段
- **灵活匹配**：按字段名进行匹配，而非严格的类型匹配
- **测试友好**：专为Go测试框架设计，集成testing.T

## 主要函数

### CompareStructs

```go
func CompareStructs(t *testing.T, expected, actual interface{}, structName string)
```

通用的结构体比较函数，支持不同类型但字段结构相同的比较。

**参数：**
- `t *testing.T`：测试上下文
- `expected interface{}`：期望的结构体
- `actual interface{}`：实际的结构体
- `structName string`：结构体名称，用于日志输出

**特性：**
- 自动处理指针解引用
- 智能类型检测（区分相同结构不同包 vs 完全不同类型）
- 详细的比较日志输出
- 失败时提供精确的差异报告

## 使用示例

```go
package main

import (
    "testing"
    "github.com/piaofutong/odas-sdk/utils/compare"
)

type User struct {
    ID   int
    Name string
    Age  int
}

type UserResponse struct {
    ID   int
    Name string
    Age  int
}

func TestUserComparison(t *testing.T) {
    expected := User{ID: 1, Name: "Alice", Age: 30}
    actual := UserResponse{ID: 1, Name: "Alice", Age: 30}
    
    // 比较不同类型但字段相同的结构体
    compare.CompareStructs(t, expected, actual, "User")
}

func TestComplexStructure(t *testing.T) {
    type Address struct {
        Street string
        City   string
    }
    
    type Person struct {
        Name    string
        Age     int
        Address *Address
        Tags    []string
    }
    
    expected := Person{
        Name: "Bob",
        Age:  25,
        Address: &Address{
            Street: "123 Main St",
            City:   "New York",
        },
        Tags: []string{"developer", "golang"},
    }
    
    actual := Person{
        Name: "Bob",
        Age:  25,
        Address: &Address{
            Street: "123 Main St",
            City:   "New York",
        },
        Tags: []string{"developer", "golang"},
    }
    
    compare.CompareStructs(t, expected, actual, "Person")
}
```

## 支持的数据类型

- **基本类型**：int, string, bool, float64 等
- **结构体**：嵌套结构体，按字段名匹配
- **指针**：自动处理 nil 检查和解引用
- **切片和数组**：逐元素比较
- **接口**：动态类型比较
- **映射（Map）**：按键值对比较

## 比较策略

1. **按字段名匹配**：优先使用字段名进行匹配，而非严格类型匹配
2. **类型转换**：对于基本类型，尝试进行类型转换后比较
3. **深度递归**：对复杂类型进行递归比较
4. **智能日志**：区分"相同结构不同包"和"完全不同类型"

## 日志输出示例

```
=== RUN   TestAuthorizationList
    compare.go:28: Pagination 相同结构不同包的类型，按字段比较: Pagination
    compare.go:45: Pagination 所有字段比较成功
    compare.go:28: AuthorizationList[0] 相同结构不同包的类型，按字段比较: ListData
    compare.go:45: AuthorizationList[0] 所有字段比较成功
    inout_test.go:148: 授权列表接口比较完成，共比较了 5 个列表项
--- PASS: TestAuthorizationList (1.20s)
```

## 注意事项

1. **导出字段**：只比较导出的字段（首字母大写）
2. **性能考虑**：使用反射，对于大型结构体可能有性能开销
3. **类型安全**：虽然支持跨类型比较，但建议确保字段语义一致
4. **测试专用**：此包专为测试场景设计，不建议在生产代码中使用

## 扩展功能

包中还提供了严格类型匹配的比较函数（以 `compare` 开头的函数），可用于需要严格类型检查的场景。这些函数要求被比较的对象类型完全相同。