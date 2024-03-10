package main

import ("fmt"
"net/http"
"io"
"encoding/xml"
"html/template"
)
/*
type Sitemapindex struct {
	Locations []string `xml:"url>loc"`
}
*/
type News struct {
	Titles []string `xml:"url>news>title"`
	Dates []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Date string
	Location string
}

type NewsAggPage struct {
	Title string
	//News string
	News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){

	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &n)
	
	for idx,_ := range n.Locations{
		news_map[n.Titles[idx]] = NewsMap{n.Dates[idx],n.Locations[idx]}
	}

	p := NewsAggPage{Title: "All Around - A2 News", News: news_map}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w,p)

}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Let's go</h1>")
}


func main() {
	//var s Sitemapindex
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000",nil)

}