package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type PageData struct {
	Title   string   `json:"title"`
	Links   []string `json:"links"`
	Content string   `json:"content"`
}

var (
	results []PageData
	mu      sync.Mutex
)

func isValidLink(link string) bool {
	invalidPrefixes := []string{"#", "tel:", "mailto:", "javascript:", "about:"}
	for _, prefix := range invalidPrefixes {
		if strings.HasPrefix(link, prefix) {
			return false
		}
	}
	return true
}

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var visited sync.Map

	c.OnRequest(func(r *colly.Request) {
		if _, exists := visited.Load(r.URL.String()); exists || r.Depth > 2 {
			r.Abort()
			return
		}
		visited.Store(r.URL.String(), true)

		r.Headers.Set("User-Agent", userAgents[rng.Intn(len(userAgents))])
	})

	c.SetRequestTimeout(60 * time.Second)
	c.OnHTML("html", func(e *colly.HTMLElement) {
		var content strings.Builder
		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			content.WriteString(strings.TrimSpace(el.Text) + "\n")
		})

		page := PageData{
			Title:   e.ChildText("title"),
			Content: content.String(),
		}

		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")
			if !isValidLink(link) {
				return
			}

			absoluteURL := e.Request.AbsoluteURL(link)
			if absoluteURL == "" || !strings.HasPrefix(absoluteURL, "http") {
				return
			}

			if strings.Contains(absoluteURL, "logout") || strings.Contains(absoluteURL, "x.com") {
				return
			}

			// **改这里：使用 sync.Map 代替普通 map**
			if _, exists := visited.Load(absoluteURL); !exists {
				visited.Store(absoluteURL, true) // **确保并发安全**
				if err := c.Visit(absoluteURL); err != nil {
					log.Println("Failed to visit link:", err)
				}
			}
		})

		mu.Lock()
		results = append(results, page)
		mu.Unlock()
		log.Println("爬取完成:", page.Title)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("请求失败:", r.Request.URL, err)
	})

	c.OnScraped(func(_ *colly.Response) {
		file, err := os.Create("data.json")
		if err != nil {
			log.Fatal("无法创建JSON文件：", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(results); err != nil {
			log.Fatal("写入文件失败：", err)
		}
		log.Println("数据已保存到 data.json")
	})

	startURL := "https://example.com"
	log.Println("开始爬取:", startURL)
	if err := c.Visit(startURL); err != nil {
		log.Println("访问链接失败：", err)
	}
	c.Wait()
}
