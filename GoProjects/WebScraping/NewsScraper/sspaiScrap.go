package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/spf13/viper"
)

type Article struct {
	Title string
	Link  string
}

func main() {
	// 初始化配置
	initConfig()

	browser, cleanup := initBrowser()
	defer cleanup()

	page := navigatePage(browser, viper.GetString("targetURL"))
	performScrolling(page)
	articles := extractArticles(page)
	saveArticles(articles)
}

// 初始化配置
func initConfig() {
	viper.SetConfigName("config") // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 查找配置文件的路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("❌ 无法读取配置文件: %v", err)
	}
}

// 初始化浏览器实例
func initBrowser() (*rod.Browser, func()) {
	log.Println("🚀 启动 Edge 浏览器...")
	launcher := launcher.New().
		Leakless(false).
		Headless(false).
		Bin(viper.GetString("edgePath"))

	controlURL, err := launcher.Launch()
	if err != nil {
		log.Fatal("❌ 浏览器启动失败:", err)
	}

	browser := rod.New().
		ControlURL(controlURL).
		Trace(true).
		SlowMotion(1 * time.Second)

	if err := browser.Connect(); err != nil {
		log.Fatal("❌ 浏览器连接失败:", err)
	}

	return browser, func() {
		if err := browser.Close(); err != nil {
			log.Println("⚠️ 关闭浏览器时出错:", err)
		}
	}
}

// 导航到指定页面
func navigatePage(browser *rod.Browser, url string) *rod.Page {
	page, err := browser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		log.Fatal("❌ 创建页面失败:", err)
	}

	if err := page.WaitLoad(); err != nil {
		log.Fatal("❌ 页面加载失败:", err)
	}
	log.Println("✅ 页面加载完成！")

	// 使用 Eval 方法获取页面标题
	title, err := page.Eval("() => document.title")
	if err != nil {
		log.Println("⚠️ 获取页面标题失败:", err)
	} else {
		log.Println("📌 网页标题:", title.Value.String())
	}
	return page
}

// 执行滚动加载
func performScrolling(page *rod.Page) {
	log.Println("📢 开始滚动加载更多内容...")

	scrollScript := "() => window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' })"

	for i := 0; i < viper.GetInt("maxScrolls"); i++ {
		if _, err := page.Eval(scrollScript); err != nil {
			log.Println("⚠️ 滚动操作失败:", err)
			break
		}

		// 使用更可靠的等待方式
		if err := page.WaitIdle(time.Minute); err != nil {
			log.Println("⚠️ 等待页面空闲失败:", err)
			break
		}

		time.Sleep(viper.GetDuration("scrollInterval"))
	}
	log.Println("✅ 滚动加载完成")
}

// 提取文章信息
func extractArticles(page *rod.Page) []Article {
	elements, err := page.Elements(viper.GetString("articleSelector"))
	if err != nil {
		log.Fatal("❌ 获取文章元素失败:", err)
	}

	unique := make(map[string]struct{})
	var articles []Article

	for _, el := range elements {
		article, err := parseArticle(el)
		if err != nil {
			log.Println("⚠️ 解析文章失败:", err)
			continue
		}

		key := fmt.Sprintf("%s|%s", article.Title, article.Link)
		if _, exists := unique[key]; !exists && article.Valid() {
			unique[key] = struct{}{}
			articles = append(articles, article)
			log.Printf("📝 发现文章: %s - %s\n", article.Title, article.Link)
		}
	}

	if len(articles) == 0 {
		log.Println("⚠️ 未找到有效文章，请检查选择器配置")
	}
	return articles
}

// 解析单个文章元素
func parseArticle(el *rod.Element) (Article, error) {
	var article Article

	// 获取标题
	titleEl, err := el.Element(viper.GetString("titleSelector"))
	if err != nil {
		return article, fmt.Errorf("获取标题元素失败: %w", err)
	}
	article.Title = processTitle(titleEl.MustText())

	// 获取链接
	linkEl, err := el.Element("a")
	if err != nil {
		return article, fmt.Errorf("获取链接元素失败: %w", err)
	}

	href, err := linkEl.Attribute("href")
	if err != nil || href == nil {
		return article, fmt.Errorf("获取链接地址失败: %w", err)
	}
	article.Link = normalizeLink(*href)

	return article, nil
}

// 处理标题格式
func processTitle(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ToLower(title)
	return strings.Join(strings.Fields(title), " ") // 处理所有空白字符
}

// 标准化链接格式
func normalizeLink(link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}
	return strings.TrimSuffix(viper.GetString("targetURL"), "/") + "/" + strings.TrimPrefix(link, "/")
}

// 保存文章到文件
func saveArticles(articles []Article) {
	if len(articles) == 0 {
		return
	}

	var builder strings.Builder
	for _, a := range articles {
		builder.WriteString(fmt.Sprintf("%s - %s\n", a.Title, a.Link))
	}

	if err := os.WriteFile(viper.GetString("outputFileName"), []byte(builder.String()), 0644); err != nil {
		log.Fatal("❌ 文件写入失败:", err)
	}
	log.Printf("✅ 成功保存 %d 篇文章到 %s", len(articles), viper.GetString("outputFileName"))
}

// 校验文章有效性
func (a Article) Valid() bool {
	return a.Title != "" && a.Link != "" &&
		strings.HasPrefix(a.Link, "http") &&
		!strings.Contains(a.Link, "about:blank")
}
