package main

import (
	"fmt"
	"os"

	"github.com/crazykun/ask"
)

func main() {
	fmt.Println("=== Ask 库基本使用示例 ===\n")

	// 1. 三目运算符基本用法
	fmt.Println("1. 三目运算符 If:")
	score := 85
	result := ask.If(score > 60, "及格", "不及格")
	fmt.Printf("   分数 %d: %s\n", score, result)

	age := 16
	status := ask.If(age >= 18, "成年人", "未成年人")
	fmt.Printf("   年龄 %d: %s\n", age, status)

	// 2. 字符串非空判断
	fmt.Println("\n2. 字符串非空判断:")
	var username string
	greeting := ask.If(username, "欢迎 "+username, "请先登录")
	fmt.Printf("   空用户名: %s\n", greeting)

	username = "Alice"
	greeting = ask.If(username, "欢迎 "+username, "请先登录")
	fmt.Printf("   用户名 '%s': %s\n", username, greeting)

	// 3. 空值合并运算符
	fmt.Println("\n3. 空值合并 Ifelse:")
	var title string
	displayTitle := ask.Ifelse(title, "无标题")
	fmt.Printf("   空标题: %s\n", displayTitle)

	title = "Go 编程指南"
	displayTitle = ask.Ifelse(title, "无标题")
	fmt.Printf("   标题 '%s': %s\n", title, displayTitle)

	// 4. 数值默认值
	fmt.Println("\n4. 数值默认值:")
	var port int
	serverPort := ask.Ifelse(port, 8080)
	fmt.Printf("   默认端口: %d\n", serverPort)

	port = 3000
	serverPort = ask.Ifelse(port, 8080)
	fmt.Printf("   自定义端口: %d\n", serverPort)

	// 5. 多值合并 Coalesce
	fmt.Println("\n5. 多值合并 Coalesce:")
	var cmdHost, envHost, configHost string
	envHost = os.Getenv("HOST") // 通常为空
	finalHost := ask.Coalesce(cmdHost, envHost, configHost, "localhost")
	fmt.Printf("   最终主机: %s\n", finalHost)

	// 模拟有值的情况
	configHost = "config.example.com"
	finalHost = ask.Coalesce(cmdHost, envHost, configHost, "localhost")
	fmt.Printf("   配置主机: %s\n", finalHost)

	// 6. 错误处理
	fmt.Println("\n6. 错误处理:")
	var err error
	message := ask.If(err, "操作失败", "操作成功")
	fmt.Printf("   无错误: %s\n", message)

	err = fmt.Errorf("网络连接失败")
	message = ask.If(err, "操作失败: "+err.Error(), "操作成功")
	fmt.Printf("   有错误: %s\n", message)

	// 7. 指针检查
	fmt.Println("\n7. 指针检查:")
	var user *User
	info := ask.If(user, "用户已登录", "用户未登录")
	fmt.Printf("   空指针: %s\n", info)

	user = &User{Name: "Bob", Email: "bob@example.com"}
	info = ask.If(user, "用户: "+user.Name, "用户未登录")
	fmt.Printf("   有效用户: %s\n", info)

	// 8. 切片检查
	fmt.Println("\n8. 切片检查:")
	var items []string
	listStatus := ask.If(items, "有数据", "无数据")
	fmt.Printf("   空切片: %s\n", listStatus)

	items = []string{"item1", "item2"}
	listStatus = ask.If(items, fmt.Sprintf("有 %d 项数据", len(items)), "无数据")
	fmt.Printf("   非空切片: %s\n", listStatus)
}

type User struct {
	Name  string
	Email string
}