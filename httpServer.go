package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go is running")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Manula Alahakkon is learning GO")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)

}