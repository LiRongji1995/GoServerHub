package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func main() {
	c := colly.NewCollector()

	// 设置 User-Agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"

	// 设置超时
	c.SetRequestTimeout(60 * time.Second)

	// 访问测试
	err := c.Visit("https://example.com")
	if err != nil {
		fmt.Println("访问失败:", err)
	} else {
		fmt.Println("访问成功!")
	}
}
