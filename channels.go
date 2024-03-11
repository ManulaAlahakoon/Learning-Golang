package main

import "fmt"

func pass(c chan int, value int){
	c <- value*10 //Send
}
func main() {
	channelA := make(chan int)

	go pass(channelA, 5)
	go pass(channelA,30)

	//v1 := <-channelA  //Get
	//v2 := <-channelA

	v1, v2 := <- channelA, <- channelA
	fmt.Println(v1,v2)
}