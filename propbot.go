package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getHtml(url string) string {

	return ""

}

// Things I'd need to get:
//   * Current Price
//   * Place bid link
//

var (
	url = "https://www.propertyroom.com/l/apple-macbook-air-laptop/16170499?source=CATPANEL"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("propertyroom.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[*] Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("[-] Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[*] Visited", r.Request.URL)
	})

	c.OnHTML(".current-price-propductpage", func(e *colly.HTMLElement) {
		// I believe that this is where you do the needfull
		// e.Request.Visit(e.Attr("href"))
		fmt.Printf("Found: %+v", e.DOM)
	})

	// c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	// 	fmt.Println("First column of a table row:", e.Text)
	// })

	// c.OnXML("//h1", func(e *colly.XMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)

}
