
package main

import ("fmt"
"net/http"
"io"
"encoding/xml"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>titles"`
	Dates []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)
	//resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	
	for _, Location string := range s.Locations {
			resp, _ := http.Get(Location)
			bytes, _ := io.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &n)
	}
}

