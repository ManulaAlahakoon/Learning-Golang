package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func pass(c chan int, value int){
	defer wg.Done()
	c <- value*10 //Send
}
func main() {
	channelA := make(chan int,10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go pass(channelA,i)
	}
	wg.Wait()
	close(channelA)

	for item := range channelA {
		fmt.Println(item)
	}
	
	//go pass(channelA, 5)
	//go pass(channelA,30)

	//v1 := <-channelA  //Get
	//v2 := <-channelA

	//v1, v2 := <- channelA, <- channelA
	//fmt.Println(v1,v2)
}