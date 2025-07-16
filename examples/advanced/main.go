package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/crazykun/ask"
)

func main() {
	fmt.Println("=== Ask 库高级使用示例 ===\n")

	// 1. 配置管理示例
	configExample()

	// 2. API 响应构建示例
	apiResponseExample()

	// 3. 模板数据准备示例
	templateDataExample()

	// 4. 链式操作示例
	chainedOperationsExample()
}

// 配置管理示例
func configExample() {
	fmt.Println("1. 配置管理示例:")

	// 模拟环境变量
	os.Setenv("APP_PORT", "3000")
	os.Setenv("APP_DEBUG", "true")

	config := &AppConfig{
		Host:    ask.Coalesce(os.Getenv("APP_HOST"), "localhost"),
		Port:    ask.Ifelse(getEnvInt("APP_PORT"), 8080),
		Debug:   ask.If(os.Getenv("APP_DEBUG") == "true", true, false),
		LogFile: ask.Ifelse(os.Getenv("LOG_FILE"), "/var/log/app.log"),
		Workers: ask.Ifelse(getEnvInt("WORKERS"), 4),
	}

	fmt.Printf("   配置: %+v\n\n", config)
}

// API 响应构建示例
func apiResponseExample() {
	fmt.Println("2. API 响应构建示例:")

	// 模拟用户数据
	users := []*User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Avatar: "", IsActive: true},
		{ID: 2, Name: "", Email: "bob@example.com", Avatar: "/avatars/bob.jpg", IsActive: false},
		{ID: 3, Name: "Charlie", Email: "", Avatar: "", IsActive: true},
	}

	var responses []UserResponse
	for _, user := range users {
		response := UserResponse{
			ID:     user.ID,
			Name:   ask.Coalesce(user.Name, extractNameFromEmail(user.Email), "匿名用户"),
			Email:  ask.Ifelse(user.Email, "未设置"),
			Avatar: ask.Ifelse(user.Avatar, "/default-avatar.png"),
			Status: ask.If(user.IsActive, "在线", "离线"),
		}
		responses = append(responses, response)
	}

	for _, resp := range responses {
		fmt.Printf("   用户: %+v\n", resp)
	}
	fmt.Println()
}

// 模板数据准备示例
func templateDataExample() {
	fmt.Println("3. 模板数据准备示例:")

	articles := []*Article{
		{Title: "Go 语言入门", Author: "Alice", Content: "这是一篇关于 Go 的文章...", Published: true, Views: 1500},
		{Title: "", Author: "", Content: "未完成的文章", Published: false, Views: 0},
		{Title: "数据结构与算法", Author: "Bob", Content: "", Published: true, Views: 2300},
	}

	for i, article := range articles {
		data := map[string]interface{}{
			"title":       ask.Ifelse(article.Title, "无标题"),
			"author":      ask.Coalesce(article.Author, "匿名作者"),
			"content":     ask.Ifelse(article.Content, "暂无内容"),
			"status":      ask.If(article.Published, "已发布", "草稿"),
			"views":       ask.If(article.Views > 0, fmt.Sprintf("%d 次浏览", article.Views), "暂无浏览"),
			"popularity":  ask.If(article.Views > 1000, "热门", "普通"),
		}

		fmt.Printf("   文章 %d: %v\n", i+1, data)
	}
	fmt.Println()
}

// 链式操作示例
func chainedOperationsExample() {
	fmt.Println("4. 链式操作示例:")

	// 复杂的条件逻辑
	processOrder := func(order *Order) string {
		// 状态检查
		statusMsg := ask.If(order == nil, "订单不存在", 
			ask.If(order.IsPaid, 
				ask.If(order.IsShipped, "已发货", "待发货"),
				ask.If(order.IsCancelled, "已取消", "待支付")))

		// 优先级计算
		priority := ask.If(order != nil && order.IsVIP, "高优先级",
			ask.If(order != nil && order.Amount > 1000, "中优先级", "普通优先级"))

		// 处理建议
		suggestion := ask.If(order == nil, "请检查订单号",
			ask.If(order.IsCancelled, "联系客服处理",
				ask.If(!order.IsPaid, "请尽快支付",
					ask.If(!order.IsShipped, "正在准备发货", "请耐心等待收货"))))

		return fmt.Sprintf("状态: %s | 优先级: %s | 建议: %s", statusMsg, priority, suggestion)
	}

	// 测试不同订单状态
	orders := []*Order{
		nil,
		{ID: 1, Amount: 500, IsPaid: false, IsShipped: false, IsCancelled: false, IsVIP: false},
		{ID: 2, Amount: 1500, IsPaid: true, IsShipped: false, IsCancelled: false, IsVIP: true},
		{ID: 3, Amount: 800, IsPaid: true, IsShipped: true, IsCancelled: false, IsVIP: false},
		{ID: 4, Amount: 300, IsPaid: false, IsShipped: false, IsCancelled: true, IsVIP: false},
	}

	for i, order := range orders {
		result := processOrder(order)
		fmt.Printf("   订单 %d: %s\n", i+1, result)
	}
}

// 辅助函数
func getEnvInt(key string) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return 0
}

func extractNameFromEmail(email string) string {
	if email == "" {
		return ""
	}
	for i, c := range email {
		if c == '@' {
			return email[:i]
		}
	}
	return email
}

// 数据结构
type AppConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
	LogFile string `json:"log_file"`
	Workers int    `json:"workers"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsActive bool   `json:"is_active"`
}

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Status string `json:"status"`
}

type Article struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	Views     int    `json:"views"`
}

type Order struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	IsPaid      bool    `json:"is_paid"`
	IsShipped   bool    `json:"is_shipped"`
	IsCancelled bool    `json:"is_cancelled"`
	IsVIP       bool    `json:"is_vip"`
}