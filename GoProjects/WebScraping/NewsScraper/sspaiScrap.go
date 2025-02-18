package main

import (
	"fmt"
	"os"
	"strings"
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

	browser := rod.New().
		ControlURL(launchURL).
		Trace(true).
		SlowMotion(1 * time.Second).
		MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://sspai.com")

	// ✅ 等待页面加载
	page.MustWaitLoad()
	fmt.Println("✅ 页面加载完成！")

	// ✅ 获取网页标题
	title, err := page.Eval("() => document.title")
	if err != nil {
		fmt.Println("⚠️ 获取网页标题失败:", err)
	} else {
		fmt.Println("📌 网页标题:", title)
	}

	// ✅ 等待文章列表加载
	page.MustWaitElementsMoreThan(".articleCard", 1)
	fmt.Println("📢 开始爬取文章标题和链接...")

	// ✅ 触发滚动加载
	for i := 0; i < 5; i++ { // 滚动 5 次，加载更多文章
		page.Eval("() => window.scrollTo(0, document.body.scrollHeight)")
		time.Sleep(2 * time.Second) // 等待 2 秒，确保新内容加载
	}
	fmt.Println("📢 滚动加载完成，开始爬取文章...")

	// ✅ 去重逻辑
	articles := page.MustElements(".articleCard")
	uniqueArticles := make(map[string]bool)
	var articleList []string

	for _, article := range articles {
		// **获取标题**
		titleElement := article.MustElement(".title")
		text := strings.TrimSpace(titleElement.MustText())
		text = strings.ToLower(text)               // **转换为小写，防止相同标题大小写不同**
		text = strings.ReplaceAll(text, "\n", " ") // **去掉换行，合并标题**

		// **获取文章链接**
		linkElement, err := article.Element("a")
		var link string
		if err == nil {
			href := linkElement.MustAttribute("href")
			if href != nil { // **检查指针是否为空**
				link = *href
				if !strings.HasPrefix(link, "http") {
					link = "https://sspai.com" + link
				}
			}
		}

		// **去重并存储**
		if text != "" && link != "" && !uniqueArticles[text] {
			uniqueArticles[text] = true
			fullEntry := fmt.Sprintf("%s - %s", text, link)
			fmt.Println("📝 文章:", fullEntry)
			articleList = append(articleList, fullEntry)
		}
	}

	// ✅ 存入文件
	if len(articleList) > 0 {
		err := os.WriteFile("sspai_articles.txt", []byte(strings.Join(articleList, "\n")), 0644)
		if err != nil {
			fmt.Println("⚠️ 文件写入失败:", err)
		} else {
			fmt.Println("✅ 爬取完成，已存入 sspai_articles.txt")
		}
	} else {
		fmt.Println("⚠️ 没有找到文章标题，可能选择器不对！")
	}
}
