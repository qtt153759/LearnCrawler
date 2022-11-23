package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/gocolly/colly"
	"io/ioutil"
)

// Course stores information about a coursera course
type Post struct {
	Title       					string
	Comment 							string
	Created_Date					string
	Reaction     					int

}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("facebook.com", "m.facebook.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./facebook_cache"),
		colly.MaxDepth(1),
		colly.Async(true), 
		
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	// Create another collector to scrape course details
	// detailCollector := c.Clone()

	post := Post{}
	// On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	// If attribute class is this long string return from callback
	// 	// As this a is irrelevant
	// 	if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
	// 		return
	// 	}
	// 	link := e.Attr("href")
	// 	// If link start with browse or includes either signup or login return from callback
	// 	if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
	// 		return
	// 	}

	// 	// start scaping the page under the link found
	// 	e.Request.Visit(link)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})
	
	// On every a HTML element which has name attribute call callback
	c.OnHTML(`p`, func(e *colly.HTMLElement) {
		
		post.Title=e.Text
	})
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		
		// a,err:=e.DOM.Html()
		// if err!=nil{
			
		// }
		fmt.Println("ok",e.Dom.html())
		// post.Comment+=e.Text+" | "
	})

	// Extract details of the course
	// detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
	// 	log.Println("Course found", e.Request.URL)
	// 	title := e.ChildText(".banner-title")
	// 	if title == "" {
	// 		log.Println("No title found", e.Request.URL)
	// 	}
	// 	course := Course{
	// 		Title:       title,
	// 		URL:         e.Request.URL.String(),
	// 		Creator:     e.ChildText("._1qfi0x77 >._1qfi0x77> span"),
	// 		Star:				 e.ChildText("span[data-test=number-star-rating]"),
	// 		Rating:			 e.ChildText("span[data-test=ratings-count-without-asterisks]>span"),
	// 	}
		// Iterate over rows of the table which contains different information
		// about the course
		
	// 	courses = append(courses, course)
	// })

	c.Visit("https://m.facebook.com/groups/voz.congnghe/permalink/1293314158135513/?m_entstream_source=group&paipv=0&eav=AfZ69KxXZsnoCB_jS2qOiDfkqLp7Py9a1_KHN2XNset6vIwak281UFzc_Bh3lr-2nQQ")
	c.Wait()
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")

	// Dump json to the standard output
	// enc.Encode(courses)
	content,err:=json.Marshal(post)
	if err!=nil{
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("./facebook/post.json",content,0644)
}