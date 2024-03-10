package main 

import (
	"fmt"
	"net/http"
	"html/template"
)

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	p := NewsAggPage{Title: "Learning Golang is fun", News: "Yeah...It is"}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w,p)

}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Let's go</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000",nil)
}

