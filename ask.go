// Package ask provides utility functions for conditional operations and zero value checking.
// It offers ternary-like operations and null coalescing functionality with Go generics.
package ask

import "reflect"

// If implements a ternary operator that supports both boolean conditions and zero value checking.
//   - If condition is bool type: returns trueVal if true, falseVal if false
//   - If condition is non-bool type: returns trueVal if non-zero, falseVal if zero
//   - Special handling for error type: returns trueVal if error is non-nil
//
// If 三目运算符 (支持布尔条件或值非空判断)
//   - 若 condition 为 bool 类型：true 返回 trueVal，false 返回 falseVal
//   - 若 condition 非 bool 类型：非零值返回 trueVal，零值返回 falseVal
func If[T any, C any](condition C, trueVal, falseVal T) T {
	// Fast path for boolean conditions
	if b, ok := any(condition).(bool); ok {
		if b {
			return trueVal
		}
		return falseVal
	}

	// Special handling for error type
	if err, ok := any(condition).(error); ok {
		if err != nil {
			return falseVal
		}
		return trueVal
	}

	// General zero value check
	if !IsZero(condition) {
		return trueVal
	}
	return falseVal
}

// Ifelse implements null coalescing operator (?:).
// Returns the first value if it's non-zero, otherwise returns the default value.
// Ifelse 空值合并运算符  问号冒号(?:)
// 如果第一个值为零值或nil，则返回第二个值
func Ifelse[T any](value, defaultVal T) T {
	if !IsZero(value) {
		return value
	}
	return defaultVal
}

// IsEmpty checks if a value is considered "empty" (more intuitive than zero for some types).
// For strings, slices, maps, and channels, it checks if length is 0.
// For other types, it falls back to IsZero.
// IsZero 检查值是否为零值
// 支持的类型包括：布尔、整数、浮点数、字符串、指针、切片、映射、通道、函数和接口
func IsEmpty(v any) bool {
	if v == nil {
		return true
	}

	switch x := v.(type) {
	case string:
		return len(x) == 0
	case []any:
		return len(x) == 0
	case []bool:
		return len(x) == 0
	case []int:
		return len(x) == 0
	case []int8:
		return len(x) == 0
	case []int16:
		return len(x) == 0
	case []int32:
		return len(x) == 0
	case []int64:
		return len(x) == 0
	case []uint:
		return len(x) == 0
	case []uint8:
		return len(x) == 0
	case []uint16:
		return len(x) == 0
	case []uint32:
		return len(x) == 0
	case []uint64:
		return len(x) == 0
	case []uintptr:
		return len(x) == 0
	case []float32:
		return len(x) == 0
	case []float64:
		return len(x) == 0
	case []complex64:
		return len(x) == 0
	case []complex128:
		return len(x) == 0
	case []string:
		return len(x) == 0
	case map[string]any:
		return len(x) == 0
	case map[string]string:
		return len(x) == 0
	case map[string]int:
		return len(x) == 0
	case map[int]string:
		return len(x) == 0
	case map[int]int:
		return len(x) == 0
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return rv.Len() == 0
	case reflect.Array:
		return rv.Len() == 0
	default:
		return IsZero(v)
	}
}

// IsZero checks if a value is the zero value for its type.
// Optimized with type switches to avoid reflection for common types.
// IsZero 检查值是否为零值 对于常见类型使用类型断言来优化性能，避免反射开销
// 支持的类型包括：布尔、整数、浮点数、复数、字符串、错误、指针、切片、映射、数组和结构体
// 对于复杂类型使用反射来判断是否为零值
// 反射处理复杂类型，避免了对每个类型都写一个判断函数
func IsZero(v any) bool {
	if v == nil {
		return true
	}

	// Fast path for common types to avoid reflection overhead
	// 常见类型的快速路径，避免反射开销
	switch x := v.(type) {
	case bool:
		return !x
	case int:
		return x == 0
	case int8:
		return x == 0
	case int16:
		return x == 0
	case int32:
		return x == 0
	case int64:
		return x == 0
	case uint:
		return x == 0
	case uint8:
		return x == 0
	case uint16:
		return x == 0
	case uint32:
		return x == 0
	case uint64:
		return x == 0
	case uintptr:
		return x == 0
	case float32:
		return x == 0
	case float64:
		return x == 0
	case complex64:
		return x == 0
	case complex128:
		return x == 0
	case string:
		return x == ""
	case error:
		return x == nil
	// 添加切片和映射的快速路径
	case []bool:
		return len(x) == 0
	case []int:
		return len(x) == 0
	case []int8:
		return len(x) == 0
	case []int16:
		return len(x) == 0
	case []int32:
		return len(x) == 0
	case []int64:
		return len(x) == 0
	case []uint:
		return len(x) == 0
	case []uint8:
		return len(x) == 0
	case []uint16:
		return len(x) == 0
	case []uint32:
		return len(x) == 0
	case []uint64:
		return len(x) == 0
	case []uintptr:
		return len(x) == 0
	case []float32:
		return len(x) == 0
	case []float64:
		return len(x) == 0
	case []complex64:
		return len(x) == 0
	case []complex128:
		return len(x) == 0
	case []string:
		return len(x) == 0
	case []any:
		return len(x) == 0
	case map[string]any:
		return len(x) == 0
	case map[string]string:
		return len(x) == 0
	case map[string]int:
		return len(x) == 0
	case map[int]string:
		return len(x) == 0
	case map[int]int:
		return len(x) == 0
	}

	// Reflection fallback for complex types
	// 反射处理复杂类型
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	case reflect.Slice, reflect.Map:
		// Slice and map: nil or empty slice/map is considered zero
		// 切片和映射：nil 或长度为 0 都认为是零值
		return rv.IsNil() || rv.Len() == 0
	case reflect.Array:
		// Array: nil is not applicable, but an array with all zero values is considered zero
		// 数组需要检查所有元素
		return rv.IsZero()
	case reflect.Struct:
		// Struct: nil is not applicable, but a struct with all zero values is considered zero
		// 结构体需要检查所有字段
		return rv.IsZero()
	default:
		return rv.IsZero()
	}
}

// Coalesce returns the first non-zero value from the provided arguments.
// Similar to SQL COALESCE function.
func Coalesce[T any](values ...T) T {
	for _, v := range values {
		if !IsZero(v) {
			return v
		}
	}
	var zero T
	return zero
}

// Default returns defaultVal if value is zero, otherwise returns value.
// This is an alias for Ifelse for better readability in some contexts.
func Default[T any](value, defaultVal T) T {
	return Ifelse(value, defaultVal)
}
