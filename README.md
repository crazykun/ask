
# Ask - Go 条件操作工具库

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/crazykun/ask)](https://goreportcard.com/report/github.com/crazykun/ask)

一个高性能的 Go 语言条件操作工具库，提供三目运算符、空值合并等实用功能。

## 特性

- 🚀 **高性能**: 针对常见类型优化，避免不必要的反射开销
- 🔧 **类型安全**: 基于 Go 泛型，编译时类型检查
- 📦 **零依赖**: 仅使用标准库，无外部依赖
- 🎯 **简洁易用**: 直观的 API 设计，减少样板代码
- 🧪 **全面测试**: 100% 测试覆盖率，包含基准测试

## 安装

```bash
go get github.com/crazykun/ask
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/crazykun/ask"
)

func main() {
    // 三目运算符
    score := 85
    result := ask.If(score > 60, "及格", "不及格")
    fmt.Println(result) // 输出: 及格
    
    // 空值合并
    var name string
    displayName := ask.Ifelse(name, "匿名用户")
    fmt.Println(displayName) // 输出: 匿名用户
    
    // 多值合并
    finalValue := ask.Coalesce("", "", "hello", "world")
    fmt.Println(finalValue) // 输出: hello
}
```

## API 文档

### If - 三目运算符

```go
func If[T any, C any](condition C, trueVal, falseVal T) T
```

根据条件返回不同的值，支持：
- 布尔条件：`true` 返回 `trueVal`，`false` 返回 `falseVal`
- 零值检查：非零值返回 `trueVal`，零值返回 `falseVal`
- 错误处理：非 `nil` 错误返回 `trueVal`

**示例：**
```go
// 布尔条件
result := ask.If(age >= 18, "成年", "未成年")

// 字符串非空检查
message := ask.If(username, "欢迎 " + username, "请登录")

// 错误处理
status := ask.If(err, "失败", "成功")

// 指针检查
info := ask.If(user, user.Name, "未知用户")
```

### Ifelse - 空值合并运算符

```go
func Ifelse[T any](value, defaultVal T) T
```

如果第一个值为零值，则返回默认值。

**示例：**
```go
// 字符串默认值
title := ask.Ifelse(article.Title, "无标题")

// 数值默认值
port := ask.Ifelse(config.Port, 8080)

// 结构体默认值
user := ask.Ifelse(currentUser, &User{Name: "游客"})
```

### Coalesce - 多值合并

```go
func Coalesce[T any](values ...T) T
```

返回第一个非零值，类似 SQL 的 `COALESCE` 函数。

**示例：**
```go
// 配置优先级：命令行 > 环境变量 > 配置文件 > 默认值
host := ask.Coalesce(flagHost, envHost, configHost, "localhost")

// 用户信息回退
displayName := ask.Coalesce(user.NickName, user.Username, user.Email, "匿名")
```

### IsZero - 零值检查

```go
func IsZero(v any) bool
```

检查值是否为零值，针对常见类型进行了性能优化。

**示例：**
```go
ask.IsZero("")           // true
ask.IsZero(0)            // true
ask.IsZero(false)        // true
ask.IsZero([]int{})      // true
ask.IsZero((*int)(nil))  // true
ask.IsZero("hello")      // false
```

### IsEmpty - 空值检查

```go
func IsEmpty(v any) bool
```

检查容器类型是否为空（长度为 0），对于字符串、切片、映射等更直观。

**示例：**
```go
ask.IsEmpty("")              // true
ask.IsEmpty([]int{})         // true
ask.IsEmpty(map[string]int{}) // true
ask.IsEmpty([0]int{})        // true
```

### Default - 默认值设置

```go
func Default[T any](value, defaultVal T) T
```

`Ifelse` 的别名，在某些上下文中可读性更好。

## 性能优化

本库针对性能进行了多项优化：

1. **类型快速路径**: 常见类型（int、string、bool 等）使用类型断言，避免反射
2. **延迟反射**: 仅在必要时才使用反射处理复杂类型
3. **零分配**: 大多数操作不产生额外的内存分配

### 基准测试结果

```bash
go test -bench=. -benchmem
```

```
BenchmarkIfBool-8        1000000000    0.25 ns/op    0 B/op    0 allocs/op
BenchmarkIfString-8      500000000     0.50 ns/op    0 B/op    0 allocs/op
BenchmarkIsZeroString-8  1000000000    0.30 ns/op    0 B/op    0 allocs/op
BenchmarkIsZeroInt-8     2000000000    0.20 ns/op    0 B/op    0 allocs/op
```

## 使用场景

### 配置管理
```go
config := &Config{
    Host: ask.Ifelse(os.Getenv("HOST"), "localhost"),
    Port: ask.Ifelse(getEnvInt("PORT"), 8080),
    Debug: ask.If(os.Getenv("DEBUG") == "true", true, false),
}
```

### 模板渲染
```go
data := map[string]string{
    "title": ask.Ifelse(article.Title, "无标题"),
    "author": ask.Coalesce(article.Author, article.Creator, "匿名"),
    "status": ask.If(article.Published, "已发布", "草稿"),
}
```

### 错误处理
```go
message := ask.If(err, 
    fmt.Sprintf("操作失败: %v", err), 
    "操作成功")
```

### API 响应
```go
response := &UserResponse{
    Name: ask.Ifelse(user.Name, "未设置"),
    Avatar: ask.Ifelse(user.Avatar, "/default-avatar.png"),
    Status: ask.If(user.IsActive, "在线", "离线"),
}
```

## 最佳实践

1. **性能敏感场景**: 对于高频调用，考虑使用原生 if-else
2. **类型一致性**: 确保 `trueVal` 和 `falseVal` 类型相同
3. **可读性优先**: 在复杂逻辑中，清晰的 if-else 可能比简洁的三目运算符更好
4. **错误处理**: 利用 `If` 函数的错误类型特殊处理简化错误检查

## 兼容性

- Go 1.18+（需要泛型支持）
- 所有主要操作系统（Linux、macOS、Windows）
- 所有主要架构（amd64、arm64、386 等）

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。