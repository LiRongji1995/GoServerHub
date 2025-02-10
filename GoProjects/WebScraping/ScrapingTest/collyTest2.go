package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
		//colly.AllowedDomains("example.com"),
		colly.MaxDepth(2),
		colly.AllowURLRevisit(), //Enables revisiting previously visited URLs instead of skipping them
	)
	// Set timeout to prevent handshake failures
	c.SetRequestTimeout(60 * time.Second)

	// Create a map to store visited URLs to avoid duplicates
	visited := make(map[string]bool)

	// Prevent duplicate visits by checking if a URL has already been visited
	c.OnRequest(func(r *colly.Request) {
		if visited[r.URL.String()] {
			fmt.Println("Skipping already visited URL:", r.URL.String())
			r.Abort() // Cancel the request to prevent duplicate crawling
			return
		}
		visited[r.URL.String()] = true
	})

	// Set request timeout to prevent freezing
	c.SetRequestTimeout(60 * time.Second)

	// Define a list of common User-Agent strings to avoid detection
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	}

	// Seed the random number generator to ensure different User-Agents are used
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Set a random User-Agent for each request
	c.OnRequest(func(r *colly.Request) {
		randomUserAgent := userAgents[rng.Intn(len(userAgents))] // Select a random User-Agent
		r.Headers.Set("User-Agent", randomUserAgent)             // Apply the User-Agent to the request
		fmt.Println("Using User-Agent:", randomUserAgent)        // Print the selected User-Agent
	})
	// Store crawl results
	var results []PageData

	c.OnHTML("html", func(e *colly.HTMLElement) {
		page := PageData{
			Title:   e.ChildText("title"),                // Get page title
			Content: strings.TrimSpace(e.ChildText("p")), // Extract main content
		}

		//Iterate through all <a> tags to get links
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			link := el.Attr("href")

			// Skip invalid links
			if link == "" ||
				strings.HasPrefix(link, "#") ||
				strings.HasPrefix(link, "tel:") ||
				strings.HasPrefix(link, "mailto:") ||
				strings.HasPrefix(link, "javascript:") ||
				strings.HasPrefix(link, "about:") {
				return // Ignore invalid links
			}

			absoluteURL := e.Request.AbsoluteURL(link)
			if absoluteURL == "" || !strings.HasPrefix(absoluteURL, "http") {
				return // Skip empty or non-HTTP URLs
			}

			// Skip logout and authentication-related URLs
			if strings.Contains(absoluteURL, "logout") ||
				strings.Contains(absoluteURL, "signout") ||
				strings.Contains(absoluteURL, "session_end") ||
				strings.Contains(absoluteURL, "authorize") ||
				strings.Contains(absoluteURL, "redirect_uri") ||
				strings.Contains(absoluteURL, "oauth") {
				fmt.Println("Skipping authentication/logout URL:", absoluteURL)
				return
			}

			// Skip blocked or unresponsive sites
			if strings.Contains(absoluteURL, "x.com") {
				fmt.Println("Skipping restricted site:", absoluteURL)
				return
			}

			if err := c.Visit(absoluteURL); err != nil {
				fmt.Println("Failed to visit link:", err)
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
