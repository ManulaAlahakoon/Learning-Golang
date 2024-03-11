
package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sync"
)

/*
type Sitemapindex struct {
	Locations []string `xml:"url>loc"`
}
*/

var wg sync.WaitGroup

type News struct {
	Titles    []string `xml:"url>news>title"`
	Dates     []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Date     string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}
func newsRoutine(c chan News) {
	defer wg.Done()
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	news_map := make(map[string]NewsMap)

	channel := make(chan News, 1)
	wg.Add(1)
	go newsRoutine(channel)
	wg.Wait()
	close(channel)

	for elem := range channel {
	for idx, _ := range elem.Locations {
		news_map[elem.Titles[idx]] = NewsMap{elem.Dates[idx], elem.Locations[idx]}
		}
	}
		
	p := NewsAggPage{Title: "News aggregator", News: news_map}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Let's go</h1>")
}



func main() {
	//var s Sitemapindex
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}


