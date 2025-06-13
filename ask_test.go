package ask

import (
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		cond   any
		trueV  string
		falseV string
		expect string
	}{
		{true, "yes", "no", "yes"},
		{false, "yes", "no", "no"},
		{"non-empty", "yes", "no", "yes"},
		{"", "yes", "no", "no"},
	}

	for _, test := range tests {
		if res := If(test.cond, test.trueV, test.falseV); res != test.expect {
			t.Errorf("If(%v, %v, %v) = %v; want %v", test.cond, test.trueV, test.falseV, res, test.expect)
		}
	}
}

func TestSimple(t *testing.T) {
	// 1. 基本布尔条件
	score := 85
	t.Log(If(score > 60, "及格", "不及格")) // 输出: 及格

	// 2. 字符串非空判断
	var name string
	t.Log(If(name, "有名字", "无名氏"))    // 输出: 无名氏
	t.Log(If("Alice", "有名字", "无名氏")) // 输出: 有名字

	// 3. 指针非空判断
	var ptr *int
	t.Log(If(ptr, "有指针", "无指针")) // 输出: 无指针
	// 使用指针
	val := 10
	t.Log(If(&val, "有指针", "无指针")) // 输出: 有指针

	// 4. 切片非空判断
	var emptySlice []int
	t.Log(If(emptySlice, "非空切片", "空切片"))  // 输出: 空切片
	t.Log(If([]int{1, 2}, "非空切片", "空切片")) // 输出: 非空切片

	// 5. 自定义结构体
	type User struct{ Name string }
	var u *User
	t.Log(If(u, "有用户", "无用户"))       // 输出: 无用户
	t.Log(If(&User{}, "有用户", "无用户")) // 输出: 有用户
}
