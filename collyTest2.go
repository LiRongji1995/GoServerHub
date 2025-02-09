package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

// PageData Define a data structure to store webpage information
type PageData struct {
	Title   string   `json:"title"`   //Page title
	Links   []string `json:"links"`   //All links
	Content string   `json:"content"` //Main page content
}

func main() {
	// Create a Colly crawler instance
	c := colly.NewCollector(
		colly.AllowedDomains("example.com"),
		colly.MaxDepth(2),
	)
	// Set request timeout to prevent freezing
	c.SetRequestTimeout(10)

	// Store crawl results
	var results []PageData

	c.OnHTML("html", func(e *colly.HTMLElement) {
		page := PageData{
			Title:   e.ChildText("title"),                //Get page title
			Content: strings.TrimSpace(e.ChildText("p")), //Extract main content
		}

		//Iterate through all <a> tags to get links
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")                          //Properties of HTML hyperlinks (<a> tags)
			if link != "" && !strings.HasPrefix(link, "#") { //Filter out invalid links
				page.Links = append(page.Links, link)
				// If a new link is found, add it to the crawl queue
				if err := c.Visit(e.Request.AbsoluteURL(link)); err != nil {
					fmt.Println("访问链接失败：", err)
				}
			}
		})

		//store the result
		results = append(results, page)
		fmt.Println("爬取完成", page.Title) //Print success message
	})
	// Handle request errors
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("请求失败", r.Request.URL, err) //Print error message
	})

	//After crawling is complete, save data to JSON
	c.OnScraped(func(_ *colly.Response) {
		file, err := os.Create("data.json") // Create JSON file
		if err != nil {
			log.Fatal("无法创建JSON文件：", err) // Print error if file creation fails
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Fatal("无法关闭文件", err)
			}
		}()

		jsonData, _ := json.MarshalIndent(results, "", "  ") //Format JSON
		write, err := file.Write(jsonData)                   //Write to file
		if err != nil {
			log.Fatal("写入文件失败：", err)
		}
		fmt.Printf("写入 %d 字节\n", write)
		fmt.Println("数据已保存到 data.Json") // Print success message
	})

	//Start crawling
	startURL := "https://example.com"
	fmt.Println("开始爬取:", startURL) //Print start message
	if err := c.Visit(startURL); err != nil {
		fmt.Println("访问链接失败：", err)
	}
}
