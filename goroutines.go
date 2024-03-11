package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	defer fmt.Println("DONE")
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup : ",r)
	}
}

func say(s string) {
	//defer wg.Done()
	defer cleanup()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 3 {
			panic("Stop ... SOS")
		}
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