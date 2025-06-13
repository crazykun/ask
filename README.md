
# ask

Go 语言中的三目运算符实现。

由于 Go 不支持原生的三目运算符（`condition ? trueVal : falseVal`），本库提供了一个通用的三目运算符函数 [If](file:///home/data/go/ask/ask.go#L5-L18) 和一个空值合并操作符 [Ifelse](file:///home/data/go/ask/ask.go#L22-L28)，可用于简化条件判断逻辑。

## 安装

```bash
go get github.com/crazykun/ask
```

## 使用示例

### 三目运算符：`If(condition, trueVal, falseVal)`

```go
result := ask.If(true, "yes", "no") // 返回 "yes"
result := ask.If(false, "yes", "no") // 返回 "no"
result := ask.If("non-empty", "yes", "no") // 返回 "yes"
result := ask.If("", "yes", "no") // 返回 "no"
```

### 空值合并运算符：`Ifelse(value, defaultVal)`

```go
value := ask.Ifelse("", "default") // 返回 "default"
value := ask.Ifelse("custom", "default") // 返回 "custom"
value := ask.Ifelse(0, 10) // 返回 10
value := ask.Ifelse(5, 10) // 返回 5
```

## 支持的类型

- 布尔值：`true` / `false`
- 数值类型：`int`, `float`, `uint` 等
- 字符串：`string`
- 指针、切片、映射等引用类型（自动判断是否为 `nil`）

## 函数说明

### `If[T any, C any](condition C, trueVal, falseVal T) T`

根据条件返回不同的值。支持布尔表达式或非空判断。

### `Ifelse[T any](condition, defaultVal T) T`

如果第一个值为空（零值），则返回默认值。

### `IsZero[T any](value T) bool`

通过反射判断一个值是否为零值（空字符串、0、nil 等）。

## 注意事项

- 性能敏感场景建议使用直接比较代替泛型函数。
- 所有函数均基于反射实现，不适用于非常高频调用的场景。
- 可扩展性强，适用于大多数类型的条件赋值场景。

## 贡献

欢迎提交 PR 或 Issue 来改进此项目！

## 许可证

MIT License