package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML(".sidebar-link", func(element *colly.HTMLElement) {
		element.Request.Visit(element.Attr("href"))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.Body)
	})

	c.Visit("https://gorm.io/zh_CN/docs/")
}
