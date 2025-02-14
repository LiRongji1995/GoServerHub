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

	// **使用 Edge 作为浏览器**
	chromePath := "C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe" // Edge 真实路径

	// **启动 Edge**
	launchURL := launcher.New().
		Leakless(false).
		Headless(false). // 显示浏览器窗口，方便调试
		Bin(chromePath). // 使用 Edge
		MustLaunch()

	fmt.Println("连接 Edge 浏览器...")
	browser := rod.New().
		ControlURL(launchURL).
		Trace(true).                 // 显示详细日志
		SlowMotion(1 * time.Second). // 防止执行太快，导致页面未加载
		MustConnect()
	defer browser.MustClose()

	// **打开知乎首页**
	page := browser.MustPage("https://www.zhihu.com")

	// **伪装 User-Agent 以避免被检测**
	_ = proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}.Call(page)

	// **设置 Cookie**
	page.SetCookies([]*proto.NetworkCookieParam{
		{Name: "z_c0", Value: "2|1:0|10:1739086183|4:z_c0|80:MS4xcFRHMFRRQUFBQUFtQUFBQVlBSlZUV2VubFdoVE5SZWZYRjJWdWpvRmFaRFREN2oySGtScllnPT0=|65eff346767845f4ddc9229ccd0142f221af844e04de4b1ec4111d3a0a301218", Domain: ".zhihu.com", Path: "/"},
		{Name: "_xsrf", Value: "IEOMNWrIbl5C4V1ePSpOXjVSIx5ToxS9", Domain: ".zhihu.com", Path: "/"},
		{Name: "d_c0", Value: "AAAaYB8ffBiPTsRTGHKyuGVD0wzwoTyvMtE=|1713426518", Domain: ".zhihu.com", Path: "/"},
		{Name: "_zap", Value: "41d4d829-c480-4f97-9aea-d5f8432f027a", Domain: ".zhihu.com", Path: "/"},
		{Name: "q_c1", Value: "73ec97f7ce994be0805b19e244fdaec4|1735270397000|1735270397000", Domain: ".zhihu.com", Path: "/"},
	})

	// **刷新页面让 Cookie 生效**
	page.MustReload()
	fmt.Println("页面已刷新，等待加载...")

	// **等待页面加载完成**
	page.MustWaitLoad()
	fmt.Println("页面加载完成")

	// **打印 HTML 结构，确认是否登录成功**
	html := page.MustHTML()
	fmt.Println("页面 HTML 预览:", html[:500]) // 只打印前 500 个字符，避免控制台太长

	// **等待 `Topstory-body` 出现**
	page.MustWaitElementsMoreThan("body.Topstory-body", 0)

	// **提取 `Topstory-body`**
	topstory := page.MustElement("body.Topstory-body").MustHTML()
	fmt.Println("Topstory-body HTML:", topstory)
}
