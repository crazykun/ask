package ask

import (
	"errors"
	"testing"
)

// Benchmark tests for performance analysis

func BenchmarkIfBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = If(true, "yes", "no")
	}
}

func BenchmarkIfBoolNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result string
		if true {
			result = "yes"
		} else {
			result = "no"
		}
		_ = result
	}
}

func BenchmarkIfString(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = If(str, "yes", "no")
	}
}

func BenchmarkIfStringNative(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		var result string
		if str != "" {
			result = "yes"
		} else {
			result = "no"
		}
		_ = result
	}
}

func BenchmarkIfInt(b *testing.B) {
	num := 42
	for i := 0; i < b.N; i++ {
		_ = If(num, "yes", "no")
	}
}

func BenchmarkIfIntNative(b *testing.B) {
	num := 42
	for i := 0; i < b.N; i++ {
		var result string
		if num != 0 {
			result = "yes"
		} else {
			result = "no"
		}
		_ = result
	}
}

func BenchmarkIfError(b *testing.B) {
	err := errors.New("test error")
	for i := 0; i < b.N; i++ {
		_ = If(err, "error", "ok")
	}
}

func BenchmarkIfErrorNative(b *testing.B) {
	err := errors.New("test error")
	for i := 0; i < b.N; i++ {
		var result string
		if err != nil {
			result = "error"
		} else {
			result = "ok"
		}
		_ = result
	}
}

func BenchmarkIfelse(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = Ifelse(str, "default")
	}
}

func BenchmarkIfelseNative(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		var result string
		if str != "" {
			result = str
		} else {
			result = "default"
		}
		_ = result
	}
}

func BenchmarkCoalesce2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Coalesce("", "default")
	}
}

func BenchmarkCoalesce4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Coalesce("", "", "", "default")
	}
}

func BenchmarkCoalesce8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Coalesce("", "", "", "", "", "", "", "default")
	}
}

func BenchmarkIsZeroString(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = IsZero(str)
	}
}

func BenchmarkIsZeroStringNative(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = str == ""
	}
}

func BenchmarkIsZeroInt(b *testing.B) {
	num := 42
	for i := 0; i < b.N; i++ {
		_ = IsZero(num)
	}
}

func BenchmarkIsZeroIntNative(b *testing.B) {
	num := 42
	for i := 0; i < b.N; i++ {
		_ = num == 0
	}
}

func BenchmarkIsZeroSlice(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_ = IsZero(slice)
	}
}

func BenchmarkIsZeroSliceNative(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_ = slice == nil
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = IsEmpty(str)
	}
}

func BenchmarkIsEmptyNative(b *testing.B) {
	str := "hello"
	for i := 0; i < b.N; i++ {
		_ = len(str) == 0
	}
}

func BenchmarkIsZeroStringSlice(b *testing.B) {
	slice := []string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		_ = IsZero(slice)
	}
}

func BenchmarkIsZeroStringSliceNative(b *testing.B) {
	slice := []string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		_ = len(slice) == 0
	}
}

func BenchmarkIsEmptyStringSlice(b *testing.B) {
	slice := []string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		_ = IsEmpty(slice)
	}
}

func BenchmarkIsEmptyStringSliceNative(b *testing.B) {
	slice := []string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		_ = len(slice) == 0
	}
}

// Comparative benchmarks for different approaches
func BenchmarkComplexCondition_Ask(b *testing.B) {
	user := &struct {
		Name   string
		Email  string
		Active bool
	}{
		Name:   "",
		Email:  "user@example.com",
		Active: true,
	}

	for i := 0; i < b.N; i++ {
		displayName := Coalesce(user.Name, extractEmailName(user.Email), "Anonymous")
		status := If(user.Active, "Online", "Offline")
		_ = displayName + " - " + status
	}
}

func BenchmarkComplexCondition_Native(b *testing.B) {
	user := &struct {
		Name   string
		Email  string
		Active bool
	}{
		Name:   "",
		Email:  "user@example.com",
		Active: true,
	}

	for i := 0; i < b.N; i++ {
		var displayName string
		if user.Name != "" {
			displayName = user.Name
		} else if emailName := extractEmailName(user.Email); emailName != "" {
			displayName = emailName
		} else {
			displayName = "Anonymous"
		}

		var status string
		if user.Active {
			status = "Online"
		} else {
			status = "Offline"
		}

		_ = displayName + " - " + status
	}
}

func extractEmailName(email string) string {
	for i, c := range email {
		if c == '@' {
			return email[:i]
		}
	}
	return email
}