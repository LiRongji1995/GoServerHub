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

	// 等待页面完全加载
	page.MustWaitLoad()
	fmt.Println("✅ 页面加载完成！")

	// 存储文章标题和链接
	articleSet := make(map[string]string)

	// **自动翻页爬取**
	for {
		// 获取文章
		articles := page.MustElements("div.articleCard")
		fmt.Printf("📢 发现 %d 篇文章\n", len(articles))

		for _, article := range articles {
			// **安全获取标题**
			titleElement, err := article.Element(".title")
			if err != nil || titleElement == nil {
				fmt.Println("⚠️ 找不到文章标题，跳过")
				continue
			}
			title := titleElement.MustText()

			// **安全获取链接**
			linkElement, err := article.Element("a")
			if err != nil || linkElement == nil {
				fmt.Println("⚠️ 找不到文章链接，跳过")
				continue
			}
			linkPtr, err := linkElement.Attribute("href")
			if err != nil || linkPtr == nil {
				fmt.Println("⚠️ 获取链接失败，跳过")
				continue
			}
			link := "https://sspai.com" + *linkPtr

			// **去重存储**
			articleSet[title] = link
		}

		// **查找 “加载更多” 按钮**
		nextPageButton, err := page.Element(".btn-more")
		if err != nil || nextPageButton == nil {
			fmt.Println("✅ 没有更多文章，爬取结束！")
			break
		}

		// **点击加载更多**
		fmt.Println("🔄 点击加载更多...")
		nextPageButton.MustClick()
		page.MustWaitLoad()
		time.Sleep(2 * time.Second) // 等待页面加载
	}

	// **保存到文件**
	var articleList []string
	for title, link := range articleSet {
		articleList = append(articleList, title+" - "+link)
	}

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
