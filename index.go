package main

import ("fmt"
"net/http"
"io"
"encoding/xml"
)

type Sitemapindex struct {
	Locations []string `xml:"url>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Dates []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Date string
	Location string
}

func main() {
	//var s Sitemapindex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)
	//resp.Body.Close()
	xml.Unmarshal(bytes, &n)
	
	

			for idx,_ := range n.Locations{
				news_map[n.Titles[idx]] = NewsMap{n.Dates[idx],n.Locations[idx]}
			}


	for id, data := range news_map {
		fmt.Println("\nTitle = ",id)
		fmt.Println("\nLocation = ",data.Location)
		fmt.Println("\nDate = ",data.Date)
	}
}