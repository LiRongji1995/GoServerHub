package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	fmt.Println("🚀 启动 Edge 浏览器...")
	chromePath := "C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe"

	launchURL := launcher.New().
		Leakless(false).
		Headless(false).
		Bin(chromePath).
		MustLaunch()

	browser := rod.New().ControlURL(launchURL).Trace(true).SlowMotion(1 * time.Second).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://sspai.com") // 你可以换成别的网页

	page.MustWaitLoad()
	fmt.Println("✅ 页面加载完成！")

	// 获取 HTML 源码
	html, err := page.HTML()
	if err != nil {
		fmt.Println("⚠️ 获取网页 HTML 失败:", err)
		return
	}

	// 保存到本地文件
	err = os.WriteFile("83522.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("⚠️ 保存 HTML 失败:", err)
	} else {
		fmt.Println("✅ HTML 已保存到 83522.html")
	}
}
