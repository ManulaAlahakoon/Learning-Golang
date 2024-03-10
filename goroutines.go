package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	
}

func main() {

	//First collection of goroutines
	wg.Add(1)
	go say("Hi")
	wg.Add(1)
	go say("Manula")
	wg.Wait()

	//Second collection of goroutines
	wg.Add(1)
	go say("Mother is making dinner")
	wg.Add(1)
	go say("My wife is beautiful")
	wg.Wait()

}