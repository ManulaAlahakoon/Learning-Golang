package main

import ("fmt"
"net/http"
"io"
"encoding/xml")

type UrlIndex struct {
	Locations []Location `xml:"url"` 
	//Must captalize variable names in here
	//Titles []Title  `xml:"news>news>title"`
}

type Location struct {
	Loc string `xml:"loc"`
}

type Title struct {
	Ti string `xml:",Cdata"`
}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var s UrlIndex
	xml.Unmarshal(bytes,&s)

	//fmt.Println(s.Locations)
	//fmt.Println(s.Titles)

	fmt.Println(s.Locations)

	
}