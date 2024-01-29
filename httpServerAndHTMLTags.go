package main

import ("fmt"
"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Make %s projects every month</h1>","<strong>2</strong>")
	fmt.Fprintf(w,"<p>Have small improvements often , everyday</p>")
	fmt.Fprintf(w,`<p>Live relax</p>
<h1>Live joyfully</h1>
<p>Live freely</p>`)
}

func main(){
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000",nil)
}