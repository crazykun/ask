package ask

import (
	"errors"
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name   string
		cond   any
		trueV  string
		falseV string
		expect string
	}{
		{"bool true", true, "yes", "no", "yes"},
		{"bool false", false, "yes", "no", "no"},
		{"non-empty string", "non-empty", "yes", "no", "yes"},
		{"empty string", "", "yes", "no", "no"},
		{"non-zero int", 42, "yes", "no", "yes"},
		{"zero int", 0, "yes", "no", "no"},
		{"nil pointer", (*int)(nil), "yes", "no", "no"},
		{"non-nil error", errors.New("error"), "yes", "no", "yes"},
		{"nil error", error(nil), "yes", "no", "no"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.cond, tt.trueV, tt.falseV); got != tt.expect {
				t.Errorf("If(%v, %v, %v) = %v; want %v", tt.cond, tt.trueV, tt.falseV, got, tt.expect)
			}
		})
	}
}

func TestIfelse(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		defaultV any
		expect   any
	}{
		{"non-empty string", "hello", "default", "hello"},
		{"empty string", "", "default", "default"},
		{"non-zero int", 42, 100, 42},
		{"zero int", 0, 100, 100},
		{"nil pointer", (*string)(nil), "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ifelse(tt.value, tt.defaultV); got != tt.expect {
				t.Errorf("Ifelse(%v, %v) = %v; want %v", tt.value, tt.defaultV, got, tt.expect)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		expect bool
	}{
		{"nil", nil, true},
		{"bool false", false, true},
		{"bool true", true, false},
		{"int zero", 0, true},
		{"int non-zero", 42, false},
		{"string empty", "", true},
		{"string non-empty", "hello", false},
		{"slice nil", []int(nil), true},
		{"slice empty", []int{}, true},
		{"slice non-empty", []int{1, 2}, false},
		{"map nil", map[string]int(nil), true},
		{"map empty", map[string]int{}, true},
		{"map non-empty", map[string]int{"a": 1}, false},
		{"pointer nil", (*int)(nil), true},
		{"error nil", error(nil), true},
		{"error non-nil", errors.New("error"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.value); got != tt.expect {
				t.Errorf("IsZero(%v) = %v; want %v", tt.value, got, tt.expect)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		value  any
		expect bool
	}{
		{"nil", nil, true},
		{"string empty", "", true},
		{"string non-empty", "hello", false},
		{"slice empty", []int{}, true},
		{"slice non-empty", []int{1}, false},
		{"map empty", map[string]int{}, true},
		{"map non-empty", map[string]int{"a": 1}, false},
		{"array empty", [0]int{}, true},
		{"array non-empty", [1]int{42}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.value); got != tt.expect {
				t.Errorf("IsEmpty(%v) = %v; want %v", tt.value, got, tt.expect)
			}
		})
	}
}

func TestCoalesce(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		expect string
	}{
		{"first non-zero", []string{"", "", "hello", "world"}, "hello"},
		{"all zero", []string{"", "", ""}, ""},
		{"first non-zero", []string{"first", "second"}, "first"},
		{"empty slice", []string{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Coalesce(tt.values...); got != tt.expect {
				t.Errorf("Coalesce(%v) = %v; want %v", tt.values, got, tt.expect)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		defaultV string
		expect   string
	}{
		{"non-empty value", "hello", "default", "hello"},
		{"empty value", "", "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(tt.value, tt.defaultV); got != tt.expect {
				t.Errorf("Default(%v, %v) = %v; want %v", tt.value, tt.defaultV, got, tt.expect)
			}
		})
	}
}



// Example tests for documentation
func ExampleIf() {
	score := 85
	result := If(score > 60, "及格", "不及格")
	println(result) // 输出: 及格
}

func ExampleIfelse() {
	var name string
	result := Ifelse(name, "无名氏")
	println(result) // 输出: 无名氏
}

func ExampleCoalesce() {
	result := Coalesce("", "", "hello", "world")
	println(result) // 输出: hello
}
