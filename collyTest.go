package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println("链接文本:", e.Text)
		fmt.Println("链接地址:", e.Attr("href"))
	})

	c.Visit("https://example.com")
}
