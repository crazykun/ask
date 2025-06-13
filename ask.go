package ask

import "reflect"

type ValidType interface{}

// If 三目运算符 (支持布尔条件或值非空判断)
func If[T any, C any](condition C, trueVal, falseVal T) T {
	// 如果条件本身就是布尔类型
	if b, ok := any(condition).(bool); ok {
		if b {
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
func IsZero[T ValidType](value T) bool {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return true // 无效值视为零值
	}

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool() // 布尔类型，false 视为零值
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	default:
		return v.IsZero()
	}
}
