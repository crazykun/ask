package ask

import "reflect"

// If 三目运算符 (支持布尔条件或值非空判断)
//   - 若 condition 为 bool 类型：true 返回 trueVal，false 返回 falseVal
//   - 若 condition 非 bool 类型：非零值返回 trueVal，零值返回 falseVal
func If[T any, C any](condition C, trueVal, falseVal T) T {
	// 如果条件本身就是布尔类型
	if b, ok := any(condition).(bool); ok {
		if b {
			return trueVal
		}
		return falseVal
	}

	// 判断是不是error类型 - 修复逻辑：error非nil时应该返回trueVal
	if err, ok := any(condition).(error); ok {
		if err != nil {
			return trueVal
		}
		return falseVal
	}

	if !IsZero(condition) {
		return trueVal
	}
	return falseVal
}

// Ifelse 空值合并运算符  问号冒号(?:)
// 如果第一个值为零值或nil，则返回第二个值
func Ifelse[T any](condition, defaultVal T) T {
	// 通过反射检查第一个值是否为零值
	if !IsZero(condition) {
		return condition
	}
	return defaultVal
}

// IsZero 检查值是否为零值
// 支持的类型包括：布尔、整数、浮点数、字符串、指针、切片、映射、通道、函数和接口
func IsZero(v any) bool {
	if v == nil {
		return true
	}

	// 特殊处理常见基础类型，避免反射开销
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
	}

	// 反射处理复杂类型
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map,
		reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	case reflect.Array:
		// 数组需要检查所有元素
		return rv.IsZero()
	case reflect.Struct:
		// 结构体需要检查所有字段
		return rv.IsZero()
	default:
		return rv.IsZero()
	}
}
