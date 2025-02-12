package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func main() {
	fmt.Println("程序开始执行...")

	// 指定 Chrome 路径
	chromePath := "C:/Users/John Green/AppData/Roaming/rod/browser/chrome-win/chrome.exe"

	// 启动 Chrome，不使用 headless 模式，降低封禁概率
	fmt.Println("启动 Chrome 浏览器...")
	launchURL := launcher.New().
		Leakless(false).
		Headless(false). // 让浏览器可见
		Bin(chromePath).
		MustLaunch()

	fmt.Println("连接 Chrome 浏览器...")
	browser := rod.New().ControlURL(launchURL).Trace(true).SlowMotion(1 * time.Second).MustConnect()
	defer browser.MustClose()

	// 打开知乎问题页面
	fmt.Println("打开知乎问题页面...")
	page := browser.MustPage("https://www.zhihu.com/question/314503115")

	// 设置 User-Agent
	_ = proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}.Call(page)

	// 设置 Cookies
	page.SetCookies([]*proto.NetworkCookieParam{
		{
			Name:   "z_c0",
			Value:  "2|1:0|10:1739086183|4:z_c0|80:MS4xcFRHMFRRQUFBQUFtQUFBQVlBSlZUV2VubFdoVE5SZWZYRjJWdWpvRmFaRFREN2oySGtScllnPT0=|65eff346767845f4ddc9229ccd0142f221af844e04de4b1ec4111d3a0a301218",
			Domain: ".zhihu.com",
			Path:   "/",
		},
		{
			Name:   "_xsrf",
			Value:  "IEOMNWrIbl5C4V1ePSpOXjVSIx5ToxS9",
			Domain: ".zhihu.com",
			Path:   "/",
		},
		{
			Name:   "_zap",
			Value:  "41d4d829-c480-4f97-9aea-d5f8432f027a",
			Domain: ".zhihu.com",
			Path:   "/",
		},
		{
			Name:   "d_c0",
			Value:  "AAAaYB8ffBiPTsRTGHKyuGVD0wzwoTyvMtE=|1713426518",
			Domain: ".zhihu.com",
			Path:   "/",
		},
	})

	page.MustWaitLoad()
	time.Sleep(3 * time.Second) // 额外等待 JS 渲染
	fmt.Println("页面加载完成")

	// 获取问题标题
	fmt.Println("尝试获取问题标题...")
	page.MustElement("h1.QuestionHeader-title").MustWaitVisible()
	title := page.MustElement("h1.QuestionHeader-title").MustText()
	fmt.Println("知乎问题标题:", title)

	// 获取前 5 个回答
	fmt.Println("尝试获取前 5 个回答...")
	answers := page.MustElements(".RichContent-inner")
	for i, ans := range answers[:5] {
		fmt.Printf("回答 %d: %s\n", i+1, ans.MustText())
	}

	fmt.Println("程序执行完毕")
}
