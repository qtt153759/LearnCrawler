package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"github.com/gocolly/colly"
	"io/ioutil"
)

// Course stores information about a coursera course
type Course struct {
	Title       string
	Description string
	URL					string
	Creator     string
	Star			  string
	Rating      string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("coursera.org", "www.coursera.org"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./coursera_cache"),
		colly.MaxDepth(5),
		colly.Async(true), 
		
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	// Create another collector to scrape course details
	detailCollector := c.Clone()

	courses := make([]Course, 0, 200)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")
		// If link start with browse or includes either signup or login return from callback
		if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}

		// start scaping the page under the link found
		e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		// Activate detailCollector if the link contains "coursera.org/learn"
		courseURL := e.Request.AbsoluteURL(e.Attr("href"))

		if strings.Index(courseURL, "coursera.org/learn") != -1 {
			detailCollector.Visit(courseURL)
		}
	})

	// Extract details of the course
	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText(".banner-title")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			Title:       title,
			URL:         e.Request.URL.String(),
			Creator:     e.ChildText("._1qfi0x77 >._1qfi0x77> span"),
			Star:				 e.ChildText("span[data-test=number-star-rating]"),
			Rating:			 e.ChildText("span[data-test=ratings-count-without-asterisks]>span"),
		}
		// Iterate over rows of the table which contains different information
		// about the course
		
		courses = append(courses, course)
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://coursera.org/browse")
	c.Wait()
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")

	// Dump json to the standard output
	// enc.Encode(courses)
	content,err:=json.Marshal(courses)
	if err!=nil{
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("./coursera/coursera.json",content,0644)
}