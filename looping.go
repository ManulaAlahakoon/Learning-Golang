
package main

import ("fmt"
"net/http"
"io"
"encoding/xml"
)

type SitemapIndex struct {
	Locations []Location `xml:"url"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	
	for _, Loc := range s.Locations {
		fmt.Printf("\n%s",Loc)
	}
}

