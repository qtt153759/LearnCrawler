package main

import (
	// "encoding/json"
	"fmt"
	"log"
	// "strings"
	"github.com/gocolly/colly"
	// "io/ioutil"
)

// Course stores information about a coursera course


func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("shopee.vn"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./shopee_cache"),
		colly.MaxDepth(2),
		colly.Async(true), 
		
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	// Create another collector to scrape course details
	// detailCollector := c.Clone()

	// courses := make([]Course, 0, 200)

	// On every a element which has href attribute call callback
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		// start scaping the page under the link found
		// link := e.Request.AbsoluteURL(e.Attr("data-seq"))
		log.Println("find")
		h, err:=e.DOM.Html()
		if err==nil{

			fmt.Println("ok",h)
		}

		// detailCollector.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	// c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
	// 	// Activate detailCollector if the link contains "coursera.org/learn"
	// 	courseURL := e.Request.AbsoluteURL(e.Attr("href"))

	// 	if strings.Index(courseURL, "coursera.org/learn") != -1 {
	// 		detailCollector.Visit(courseURL)
	// 	}
	// })

	// Extract details of the course
	// detailCollector.OnHTML(`div[class=_2rQP1z]`, func(e *colly.HTMLElement) {
	// 	log.Println("Course found", e.Request.URL)
	// 	title := e.ChildText("span")
	// 	if title == "" {
	// 		log.Println("No title found", e.Request.URL)
	// 	}else{
	// 		log.Println("title",title)
	// 	}
	// 	// Iterate over rows of the table which contains different information
	// 	// about the course
		
		
	// })
	// detailCollector.OnRequest(func(r *colly.Request) {
	// 	log.Println("visiting", r.URL.String())
	// })
	// Start scraping on http://coursera.com/browse
	c.Visit("https://shopee.vn/Thi%E1%BA%BFt-B%E1%BB%8B-%C4%90i%E1%BB%87n-T%E1%BB%AD-cat.11036132")
	c.Wait()
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")

	// Dump json to the standard output
	// enc.Encode(courses)
	// content,err:=json.Marshal(courses)
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }
	// ioutil.WriteFile("./coursera/coursera.json",content,0644)
}